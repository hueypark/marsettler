package rotator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hueypark/heavycannon/math/vector"
)

func TestRotateVector(t *testing.T) {
	a := assert.New(t)

	v := vector.Vector{X: 0, Y: 1}
	r := Rotator{45}
	rv := r.RotateVector(v)

	a.InEpsilon(rv.X, 1, 0.1)
	a.InEpsilon(rv.Y, 0, 0.1)
}
