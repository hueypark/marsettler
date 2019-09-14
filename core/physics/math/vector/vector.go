package vector

import (
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func ZERO() Vector {
	return Vector{0, 0}
}

func (v *Vector) Invert() {
	v.X = -v.X
	v.Y = -v.Y
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.SquareMagnitude())
}

func (v *Vector) SquareMagnitude() float64 {
	return (v.X * v.X) + (v.Y * v.Y)
}

func (v *Vector) Normalize() {
	l := v.Magnitude()
	if l > 0 {
		v.X /= l
		v.Y /= l
	}
}

func (v *Vector) Multiply(val float64) {
	v.X *= val
	v.Y *= val
}

func Multiply(v Vector, val float64) Vector {
	return Vector{v.X * val, v.Y * val}
}

func (v *Vector) Add(o Vector) {
	v.X += o.X
	v.Y += o.Y
}

func Add(lhs, rhs Vector) Vector {
	return Vector{lhs.X + rhs.X, lhs.Y + rhs.Y}
}

func (v *Vector) AddScaledVector(o Vector, scale float64) {
	v.X += o.X * scale
	v.Y += o.Y * scale
}

func (v *Vector) Subtract(o Vector) {
	v.X -= o.X
	v.Y -= o.Y
}

func Subtract(lhs, rhs Vector) Vector {
	return Vector{lhs.X - rhs.X, lhs.Y - rhs.Y}
}

func (v *Vector) Clear() {
	v.X = 0
	v.Y = 0
}

func Product(lhs, rhs Vector) Vector {
	return Vector{lhs.X * rhs.X, lhs.Y * rhs.Y}
}

func Dot(lhs, rhs Vector) float64 {
	return (lhs.X * rhs.X) + (lhs.Y * rhs.Y)
}

func Cross(lhs, rhs Vector) float64 {
	return (lhs.X * rhs.Y) - (lhs.Y * rhs.X)
}

func (v Vector) OnTheRight(o Vector) bool {
	return Cross(v, o) < 0
}

func (v Vector) Size() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) SizeSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func Invert(v Vector) Vector {
	return Vector{-v.X, -v.Y}
}
