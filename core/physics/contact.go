package physics

import (
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics/shape"
)

type Contact struct {
	lhs         Body
	rhs         Body
	normal      vector.Vector
	penetration float64
}

func (contact *Contact) SolveCollision() {
	switch contact.lhs.Shape() {
	case Bullet:
		switch contact.rhs.Shape() {
		case Bullet:
		case Circle:
			lhs := contact.lhs.(shape.Bullet)
			rhs := contact.rhs.(shape.Circle)
			contact.normal, contact.penetration = bulletToCircle(lhs, rhs)
			contact.solveCollision()
		}
	case Circle:
		switch contact.rhs.Shape() {
		case Bullet:
			lhs := contact.lhs.(shape.Circle)
			rhs := contact.rhs.(shape.Bullet)
			contact.normal, contact.penetration = bulletToCircle(rhs, lhs)
			contact.normal = contact.normal.Invert()
			contact.penetration = -contact.penetration
			contact.solveCollision()
		case Circle:
			lhs := contact.lhs.(shape.Circle)
			rhs := contact.rhs.(shape.Circle)
			contact.normal, contact.penetration = circleToCricle(lhs, rhs)
			contact.solveCollision()
		}
	}
}

func (contact *Contact) solveCollision() {
	if 0 < contact.penetration {
		halfPen := contact.penetration * 0.5

		contact.lhs.OnCollision(contact.rhs, contact.normal, contact.penetration)
		contact.rhs.OnCollision(contact.lhs, contact.normal.Invert(), contact.penetration)

		contact.lhs.SetPosition(contact.lhs.Position().Add(contact.normal.Mul(halfPen)))
		contact.rhs.SetPosition(contact.rhs.Position().Add(contact.normal.Mul(-halfPen)))
	}
}

func circleToCricle(lhs, rhs shape.Circle) (normal vector.Vector, penetration float64) {
	diff := lhs.Position().Sub(rhs.Position())

	sumRadius := lhs.Radius() + rhs.Radius()

	normal = diff.Nomalize()

	penetration = sumRadius - diff.Size()

	return normal, penetration
}

func bulletToCircle(lhs shape.Bullet, rhs shape.Circle) (normal vector.Vector, penetration float64) {
	diff := lhs.Position().Sub(rhs.Position())

	normal = diff.Nomalize()

	penetration = rhs.Radius() - diff.Size()

	return normal, penetration
}
