package test

import (
	"testing"

	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/stretchr/testify/assert"
)

func TestPointInTriangle(t *testing.T) {
	a := assert.New(t)

	triangleVertexA := vector.Vector{0, 0}
	triangleVertexB := vector.Vector{100, 0}
	triangleVertexC := vector.Vector{0, 100}

	a.True(
		PointInTriangle(
			vector.Vector{30, 30},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))

	a.False(
		PointInTriangle(
			vector.Vector{50, 200},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))

	a.False(
		PointInTriangle(
			vector.Vector{0, -10},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))

	a.False(
		PointInTriangle(
			vector.Vector{-10, 0},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))
}
