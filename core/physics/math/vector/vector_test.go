package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	a := assert.New(t)

	v1 := Vector{10, 0}
	v1.Normalize()

	nV1 := Vector{1, 0}
	a.Equal(v1, nV1)

	v2 := Vector{3, 4}
	v2.Normalize()

	nV2 := Vector{0.6, 0.8}
	a.Equal(v2, nV2)
}

func TestProduct(t *testing.T) {
	a := assert.New(t)

	v1 := Vector{1, 2}
	v2 := Vector{4, 5}
	v3 := Product(v1, v2)

	a.Equal(v3, Vector{4, 10})
}

func TestDot(t *testing.T) {
	a := assert.New(t)

	v1 := Vector{1, 2}
	v2 := Vector{4, 5}
	dot := Dot(v1, v2)

	a.Equal(dot, 14.0)
}

func TestOnRight(t *testing.T) {
	a := assert.New(t)

	v1 := Vector{0, 1}
	v2 := Vector{1, 1}

	a.True(v1.OnTheRight(v2))
}
