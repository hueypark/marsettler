package physics

import (
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics/body"
)

const RESTITUTION = 0.5

type Contact struct {
	lhs         *body.Body
	rhs         *body.Body
	normal      vector.Vector // lhs to rhs
	penetration float64
	points      []vector.Vector
}

func New(lhs, rhs *body.Body) *Contact {
	return &Contact{lhs: lhs, rhs: rhs}
}

func (c *Contact) SolveCollision() {
	c.addImpulse()
	c.solvePenetration()
}

func (c *Contact) Points() []vector.Vector {
	return c.points
}

func (c *Contact) Normal() vector.Vector {
	return c.normal
}

func (c *Contact) Penetration() float64 {
	return c.penetration
}

func (c *Contact) addImpulse() {
	for _, p := range c.points {
		relativeVelocity := vector.Sub(c.rhs.Velocity, c.lhs.Velocity)

		velAlongNormal := vector.Dot(relativeVelocity, c.normal)
		if velAlongNormal > 0 {
			return
		}

		contactVelocity := velAlongNormal * -(1 + RESTITUTION)

		inverseMassSum := c.lhs.InverseMass() + c.rhs.InverseMass()

		impulse := vector.Mul(c.normal, contactVelocity)
		impulse.Mul(1 / inverseMassSum)

		c.lhs.AddImpluse(vector.Mul(impulse, -1), vector.Sub(c.lhs.Position(), p))
		c.rhs.AddImpluse(impulse, vector.Sub(c.rhs.Position(), p))
	}
}

func (c *Contact) solvePenetration() {
	if !c.lhs.Static() {
		c.lhs.SetPosition(vector.Add(c.lhs.Position(), vector.Mul(c.normal, c.penetration*-0.5)))
	}

	if !c.rhs.Static() {
		c.rhs.SetPosition(vector.Add(c.rhs.Position(), vector.Mul(c.normal, c.penetration*0.5)))
	}
}
