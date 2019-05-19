package rotator

import (
	"math"

	"github.com/hueypark/marsettler/core/math/vector"
)

// Rotator represent roate ob object.
type Rotator struct {
	Radian float64
}

// Dir returns forward direction.
func (r Rotator) Dir() vector.Vector {
	return vector.Vector{X: math.Cos(r.Radian), Y: math.Sin(r.Radian)}
}

// Degree returns degree.
func (r Rotator) Degree() float64 {
	return r.Radian / math.Pi * 180.0
}

// Add adds radian.
func (r *Rotator) Add(radian float64) {
	r.Radian += radian
}
