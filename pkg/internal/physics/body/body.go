package body

import "github.com/hueypark/marsettler/pkg/internal/math"

type Body struct {
	id              int64
	position        *math.Vector
	rotation        math.Rotator
	Velocity        *math.Vector
	angularVelocity float64
	Shape           shape
	mass            float64
	inverseMass     float64
	inverseInertia  float64
	forceSum        *math.Vector
}

type shape interface {
	Type() Shape
}

func New(id int64, position math.Vector) *Body {
	r := Body{
		id:       id,
		position: &math.Vector{X: position.X, Y: position.Y},
		Velocity: &math.Vector{},
		forceSum: &math.Vector{},
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

func (r *Body) Position() *math.Vector {
	return &math.Vector{X: r.position.X, Y: r.position.Y}
}

func (r *Body) SetPosition(position *math.Vector) {
	r.position.Set(position)
}

func (r *Body) SetVelocity(velovity *math.Vector) {
	r.Velocity.Set(velovity)
}

func (r *Body) Rotation() math.Rotator {
	return r.rotation
}

func (r *Body) SetRotation(rotation math.Rotator) {
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

	var acceleration = &math.Vector{}
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

func (r *Body) SetShape(s shape) {
	r.Shape = s
}

func (r *Body) AddForce(force *math.Vector) {
	r.forceSum.Add(force)
}

func (r *Body) AddImpluse(impulse, contact *math.Vector) {
	r.Velocity.AddScaledVector(impulse, r.inverseMass)
	r.angularVelocity -= math.Cross(contact, impulse) * r.inverseInertia * 0.01
}
