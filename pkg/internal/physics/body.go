package physics

import (
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/physics/shape"
)

type Body struct {
	id              int64
	position        *math2d.Vector
	onSetPosition   func(position *math2d.Vector)
	rotation        math2d.Rotator
	Velocity        *math2d.Vector
	angularVelocity float64
	Shape           shape.Shape
	mass            float64
	inverseMass     float64
	inverseInertia  float64
	forceSum        *math2d.Vector
}

// NewBody creates new shape.
func NewBody(id int64, position *math2d.Vector, onSetPosition func(position *math2d.Vector)) *Body {
	r := Body{
		id:            id,
		position:      &math2d.Vector{X: position.X, Y: position.Y},
		onSetPosition: onSetPosition,
		Velocity:      &math2d.Vector{},
		forceSum:      &math2d.Vector{},
	}
	return &r
}

func (r *Body) Clear() {
	r.forceSum.Clear()
	r.Velocity.Clear()
}

func (r *Body) ID() int64 {
	return r.id
}

func (r *Body) Position() *math2d.Vector {
	return &math2d.Vector{X: r.position.X, Y: r.position.Y}
}

func (r *Body) SetPosition(position *math2d.Vector) {
	r.position.Set(position)
	if r.onSetPosition != nil {
		r.onSetPosition(position)
	}
}

func (r *Body) SetVelocity(velovity *math2d.Vector) {
	r.Velocity.Set(velovity)
}

func (r *Body) Rotation() math2d.Rotator {
	return r.rotation
}

func (r *Body) SetRotation(rotation math2d.Rotator) {
	r.rotation = rotation
}

func (r *Body) SetAngularVelocity(degrees float64) {
	r.angularVelocity = degrees
}

func (r *Body) Tick(delta float64) {
	if r.inverseMass <= 0.0 {
		return
	}

	r.position.AddScaledVector(r.Velocity, delta)

	var acceleration = &math2d.Vector{}
	acceleration.AddScaledVector(r.forceSum, r.inverseMass)

	r.Velocity.AddScaledVector(acceleration, delta)

	r.forceSum.Clear()

	r.rotation.AddScaled(r.angularVelocity, delta)
}

func (r *Body) SetMass(mass float64) {
	r.mass = mass
	r.inverseMass = 1.0 / mass
	r.inverseInertia = 1.0 / mass
}

func (r *Body) Mass() float64 {
	return r.mass
}

func (r *Body) InverseMass() float64 {
	return r.inverseMass
}

func (r *Body) SetStatic() {
	r.inverseMass = 0
}

func (r *Body) Static() bool {
	if r.inverseMass == 0 {
		return true
	}

	return false
}

func (r *Body) SetShape(s shape.Shape) {
	r.Shape = s
}

func (r *Body) AddForce(force *math2d.Vector) {
	r.forceSum.Add(force)
}

func (r *Body) AddImpluse(impulse, contact *math2d.Vector) {
	r.Velocity.AddScaledVector(impulse, r.inverseMass)
	r.angularVelocity -= math2d.Cross(contact, impulse) * r.inverseInertia * 0.01
}
