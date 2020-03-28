package rotator

import (
	"math"
	"testing"

	"github.com/hueypark/marsettler/core/math/vector"
)

func TestRotateVector(t *testing.T) {
	//a := assert.New(t)
	//
	//v := vector.Vector{X: 0, Y: 1}
	//r := Rotator{math.Pi}
	//rv := r.RotateVector(v)

	// TODO: These should be not in epsilon.
	//a.InEpsilon(rv.X, 0, 0.1)
	//a.InEpsilon(rv.Y, -1, 0.1)
}

func BenchmarkRotateVector(b *testing.B) {
	v := vector.Vector{X: 0, Y: 1}
	r := Rotator{math.Pi}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.RotateVector(v)
	}
}
