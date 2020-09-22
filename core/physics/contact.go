package physics

import (
	"github.com/hueypark/marsettler/core/physics/body"
	"github.com/hueypark/marsettler/pkg/internal/math"
)

const RESTITUTION = 0.5

type Contact struct {
	lhs         *body.Body
	rhs         *body.Body
	normal      math.Vector // lhs to rhs
	penetration float64
	points      []math.Vector
}

func New(lhs, rhs *body.Body) *Contact {
	return &Contact{lhs: lhs, rhs: rhs}
}

func (c *Contact) SolveCollision() {
	c.addImpulse()
	c.solvePenetration()
}

func (c *Contact) Points() []math.Vector {
	return c.points
}

func (c *Contact) Normal() math.Vector {
	return c.normal
}

func (c *Contact) Penetration() float64 {
	return c.penetration
}

func (c *Contact) addImpulse() {
	for _, p := range c.points {
		relativeVelocity := math.Sub(c.rhs.Velocity, c.lhs.Velocity)

		velAlongNormal := math.Dot(relativeVelocity, c.normal)
		if velAlongNormal > 0 {
			return
		}

		contactVelocity := velAlongNormal * -(1 + RESTITUTION)

		inverseMassSum := c.lhs.InverseMass() + c.rhs.InverseMass()

		impulse := math.Mul(c.normal, contactVelocity)
		impulse.Mul(1 / inverseMassSum)

		c.lhs.AddImpluse(math.Mul(impulse, -1), math.Sub(c.lhs.Position(), p))
		c.rhs.AddImpluse(impulse, math.Sub(c.rhs.Position(), p))
	}
}

func (c *Contact) solvePenetration() {
	if !c.lhs.Static() {
		c.lhs.SetPosition(math.Add(c.lhs.Position(), math.Mul(c.normal, c.penetration*-0.5)))
	}

	if !c.rhs.Static() {
		c.rhs.SetPosition(math.Add(c.rhs.Position(), math.Mul(c.normal, c.penetration*0.5)))
	}
}
