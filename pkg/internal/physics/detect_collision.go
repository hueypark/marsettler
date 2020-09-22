package physics

import (
	"math"

	math2 "github.com/hueypark/marsettler/pkg/internal/math"
	"github.com/hueypark/marsettler/pkg/internal/physics/body"
	"github.com/hueypark/marsettler/pkg/internal/physics/body/circle"
	"github.com/hueypark/marsettler/pkg/internal/physics/body/convex"
	"github.com/hueypark/marsettler/pkg/internal/physics/closest_point"
)

func (c *Contact) DetectCollision() {
	lhsType := c.lhs.Shape.Type()
	rhsType := c.rhs.Shape.Type()

	switch lhsType {
	case body.Bullet:
		switch rhsType {
		case body.Bullet:
			break
		case body.Circle:
			c.normal, c.penetration, c.points = bulletToCircle(c.lhs, c.rhs)
			break
		case body.Convex:
			c.normal, c.penetration, c.points = bulletToConvex(c.lhs, c.rhs)
			break
		}
		break
	case body.Circle:
		switch rhsType {
		case body.Bullet:
			c.swap()
			c.normal, c.penetration, c.points = bulletToCircle(c.lhs, c.rhs)
			break
		case body.Circle:
			c.normal, c.penetration, c.points = circleToCircle(c.lhs, c.rhs)
			break
		case body.Convex:
			c.normal, c.penetration, c.points = circleToConvex(c.lhs, c.rhs)
			break
		}
		break
	case body.Convex:
		switch rhsType {
		case body.Bullet:
			c.swap()
			c.normal, c.penetration, c.points = bulletToConvex(c.lhs, c.rhs)
		case body.Circle:
			c.swap()
			c.normal, c.penetration, c.points = circleToConvex(c.lhs, c.rhs)
			break
		case body.Convex:
			c.normal, c.penetration, c.points = convexToConvex(c.lhs, c.rhs)
			break
		}
		break
	}
}

func (c *Contact) swap() {
	c.lhs, c.rhs = c.rhs, c.lhs
}

func bulletToCircle(
	lhs, rhs *body.Body,
) (normal math2.Vector, penetration float64, points []math2.Vector) {
	rhsCircle := rhs.Shape.(*circle.Circle)

	normal = math2.Sub(rhs.Position(), lhs.Position())

	distanceSquared := normal.SizeSquare()

	if distanceSquared >= rhsCircle.Radius*rhsCircle.Radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	normal = normal.Normalize()
	penetration = rhsCircle.Radius - distance
	points = append(points, math2.Add(
		lhs.Position(),
		math2.Mul(normal, -0.5*penetration)))

	return normal, penetration, points
}

func bulletToConvex(
	lhs, rhs *body.Body,
) (normal math2.Vector, penetration float64, points []math2.Vector) {
	rhsConvex := rhs.Shape.(*convex.Convex)

	penetration = math.MaxFloat64

	for _, edge := range rhsConvex.Edges() {
		worldStart := rhs.Rotation().RotateVector(edge.Start)
		worldStart.Add(rhs.Position())
		worldEnd := rhs.Rotation().RotateVector(edge.End)
		worldEnd.Add(rhs.Position())
		edgeVector := math2.Sub(worldEnd, worldStart)
		pointVector := math2.Sub(lhs.Position(), worldStart)

		if !pointVector.OnTheRight(edgeVector) {
			normal = math2.Vector{}
			penetration = 0
			return normal, penetration, points
		}

		perpendicular := math2.Vector{X: -edgeVector.Y, Y: edgeVector.X}
		perpendicular = perpendicular.Normalize()

		lhsVector := math2.Sub(lhs.Position(), worldStart)

		proj := math2.Mul(perpendicular, math2.Dot(lhsVector, perpendicular))

		if proj.Size() < penetration {
			normal = perpendicular
			penetration = proj.Size()
		}
	}

	points = append(points, math2.Add(
		lhs.Position(),
		math2.Mul(normal, -0.5*penetration)))

	return normal, penetration, points
}

func circleToCircle(
	lhs, rhs *body.Body,
) (normal math2.Vector, penetration float64, points []math2.Vector) {
	lhsCircle := lhs.Shape.(*circle.Circle)
	rhsCircle := rhs.Shape.(*circle.Circle)

	normal = math2.Sub(rhs.Position(), lhs.Position())

	distanceSquared := normal.SizeSquare()
	radius := lhsCircle.Radius + rhsCircle.Radius

	if distanceSquared >= radius*radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	normal = normal.Normalize()
	penetration = radius - distance
	points = append(points, math2.Add(
		lhs.Position(),
		math2.Add(math2.Mul(normal, lhsCircle.Radius), math2.Mul(normal, -0.5*penetration))))

	return normal, penetration, points
}

func circleToConvex(
	l, r *body.Body,
) (normal math2.Vector, penetration float64, points []math2.Vector) {
	lCircle := l.Shape.(*circle.Circle)
	rConvex := r.Shape.(*convex.Convex)

	minPenetration := math.MaxFloat64
	var minPenEdge convex.Edge
	for _, edge := range rConvex.Edges() {
		edgeNormal := r.Rotation().RotateVector(edge.Normal)
		edgeStart := r.Rotation().RotateVector(edge.Start)

		pen := -math2.Dot(edgeNormal, math2.Sub(l.Position(), math2.Add(r.Position(), edgeStart)))

		if pen < -lCircle.Radius {
			return math2.Zero(), 0, nil
		}

		if pen < minPenetration {
			minPenetration = pen
			minPenEdge = edge
		}
	}

	edgeStart := math2.Add(r.Position(), r.Rotation().RotateVector(minPenEdge.Start))
	edgeEnd := math2.Add(r.Position(), r.Rotation().RotateVector(minPenEdge.End))
	edgeNormal := r.Rotation().RotateVector(minPenEdge.Normal)

	p := closest_point.LineSegmentToPoint(l.Position(), edgeStart, edgeEnd)
	if 0 < math2.Dot(edgeNormal, math2.Sub(p, l.Position())) {
		normal = math2.Sub(l.Position(), p)

		penetration = lCircle.Radius + math2.Sub(l.Position(), p).Size()
	} else {
		normal = math2.Sub(p, l.Position())
		if lCircle.Radius < normal.Size() {
			return math2.Zero(), 0, nil
		}

		penetration = lCircle.Radius - math2.Sub(p, l.Position()).Size()
	}

	normal = normal.Normalize()
	points = append(points, math2.Add(p, math2.Mul(normal, 0.5*penetration)))

	return normal, penetration, points
}

func convexToConvex(
	l, r *body.Body,
) (normal math2.Vector, penetration float64, points []math2.Vector) {
	lConvex := l.Shape.(*convex.Convex)
	rConvex := r.Shape.(*convex.Convex)

	lPenetration, lNormal, lPoint := findAxisLeastPenetration(lConvex, rConvex, l.Position(), r.Position(), l.Rotation(), r.Rotation())
	if lPenetration < 0.0 {
		return normal, penetration, points
	}

	rPenetration, rNormal, rPoint := findAxisLeastPenetration(rConvex, lConvex, r.Position(), l.Position(), r.Rotation(), l.Rotation())
	if rPenetration < 0.0 {
		return normal, penetration, points
	}

	if lPenetration < rPenetration {
		normal = lNormal
		penetration = lPenetration
		points = append(points, lPoint)
	} else {
		normal = math2.Invert(rNormal)
		penetration = rPenetration
		points = append(points, rPoint)
	}

	return normal, penetration, points
}

func findAxisLeastPenetration(
	l, r *convex.Convex, lPos, rPos math2.Vector, lRot, rRot math2.Rotator,
) (minPenetration float64, bestNormal math2.Vector, bestPoint math2.Vector) {
	minPenetration = math.MaxFloat64

	for _, edge := range l.Edges() {
		normal := lRot.RotateVector(edge.Normal)
		s := r.Support(math2.Invert(normal), rRot)

		v := math2.Add(lRot.RotateVector(edge.Start), lPos)
		v.Sub(rPos)

		penetration := -math2.Dot(normal, math2.Sub(s, v))

		if penetration < minPenetration {
			bestNormal = normal
			minPenetration = penetration
			bestPoint = math2.Add(math2.Add(s, rPos), math2.Mul(bestNormal, penetration*0.5))
		}
	}

	return minPenetration, bestNormal, bestPoint
}
