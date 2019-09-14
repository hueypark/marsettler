package physics

import (
	"math"

	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics/body"
	"github.com/hueypark/marsettler/core/physics/body/circle"
	"github.com/hueypark/marsettler/core/physics/body/convex"
	"github.com/hueypark/marsettler/core/physics/closest_point"
	"github.com/hueypark/marsettler/core/physics/math/rotator"
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

func bulletToCircle(lhs, rhs *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	rhsCircle := rhs.Shape.(*circle.Circle)

	normal = vector.Sub(rhs.Position(), lhs.Position())

	distanceSquared := normal.SizeSquare()

	if distanceSquared >= rhsCircle.Radius*rhsCircle.Radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	normal.Normalize()
	penetration = rhsCircle.Radius - distance
	points = append(points, vector.Add(
		lhs.Position(),
		vector.Mul(normal, -0.5*penetration)))

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
		edgeVector := vector.Sub(worldEnd, worldStart)
		pointVector := vector.Sub(lhs.Position(), worldStart)

		if !pointVector.OnTheRight(edgeVector) {
			normal = vector.Vector{}
			penetration = 0
			return normal, penetration, points
		}

		perpendicular := vector.Vector{X: -edgeVector.Y, Y: edgeVector.X}
		perpendicular.Normalize()

		lhsVector := vector.Sub(lhs.Position(), worldStart)

		proj := vector.Mul(perpendicular, vector.Dot(lhsVector, perpendicular))

		if proj.Size() < penetration {
			normal = perpendicular
			penetration = proj.Size()
		}
	}

	points = append(points, vector.Add(
		lhs.Position(),
		vector.Mul(normal, -0.5*penetration)))

	return normal, penetration, points
}

func circleToCircle(lhs, rhs *body.Body) (normal vector.Vector, penetration float64, points []vector.Vector) {
	lhsCircle := lhs.Shape.(*circle.Circle)
	rhsCircle := rhs.Shape.(*circle.Circle)

	normal = vector.Sub(rhs.Position(), lhs.Position())

	distanceSquared := normal.SizeSquare()
	radius := lhsCircle.Radius + rhsCircle.Radius

	if distanceSquared >= radius*radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	normal.Normalize()
	penetration = radius - distance
	points = append(points, vector.Add(
		lhs.Position(),
		vector.Add(vector.Mul(normal, lhsCircle.Radius), vector.Mul(normal, -0.5*penetration))))

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

		pen := -vector.Dot(edgeNormal, vector.Sub(l.Position(), vector.Add(r.Position(), edgeStart)))

		if pen < -lCircle.Radius {
			return vector.Zero(), 0, nil
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
	if 0 < vector.Dot(edgeNormal, vector.Sub(p, l.Position())) {
		normal = vector.Sub(l.Position(), p)

		penetration = lCircle.Radius + vector.Sub(l.Position(), p).Size()
	} else {
		normal = vector.Sub(p, l.Position())
		if lCircle.Radius < normal.Size() {
			return vector.Zero(), 0, nil
		}

		penetration = lCircle.Radius - vector.Sub(p, l.Position()).Size()
	}

	normal.Normalize()
	points = append(points, vector.Add(p, vector.Mul(normal, 0.5*penetration)))

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
		v.Sub(rPos)

		penetration := -vector.Dot(normal, vector.Sub(s, v))

		if penetration < minPenetration {
			bestNormal = normal
			minPenetration = penetration
			bestPoint = vector.Add(vector.Add(s, rPos), vector.Mul(bestNormal, penetration*0.5))
		}
	}

	return minPenetration, bestNormal, bestPoint
}
