package convex

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hueypark/marsettler/core/physics/math/rotator"
	"github.com/hueypark/marsettler/core/physics/math/vector"
)

func TestNewConvex(t *testing.T) {
	a := assert.New(t)

	vertices := []vector.Vector{{0, 0}, {100, 0}, {100, -10}, {150, 100}, {100, 200}, {0, 210}, {-50, 100}, {30, 30}, {75, 30}}
	hull := []vector.Vector{{-50, 100}, {0, 0}, {100, -10}, {150, 100}, {100, 200}, {0, 210}}

	c := New(vertices)

	a.Equal(hull, c.Hull())
}

func TestEdge(t *testing.T) {
	a := assert.New(t)

	vertices := []vector.Vector{
		{0, 0},
		{100, 0},
		{0, 100},
		{100, 100}}

	c := New(vertices)

	edges := c.Edges()
	for i, edge := range edges {
		nextIndex := i + 1
		if len(edges) <= nextIndex {
			nextIndex = 0
		}

		nextEdge := edges[nextIndex]
		a.True(vector.Subtract(nextEdge.End, nextEdge.Start).OnTheRight(vector.Subtract(edge.End, edge.Start)))
	}
}

func TestInHull(t *testing.T) {
	a := assert.New(t)

	vertices := []vector.Vector{
		{0, 0},
		{100, 0},
		{0, 100},
		{100, 100}}

	c := New(vertices)

	a.True(c.InHull(vector.ZERO(), rotator.ZERO(), vector.Vector{50, 50}))
	a.False(c.InHull(vector.ZERO(), rotator.ZERO(), vector.Vector{50, -50}))
}

func TestSupport(t *testing.T) {
	a := assert.New(t)

	c := New(
		[]vector.Vector{
			{0, 0},
			{100, 0},
			{0, 100},
			{100, 100}})

	a.Equal(c.Support(vector.Vector{1, 1}, rotator.ZERO()), vector.Vector{100, 100})
}
