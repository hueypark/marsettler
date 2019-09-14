package vector

import "math"

// Vector is 2 dimensional vector.
type Vector struct {
	X float64
	Y float64
}

func Zero() Vector {
	return Vector{X: 0, Y: 0}
}

// Add returns add vector.
func (v Vector) Add(other Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y}
}

func Add(lhs, rhs Vector) Vector {
	return Vector{lhs.X + rhs.X, lhs.Y + rhs.Y}
}

func (v Vector) AddScaledVector(other Vector, scale float64) (result Vector) {
	result.X = v.X + (other.X * scale)
	result.Y = v.Y + (other.Y * scale)

	return result
}

// Sub returns subtract vector.
func (v Vector) Sub(other Vector) Vector {
	return Vector{v.X - other.X, v.Y - other.Y}
}

func Sub(lhs, rhs Vector) Vector {
	return Vector{lhs.X - rhs.X, lhs.Y - rhs.Y}
}

func (v Vector) Invert() Vector {
	return Vector{-v.X, -v.Y}
}

func Invert(v Vector) Vector {
	return Vector{-v.X, -v.Y}
}

// Normalize returns normalized vector.
func (v Vector) Normalize() Vector {
	dot := v.Dot(v)

	if dot == 0 {
		return Vector{0, 0}
	}
	return v.Mul(1 / math.Sqrt(dot))
}

// Mul returns the scalar product of v and other.
func (v Vector) Mul(val float64) Vector {
	return Vector{v.X * val, v.Y * val}
}

func Mul(v Vector, val float64) Vector {
	return Vector{v.X * val, v.Y * val}
}

// Dot returns the dot product of v and other.
func (v Vector) Dot(other Vector) float64 {
	return (v.X * other.X) + (v.Y * other.Y)
}

func Dot(lhs, rhs Vector) float64 {
	return (lhs.X * rhs.X) + (lhs.Y * rhs.Y)
}

// Cross returns cross product.
func (v Vector) Cross(other Vector) float64 {
	return (v.X * other.Y) - (v.Y * other.X)
}

func Cross(lhs, rhs Vector) float64 {
	return (lhs.X * rhs.Y) - (lhs.Y * rhs.X)
}

// Size returns size.
func (v Vector) Size() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

// SizeSquare returns size squared.
func (v Vector) SizeSquare() float64 {
	return (v.X * v.X) + (v.Y * v.Y)
}

func (v *Vector) Clear() {
	v.X = 0
	v.Y = 0
}

func (v Vector) OnTheRight(o Vector) bool {
	return Cross(v, o) < 0
}
