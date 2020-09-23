package math

import "math"

// Vector is 2 dimensional vector.
type Vector struct {
	X float64
	Y float64
}

// Add adds vector.
func (v *Vector) Add(other *Vector) {
	v.X += other.X
	v.Y += other.Y
}

func Add(lhs, rhs *Vector) *Vector {
	return &Vector{lhs.X + rhs.X, lhs.Y + rhs.Y}
}

func (v *Vector) AddScaledVector(other *Vector, scale float64) {
	v.X += other.X * scale
	v.Y += other.Y * scale
}

// Clone returns clone of vector.
func (v *Vector) Clone() *Vector {
	return &Vector{v.X, v.Y}
}

// Set sets vector.
func (v *Vector) Set(other *Vector) {
	v.X = other.X
	v.Y = other.Y
}

// Sub subtracts vector.
func (v *Vector) Sub(other *Vector) {
	v.X -= other.X
	v.Y -= other.Y
}

func Sub(lhs, rhs *Vector) *Vector {
	return &Vector{lhs.X - rhs.X, lhs.Y - rhs.Y}
}

func (v *Vector) Invert() {
	v.X = -v.X
	v.Y = -v.Y
}

func Invert(v *Vector) *Vector {
	return &Vector{-v.X, -v.Y}
}

// Normalize normalizes vector.
func (v *Vector) Normalize() {
	dot := Dot(v, v)

	if dot == 0 {
		return
	}

	v.Mul(1 / math.Sqrt(dot))
}

// Mul makes vector scalar product of v and other.
func (v *Vector) Mul(val float64) {
	v.X *= val
	v.Y *= val
}

func Mul(v *Vector, val float64) *Vector {
	return &Vector{v.X * val, v.Y * val}
}

func Dot(lhs, rhs *Vector) float64 {
	return (lhs.X * rhs.X) + (lhs.Y * rhs.Y)
}

func Cross(lhs, rhs *Vector) float64 {
	return (lhs.X * rhs.Y) - (lhs.Y * rhs.X)
}

func (v *Vector) Right() *Vector {
	return &Vector{v.Y, -v.X}
}

// Size returns size.
func (v *Vector) Size() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

// SizeSquare returns size squared.
func (v *Vector) SizeSquare() float64 {
	return (v.X * v.X) + (v.Y * v.Y)
}

func (v *Vector) Clear() {
	v.X = 0
	v.Y = 0
}

func OnTheRight(lhs, rhs *Vector) bool {
	return Cross(lhs, rhs) < 0
}

// Zero returns true if its zero.
func (v *Vector) Zero() bool {
	if v.X == 0 && v.Y == 0 {
		return true
	}

	return false
}
