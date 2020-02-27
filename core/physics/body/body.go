package body

import (
	rotator2 "github.com/hueypark/marsettler/core/math/rotator"
	"github.com/hueypark/marsettler/core/math/vector"
)

type Body struct {
	id              int64
	position        vector.Vector
	rotation        rotator2.Rotator
	Velocity        vector.Vector
	angularVelocity float64
	Shape           shape
	mass            float64
	inverseMass     float64
	inverseInertia  float64
	forceSum        vector.Vector
}

type shape interface {
	Type() Shape
}

func New(id int64, position vector.Vector) *Body {
	r := Body{
		id:       id,
		position: position,
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

func (r *Body) Position() vector.Vector {
	return r.position
}

func (r *Body) SetPosition(position vector.Vector) {
	r.position = position
}

func (r *Body) SetVelocity(velovity vector.Vector) {
	r.Velocity = velovity
}

func (r *Body) Rotation() rotator2.Rotator {
	return r.rotation
}

func (r *Body) SetRotation(rotation rotator2.Rotator) {
	r.rotation = rotation
}

func (r *Body) SetAngularVelocity(degrees float64) {
	r.angularVelocity = degrees
}

func (r *Body) Tick(delta float64) {
	if r.inverseMass <= 0.0 {
		return
	}

	r.position = r.position.AddScaledVector(r.Velocity, delta)

	var acceleration = vector.Zero()
	acceleration = acceleration.AddScaledVector(r.forceSum, r.inverseMass)

	r.Velocity = r.Velocity.AddScaledVector(acceleration, delta)

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

func (r *Body) AddForce(force vector.Vector) {
	r.forceSum = r.forceSum.Add(force)
}

func (r *Body) AddImpluse(impulse, contact vector.Vector) {
	r.Velocity = r.Velocity.AddScaledVector(impulse, r.inverseMass)
	r.angularVelocity -= vector.Cross(contact, impulse) * r.inverseInertia * 0.01
}
