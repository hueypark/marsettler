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

// Sub returns subtract vector.
func (v Vector) Sub(other Vector) Vector {
	return Vector{v.X - other.X, v.Y - other.Y}
}

func (v Vector) Invert() Vector {
	return Vector{-v.X, -v.Y}
}

// Nomalize returns nomalized vector.
func (v Vector) Nomalize() Vector {
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

// Dot returns the dot product of v and other.
func (v Vector) Dot(other Vector) float64 {
	return (v.X * other.X) + (v.Y * other.Y)
}

// Cross returns cross product.
func (v Vector) Cross(other Vector) float64 {
	return (v.X * other.Y) - (v.Y * other.X)
}

// Size returns size.
func (v Vector) Size() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

// SizeSquare returns size squared.
func (v Vector) SizeSquare() float64 {
	return (v.X * v.X) + (v.Y * v.Y)
}
