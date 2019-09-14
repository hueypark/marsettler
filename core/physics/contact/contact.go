package contact

import (
	"github.com/hueypark/heavycannon/body"
	"github.com/hueypark/heavycannon/math/vector"
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
		relativeVelocity := vector.Subtract(c.rhs.Velocity, c.lhs.Velocity)

		velAlongNormal := vector.Dot(relativeVelocity, c.normal)
		if velAlongNormal > 0 {
			return
		}

		contactVelocity := velAlongNormal * -(1 + RESTITUTION)

		inverseMassSum := c.lhs.InverseMass() + c.rhs.InverseMass()

		impulse := vector.Multiply(c.normal, contactVelocity)
		impulse.Multiply(1 / inverseMassSum)

		c.lhs.AddImpluse(vector.Multiply(impulse, -1), vector.Subtract(c.lhs.Position(), p))
		c.rhs.AddImpluse(impulse, vector.Subtract(c.rhs.Position(), p))
	}
}

func (c *Contact) solvePenetration() {
	if !c.lhs.Static() {
		c.lhs.SetPosition(vector.Add(c.lhs.Position(), vector.Multiply(c.normal, c.penetration*-0.5)))
	}

	if !c.rhs.Static() {
		c.rhs.SetPosition(vector.Add(c.rhs.Position(), vector.Multiply(c.normal, c.penetration*0.5)))
	}
}
