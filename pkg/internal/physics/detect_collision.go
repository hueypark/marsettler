package physics

import (
	"math"

	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/physics/closest_point"
	"github.com/hueypark/marsettler/pkg/internal/physics/shape"
)

func (c *Contact) DetectCollision() {
	lhsType := c.lhs.Shape.Type()
	rhsType := c.rhs.Shape.Type()

	switch lhsType {
	case shape.BulletType:
		switch rhsType {
		case shape.BulletType:
			break
		case shape.CircleType:
			c.normal, c.penetration, c.points = bulletToCircle(c.lhs, c.rhs)
			break
		case shape.ConvecType:
			c.normal, c.penetration, c.points = bulletToConvex(c.lhs, c.rhs)
			break
		}
		break
	case shape.CircleType:
		switch rhsType {
		case shape.BulletType:
			c.swap()
			c.normal, c.penetration, c.points = bulletToCircle(c.lhs, c.rhs)
			break
		case shape.CircleType:
			c.normal, c.penetration, c.points = circleToCircle(c.lhs, c.rhs)
			break
		case shape.ConvecType:
			c.normal, c.penetration, c.points = circleToConvex(c.lhs, c.rhs)
			break
		}
		break
	case shape.ConvecType:
		switch rhsType {
		case shape.BulletType:
			c.swap()
			c.normal, c.penetration, c.points = bulletToConvex(c.lhs, c.rhs)
		case shape.CircleType:
			c.swap()
			c.normal, c.penetration, c.points = circleToConvex(c.lhs, c.rhs)
			break
		case shape.ConvecType:
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
	lhs, rhs *Body,
) (normal *math2d.Vector, penetration float64, points []*math2d.Vector) {
	rhsCircle := rhs.Shape.(*shape.Circle)

	normal = math2d.Sub(rhs.Position(), lhs.Position())

	distanceSquared := normal.SizeSquare()

	if distanceSquared >= rhsCircle.Radius*rhsCircle.Radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	normal.Normalize()
	penetration = rhsCircle.Radius - distance
	points = append(points, math2d.Add(
		lhs.Position(),
		math2d.Mul(normal, -0.5*penetration)))

	return normal, penetration, points
}

func bulletToConvex(
	lhs, rhs *Body,
) (normal *math2d.Vector, penetration float64, points []*math2d.Vector) {
	rhsConvex := rhs.Shape.(*shape.Convex)

	penetration = math.MaxFloat64

	for _, edge := range rhsConvex.Edges() {
		worldStart := rhs.Rotation().RotateVector(edge.Start)
		worldStart.Add(rhs.Position())
		worldEnd := rhs.Rotation().RotateVector(edge.End)
		worldEnd.Add(rhs.Position())
		edgeVector := math2d.Sub(worldEnd, worldStart)
		pointVector := math2d.Sub(lhs.Position(), worldStart)

		if !math2d.OnTheRight(pointVector, edgeVector) {
			normal = &math2d.Vector{}
			penetration = 0
			return normal, penetration, points
		}

		perpendicular := &math2d.Vector{X: -edgeVector.Y, Y: edgeVector.X}
		perpendicular.Normalize()

		lhsVector := math2d.Sub(lhs.Position(), worldStart)

		proj := math2d.Mul(perpendicular, math2d.Dot(lhsVector, perpendicular))

		if proj.Size() < penetration {
			normal = perpendicular
			penetration = proj.Size()
		}
	}

	points = append(points, math2d.Add(
		lhs.Position(),
		math2d.Mul(normal, -0.5*penetration)))

	return normal, penetration, points
}

func circleToCircle(
	lhs, rhs *Body,
) (normal *math2d.Vector, penetration float64, points []*math2d.Vector) {
	lhsCircle := lhs.Shape.(*shape.Circle)
	rhsCircle := rhs.Shape.(*shape.Circle)

	normal = math2d.Sub(rhs.Position(), lhs.Position())

	distanceSquared := normal.SizeSquare()
	radius := lhsCircle.Radius + rhsCircle.Radius

	if distanceSquared >= radius*radius {
		return
	}

	distance := math.Sqrt(distanceSquared)

	normal.Normalize()
	penetration = radius - distance
	points = append(points, math2d.Add(
		lhs.Position(),
		math2d.Add(math2d.Mul(normal, lhsCircle.Radius), math2d.Mul(normal, -0.5*penetration))))

	return normal, penetration, points
}

func circleToConvex(
	l, r *Body,
) (normal *math2d.Vector, penetration float64, points []*math2d.Vector) {
	lCircle := l.Shape.(*shape.Circle)
	rConvex := r.Shape.(*shape.Convex)

	minPenetration := math.MaxFloat64
	var minPenEdge shape.Edge
	for _, edge := range rConvex.Edges() {
		edgeNormal := r.Rotation().RotateVector(edge.Normal)
		edgeStart := r.Rotation().RotateVector(edge.Start)

		pen := -math2d.Dot(edgeNormal, math2d.Sub(l.Position(), math2d.Add(r.Position(), edgeStart)))

		if pen < -lCircle.Radius {
			return &math2d.Vector{}, 0, nil
		}

		if pen < minPenetration {
			minPenetration = pen
			minPenEdge = edge
		}
	}

	edgeStart := math2d.Add(r.Position(), r.Rotation().RotateVector(minPenEdge.Start))
	edgeEnd := math2d.Add(r.Position(), r.Rotation().RotateVector(minPenEdge.End))
	edgeNormal := r.Rotation().RotateVector(minPenEdge.Normal)

	p := closest_point.LineSegmentToPoint(l.Position(), edgeStart, edgeEnd)
	if 0 < math2d.Dot(edgeNormal, math2d.Sub(p, l.Position())) {
		normal = math2d.Sub(l.Position(), p)

		penetration = lCircle.Radius + math2d.Sub(l.Position(), p).Size()
	} else {
		normal = math2d.Sub(p, l.Position())
		if lCircle.Radius < normal.Size() {
			return &math2d.Vector{}, 0, nil
		}

		penetration = lCircle.Radius - math2d.Sub(p, l.Position()).Size()
	}

	normal.Normalize()
	points = append(points, math2d.Add(p, math2d.Mul(normal, 0.5*penetration)))

	return normal, penetration, points
}

func convexToConvex(
	l, r *Body,
) (normal *math2d.Vector, penetration float64, points []*math2d.Vector) {
	lConvex := l.Shape.(*shape.Convex)
	rConvex := r.Shape.(*shape.Convex)

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
		normal = math2d.Invert(rNormal)
		penetration = rPenetration
		points = append(points, rPoint)
	}

	return normal, penetration, points
}

func findAxisLeastPenetration(
	l, r *shape.Convex, lPos, rPos *math2d.Vector, lRot, rRot math2d.Rotator,
) (minPenetration float64, bestNormal *math2d.Vector, bestPoint *math2d.Vector) {
	minPenetration = math.MaxFloat64

	for _, edge := range l.Edges() {
		normal := lRot.RotateVector(edge.Normal)
		s := r.Support(math2d.Invert(normal), rRot)

		v := math2d.Add(lRot.RotateVector(edge.Start), lPos)
		v.Sub(rPos)

		penetration := -math2d.Dot(normal, math2d.Sub(s, v))

		if penetration < minPenetration {
			bestNormal = normal
			minPenetration = penetration
			bestPoint = math2d.Add(math2d.Add(s, rPos), math2d.Mul(bestNormal, penetration*0.5))
		}
	}

	return minPenetration, bestNormal, bestPoint
}
