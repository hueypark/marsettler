package convex

import (
	"math"

	"github.com/hueypark/marsettler/core/physics/body"
	math2 "github.com/hueypark/marsettler/pkg/internal/math"
)

type Convex struct {
	vertices []math2.Vector
	hull     []math2.Vector
	edges    []Edge
}

type Edge struct {
	Start  math2.Vector
	End    math2.Vector
	Normal math2.Vector
}

func New(vertices []math2.Vector) *Convex {
	c := Convex{vertices, nil, nil}

	return &c
}

func (c *Convex) Type() body.Shape {
	return body.Convex
}

// Hull is ccw
func (c *Convex) Hull() []math2.Vector {
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
			r := math2.NewRotator(math.Pi * 0.5)
			normal := r.RotateVector(math2.Sub(start, end))
			normal = normal.Normalize()
			c.edges = append(c.edges, Edge{
				start,
				end,
				normal})
		}
	}

	return c.edges
}

func MinkowskiDifference(
	a Convex, posA math2.Vector, rotA math2.Rotator, b Convex, posB math2.Vector, rotB math2.Rotator,
) *Convex {
	var vertices []math2.Vector

	for _, vertexA := range a.Hull() {
		for _, vertexB := range b.Hull() {
			vertexRotA := rotA.RotateVector(vertexA)
			vertexRotB := rotB.RotateVector(vertexB)
			worldA := math2.Add(vertexRotA, posA)
			worldB := math2.Sub(math2.Vector{}, math2.Add(vertexRotB, posB))
			vertices = append(vertices, math2.Add(worldA, worldB))
		}
	}

	return New(vertices)
}

func (c *Convex) Support(dir math2.Vector, rot math2.Rotator) (bestVertex math2.Vector) {
	bestProjection := -math.MaxFloat64

	for _, vertex := range c.Hull() {
		projection := math2.Dot(rot.RotateVector(vertex), dir)

		if bestProjection < projection {
			bestVertex = rot.RotateVector(vertex)
			bestProjection = projection
		}
	}

	return bestVertex
}

func (c *Convex) quickHull(points []math2.Vector, start, end math2.Vector) []math2.Vector {
	pointDistanceIndicators := c.getLhsPointDistanceIndicatorMap(points, start, end)
	if len(pointDistanceIndicators) == 0 {
		return []math2.Vector{end}
	}

	farthestPoint := c.getFarthestPoint(pointDistanceIndicators)

	newPoints := []math2.Vector{}
	for point := range pointDistanceIndicators {
		newPoints = append(newPoints, point)
	}

	return append(
		c.quickHull(newPoints, farthestPoint, end),
		c.quickHull(newPoints, start, farthestPoint)...)
}

func (c *Convex) InHull(position math2.Vector, rotation math2.Rotator, point math2.Vector) bool {
	for _, edge := range c.Edges() {
		if math2.Sub(point, math2.Add(position, rotation.RotateVector(edge.Start))).OnTheRight(math2.Sub(math2.Add(position, rotation.RotateVector(edge.End)), math2.Add(position, rotation.RotateVector(edge.Start)))) == false {
			return false
		}
	}

	return true
}

func (c *Convex) getExtremePoints() (minX, maxX math2.Vector) {
	minX = math2.Vector{math.MaxFloat64, 0}
	maxX = math2.Vector{-math.MaxFloat64, 0}

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
	points []math2.Vector, start, end math2.Vector,
) map[math2.Vector]float64 {
	pointDistanceIndicatorMap := make(map[math2.Vector]float64)

	for _, point := range points {
		distanceIndicator := c.getDistanceIndicator(point, start, end)
		if distanceIndicator > 0 {
			pointDistanceIndicatorMap[point] = distanceIndicator
		}
	}

	return pointDistanceIndicatorMap
}

func (c *Convex) getDistanceIndicator(point, start, end math2.Vector) float64 {
	vLine := math2.Sub(end, start)

	vPoint := math2.Sub(point, start)

	return math2.Cross(vLine, vPoint)
}

func (c *Convex) getFarthestPoint(
	pointDistanceIndicatorMap map[math2.Vector]float64,
) (farthestPoint math2.Vector) {
	maxDistanceIndicator := -math.MaxFloat64
	for point, distanceIndicator := range pointDistanceIndicatorMap {
		if maxDistanceIndicator < distanceIndicator {
			maxDistanceIndicator = distanceIndicator
			farthestPoint = point
		}
	}

	return farthestPoint
}
