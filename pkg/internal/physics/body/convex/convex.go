package convex

import (
	"math"

	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/physics/body"
)

type Convex struct {
	vertices []*math2d.Vector
	hull     []*math2d.Vector
	edges    []Edge
}

type Edge struct {
	Start  *math2d.Vector
	End    *math2d.Vector
	Normal *math2d.Vector
}

func New(vertices []*math2d.Vector) *Convex {
	c := Convex{vertices, nil, nil}

	return &c
}

func (c *Convex) Type() body.Shape {
	return body.Convex
}

// Hull is ccw
func (c *Convex) Hull() []*math2d.Vector {
	if c.hull == nil {
		minX, maxX := c.getExtremePoints()
		c.hull = append(c.quickHull(c.vertices, maxX, minX), c.quickHull(c.vertices, minX, maxX)...)
	}

	return c.hull
}

// Edge is ccw
func (c *Convex) Edges() []Edge {
	if c.edges == nil {
		hull := c.Hull()
		for i, start := range hull {
			nextIndex := i + 1
			if len(hull) <= nextIndex {
				nextIndex = 0
			}
			end := hull[nextIndex]
			r := math2d.NewRotator(math.Pi * 0.5)
			normal := r.RotateVector(math2d.Sub(start, end))
			normal.Normalize()
			c.edges = append(c.edges, Edge{
				start,
				end,
				normal})
		}
	}

	return c.edges
}

func MinkowskiDifference(
	a Convex,
	posA *math2d.Vector,
	rotA math2d.Rotator,
	b Convex,
	posB *math2d.Vector,
	rotB math2d.Rotator,
) *Convex {
	var vertices []*math2d.Vector

	for _, vertexA := range a.Hull() {
		for _, vertexB := range b.Hull() {
			vertexRotA := rotA.RotateVector(vertexA)
			vertexRotB := rotB.RotateVector(vertexB)
			worldA := math2d.Add(vertexRotA, posA)
			worldB := math2d.Sub(&math2d.Vector{}, math2d.Add(vertexRotB, posB))
			vertices = append(vertices, math2d.Add(worldA, worldB))
		}
	}

	return New(vertices)
}

func (c *Convex) Support(dir *math2d.Vector, rot math2d.Rotator) (bestVertex *math2d.Vector) {
	bestVertex = &math2d.Vector{}
	bestProjection := -math.MaxFloat64

	for _, vertex := range c.Hull() {
		projection := math2d.Dot(rot.RotateVector(vertex), dir)

		if bestProjection < projection {
			bestVertex = rot.RotateVector(vertex)
			bestProjection = projection
		}
	}

	return bestVertex
}

func (c *Convex) quickHull(points []*math2d.Vector, start, end *math2d.Vector) []*math2d.Vector {
	pointDistanceIndicators := c.getLhsPointDistanceIndicatorMap(points, start, end)
	if len(pointDistanceIndicators) == 0 {
		return []*math2d.Vector{end}
	}

	farthestPoint := c.getFarthestPoint(pointDistanceIndicators)

	var newPoints []*math2d.Vector
	for point := range pointDistanceIndicators {
		newPoints = append(newPoints, point)
	}

	return append(
		c.quickHull(newPoints, farthestPoint, end),
		c.quickHull(newPoints, start, farthestPoint)...)
}

func (c *Convex) InHull(
	position *math2d.Vector, rotation math2d.Rotator, point *math2d.Vector,
) bool {
	for _, edge := range c.Edges() {
		if math2d.OnTheRight(
			math2d.Sub(point, math2d.Add(position, rotation.RotateVector(edge.Start))),
			math2d.Sub(
				math2d.Add(position, rotation.RotateVector(edge.End)),
				math2d.Add(position, rotation.RotateVector(edge.Start)))) == false {
			return false
		}
	}

	return true
}

func (c *Convex) getExtremePoints() (minX, maxX *math2d.Vector) {
	minX = &math2d.Vector{X: math.MaxFloat64, Y: 0}
	maxX = &math2d.Vector{X: -math.MaxFloat64, Y: 0}

	for _, p := range c.vertices {
		if p.X < minX.X {
			minX = p
		}

		if maxX.X < p.X {
			maxX = p
		}
	}

	return minX, maxX
}

func (c *Convex) getLhsPointDistanceIndicatorMap(
	points []*math2d.Vector, start, end *math2d.Vector,
) map[*math2d.Vector]float64 {
	pointDistanceIndicatorMap := make(map[*math2d.Vector]float64)

	for _, point := range points {
		distanceIndicator := c.getDistanceIndicator(point, start, end)
		if distanceIndicator > 0 {
			pointDistanceIndicatorMap[point] = distanceIndicator
		}
	}

	return pointDistanceIndicatorMap
}

func (c *Convex) getDistanceIndicator(point, start, end *math2d.Vector) float64 {
	vLine := math2d.Sub(end, start)

	vPoint := math2d.Sub(point, start)

	return math2d.Cross(vLine, vPoint)
}

func (c *Convex) getFarthestPoint(
	pointDistanceIndicatorMap map[*math2d.Vector]float64,
) (farthestPoint *math2d.Vector) {
	farthestPoint = &math2d.Vector{}

	maxDistanceIndicator := -math.MaxFloat64
	for point, distanceIndicator := range pointDistanceIndicatorMap {
		if maxDistanceIndicator < distanceIndicator {
			maxDistanceIndicator = distanceIndicator
			farthestPoint.Set(point)
		}
	}

	return farthestPoint
}
