package rotator

import (
	"math"

	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics/math/matrix"
	"github.com/hueypark/marsettler/core/physics/math/util"
)

type Rotator struct {
	Degrees float64
}

func ZERO() Rotator {
	return Rotator{0}
}

func (r *Rotator) Add(degrees float64) {
	r.Degrees += degrees
}

func (r *Rotator) AddScaled(degrees, scale float64) {
	r.Degrees += degrees * scale
}

func (r Rotator) RotateVector(v vector.Vector) vector.Vector {
	return r.RotationMatrix().TransformVector(v)
}

func (r Rotator) RotationMatrix() (m matrix.Matrix) {
	rad := util.DegToRad(r.Degrees)

	c := math.Cos(rad)
	s := math.Sin(rad)

	m.M[0][0] = c
	m.M[0][1] = -s
	m.M[1][0] = s
	m.M[1][1] = c

	return m
}
