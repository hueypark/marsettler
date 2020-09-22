package test

import (
	"testing"

	"github.com/hueypark/marsettler/pkg/internal/math"
	"github.com/stretchr/testify/assert"
)

func TestPointInTriangle(t *testing.T) {
	a := assert.New(t)

	triangleVertexA := math.Vector{0, 0}
	triangleVertexB := math.Vector{100, 0}
	triangleVertexC := math.Vector{0, 100}

	a.True(
		PointInTriangle(
			math.Vector{30, 30},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))

	a.False(
		PointInTriangle(
			math.Vector{50, 200},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))

	a.False(
		PointInTriangle(
			math.Vector{0, -10},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))

	a.False(
		PointInTriangle(
			math.Vector{-10, 0},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))
}
