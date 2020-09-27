package physics

import "github.com/hueypark/marsettler/pkg/internal/math2d"

const RESTITUTION = 0.5

type Contact struct {
	lhs         *Body
	rhs         *Body
	normal      *math2d.Vector // lhs to rhs
	penetration float64
	points      []*math2d.Vector
}

func New(lhs, rhs *Body) *Contact {
	return &Contact{
		lhs:    lhs,
		rhs:    rhs,
		normal: &math2d.Vector{}}
}

func (c *Contact) SolveCollision() {
	c.addImpulse()
	c.solvePenetration()
}

func (c *Contact) Points() []*math2d.Vector {
	return c.points
}

func (c *Contact) Normal() *math2d.Vector {
	return c.normal.Clone()
}

func (c *Contact) Penetration() float64 {
	return c.penetration
}

func (c *Contact) addImpulse() {
	for _, p := range c.points {
		relativeVelocity := math2d.Sub(c.rhs.Velocity, c.lhs.Velocity)

		velAlongNormal := math2d.Dot(relativeVelocity, c.normal)
		if velAlongNormal > 0 {
			return
		}

		contactVelocity := velAlongNormal * -(1 + RESTITUTION)

		inverseMassSum := c.lhs.InverseMass() + c.rhs.InverseMass()

		impulse := math2d.Mul(c.normal, contactVelocity)
		impulse.Mul(1 / inverseMassSum)

		c.lhs.AddImpluse(math2d.Mul(impulse, -1), math2d.Sub(c.lhs.Position(), p))
		c.rhs.AddImpluse(impulse, math2d.Sub(c.rhs.Position(), p))
	}
}

func (c *Contact) solvePenetration() {
	if !c.lhs.Static() {
		c.lhs.SetPosition(math2d.Add(c.lhs.Position(), math2d.Mul(c.normal, c.penetration*-0.5)))
	}

	if !c.rhs.Static() {
		c.rhs.SetPosition(math2d.Add(c.rhs.Position(), math2d.Mul(c.normal, c.penetration*0.5)))
	}
}
