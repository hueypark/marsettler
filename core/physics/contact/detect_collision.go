package contact

import (
	"math"

	"github.com/hueypark/heavycannon/body"
	"github.com/hueypark/heavycannon/closest_point"
	"github.com/hueypark/heavycannon/math/rotator"
	"github.com/hueypark/heavycannon/math/vector"
	"github.com/hueypark/heavycannon/shape"
	"github.com/hueypark/heavycannon/shape/circle"
	"github.com/hueypark/heavycannon/shape/convex"
)

func (c *Contact) DetectCollision() {
	lhsType := c.lhs.Shape.Type()
	rhsType := c.rhs.Shape.Type()

	switch lhsType {
	case shape.BULLET:
		switch rhsType {
		case shape.BULLET:
			break
		case shape.CIRCLE:
			c.normal, c.penetration, c.points = bulletToCircle(c.lhs, c.rhs)
			break
		case shape.CONVEX:
			c.normal, c.penetration, c.points = bulletToConvex(c.lhs, c.rhs)
			break
		}
		break
	case shape.CIRCLE:
		switch rhsType {
		case shape.BULLET:
			c.swap()
			c.normal, c.penetration, c.points = bulletToCircle(c.lhs, c.rhs)
			break
		case shape.CIRCLE:
			c.normal, c.penetration, c.points = circleToCircle(c.lhs, c.rhs)
			break
		case shape.CONVEX:
			c.normal, c.penetration, c.points = circleToConvex(c.lhs, c.rhs)
			break
		}
		break
	case shape.CONVEX:
		switch rhsType {
		case shape.BULLET:
			c.swap()
			c.normal, c.penetration, c.points = bulletToConvex(c.lhs, c.rhs)
		case shape.CIRCLE:
			c.swap()
			c.normal, c.penetration, c.points = circleToConvex(c.lhs, c.rhs)
			break
		case shape.CONVEX:
			c.normal, c.penetration, c.points = convexToConvex(c.lhs, c.rhs)
			break
		}
		break
	}
}

func (c *Contact) swap() {
	c.lhs, c.rhs = c.rhs, c.lhs
}

func bulletToCircle(lhs, rhs *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	rhsCircle := rhs.Shape.(*circle.Circle)

	normal = vector.Subtract(rhs.Position(), lhs.Position())

	distanceSquared := normal.SizeSquared()

	if distanceSquared >= rhsCircle.Radius*rhsCircle.Radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	normal.Normalize()
	penetration = rhsCircle.Radius - distance
	points = append(points, vector.Add(
		lhs.Position(),
		vector.Multiply(normal, -0.5*penetration)))

	return normal, penetration, points
}

func bulletToConvex(lhs, rhs *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	rhsConvex := rhs.Shape.(*convex.Convex)

	penetration = math.MaxFloat64

	for _, edge := range rhsConvex.Edges() {
		worldStart := rhs.Rotation().RotateVector(edge.Start)
		worldStart.Add(rhs.Position())
		worldEnd := rhs.Rotation().RotateVector(edge.End)
		worldEnd.Add(rhs.Position())
		edgeVector := vector.Subtract(worldEnd, worldStart)
		pointVector := vector.Subtract(lhs.Position(), worldStart)

		if !pointVector.OnTheRight(edgeVector) {
			normal = vector.Vector{}
			penetration = 0
			return normal, penetration, points
		}

		perpendicular := vector.Vector{X: -edgeVector.Y, Y: edgeVector.X}
		perpendicular.Normalize()

		lhsVector := vector.Subtract(lhs.Position(), worldStart)

		proj := vector.Multiply(perpendicular, vector.Dot(lhsVector, perpendicular))

		if proj.Size() < penetration {
			normal = perpendicular
			penetration = proj.Size()
		}
	}

	points = append(points, vector.Add(
		lhs.Position(),
		vector.Multiply(normal, -0.5*penetration)))

	return normal, penetration, points
}

func circleToCircle(lhs, rhs *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	lhsCircle := lhs.Shape.(*circle.Circle)
	rhsCircle := rhs.Shape.(*circle.Circle)

	normal = vector.Subtract(rhs.Position(), lhs.Position())

	distanceSquared := normal.SizeSquared()
	radius := lhsCircle.Radius + rhsCircle.Radius

	if distanceSquared >= radius*radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	normal.Normalize()
	penetration = radius - distance
	points = append(points, vector.Add(
		lhs.Position(),
		vector.Add(vector.Multiply(normal, lhsCircle.Radius), vector.Multiply(normal, -0.5*penetration))))

	return normal, penetration, points
}

func circleToConvex(l, r *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	lCircle := l.Shape.(*circle.Circle)
	rConvex := r.Shape.(*convex.Convex)

	minPenetration := math.MaxFloat64
	var minPenEdge convex.Edge
	for _, edge := range rConvex.Edges() {
		edgeNormal := r.Rotation().RotateVector(edge.Normal)
		edgeStart := r.Rotation().RotateVector(edge.Start)

		pen := -vector.Dot(edgeNormal, vector.Subtract(l.Position(), vector.Add(r.Position(), edgeStart)))

		if pen < -lCircle.Radius {
			return vector.ZERO(), 0, nil
		}

		if pen < minPenetration {
			minPenetration = pen
			minPenEdge = edge
		}
	}

	edgeStart := vector.Add(r.Position(), r.Rotation().RotateVector(minPenEdge.Start))
	edgeEnd := vector.Add(r.Position(), r.Rotation().RotateVector(minPenEdge.End))
	edgeNormal := r.Rotation().RotateVector(minPenEdge.Normal)

	p := closest_point.LineSegmentToPoint(l.Position(), edgeStart, edgeEnd)
	if 0 < vector.Dot(edgeNormal, vector.Subtract(p, l.Position())) {
		normal = vector.Subtract(l.Position(), p)

		penetration = lCircle.Radius + vector.Subtract(l.Position(), p).Size()
	} else {
		normal = vector.Subtract(p, l.Position())
		if lCircle.Radius < normal.Size() {
			return vector.ZERO(), 0, nil
		}

		penetration = lCircle.Radius - vector.Subtract(p, l.Position()).Size()
	}

	normal.Normalize()
	points = append(points, vector.Add(p, vector.Multiply(normal, 0.5*penetration)))

	return normal, penetration, points
}

func convexToConvex(l, r *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
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
		normal = vector.Invert(rNormal)
		penetration = rPenetration
		points = append(points, rPoint)
	}

	return normal, penetration, points
}

func findAxisLeastPenetration(l, r *convex.Convex, lPos, rPos vector.Vector, lRot, rRot rotator.Rotator) (minPenetration float64, bestNormal vector.Vector, bestPoint vector.Vector) {
	minPenetration = math.MaxFloat64

	for _, edge := range l.Edges() {
		normal := lRot.RotateVector(edge.Normal)
		s := r.Support(vector.Invert(normal), rRot)

		v := vector.Add(lRot.RotateVector(edge.Start), lPos)
		v.Subtract(rPos)

		penetration := -vector.Dot(normal, vector.Subtract(s, v))

		if penetration < minPenetration {
			bestNormal = normal
			minPenetration = penetration
			bestPoint = vector.Add(vector.Add(s, rPos), vector.Multiply(bestNormal, penetration*0.5))
		}
	}

	return minPenetration, bestNormal, bestPoint
}
