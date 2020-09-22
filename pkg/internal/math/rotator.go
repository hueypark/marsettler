package math

import "math"

type Rotator struct {
	radian float64
}

func NewRotator(radian float64) *Rotator {
	return &Rotator{radian}
}

func ZERO() Rotator {
	return Rotator{0}
}

func (r *Rotator) Add(radian float64) {
	r.radian += radian
}

func (r *Rotator) AddScaled(radian, scale float64) {
	r.radian += radian * scale
}

func (r *Rotator) Degree() float64 {
	return r.radian / math.Pi * 180.0
}

func (r *Rotator) Dir() Vector {
	return Vector{X: math.Cos(r.radian), Y: math.Sin(r.radian)}
}

func (r *Rotator) Radian() float64 {
	return r.radian
}

func (r Rotator) RotateVector(v Vector) Vector {
	return r.RotationMatrix().TransformVector(v)
}

func (r Rotator) RotationMatrix() (m Matrix) {
	c := math.Cos(r.radian)
	s := math.Sin(r.radian)

	m.M[0][0] = c
	m.M[0][1] = -s
	m.M[1][0] = s
	m.M[1][1] = c

	return m
}
