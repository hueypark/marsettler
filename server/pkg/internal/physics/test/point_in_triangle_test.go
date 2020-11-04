package test

import (
	"testing"

	"github.com/hueypark/marsettler/server/pkg/internal/math2d"
	"github.com/stretchr/testify/assert"
)

func TestPointInTriangle(t *testing.T) {
	a := assert.New(t)

	triangleVertexA := math2d.Vector{0, 0}
	triangleVertexB := math2d.Vector{100, 0}
	triangleVertexC := math2d.Vector{0, 100}

	a.True(
		PointInTriangle(
			math2d.Vector{30, 30},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))

	a.False(
		PointInTriangle(
			math2d.Vector{50, 200},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))

	a.False(
		PointInTriangle(
			math2d.Vector{0, -10},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))

	a.False(
		PointInTriangle(
			math2d.Vector{-10, 0},
			triangleVertexA,
			triangleVertexB,
			triangleVertexC))
}
