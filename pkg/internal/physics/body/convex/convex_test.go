package convex

import (
	"testing"

	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/stretchr/testify/assert"
)

func TestNewConvex(t *testing.T) {
	a := assert.New(t)

	vertices := []math2d.Vector{{0, 0}, {100, 0}, {100, -10}, {150, 100}, {100, 200}, {0, 210}, {-50, 100}, {30, 30}, {75, 30}}
	hull := []math2d.Vector{{-50, 100}, {0, 0}, {100, -10}, {150, 100}, {100, 200}, {0, 210}}

	c := New(vertices)

	a.Equal(hull, c.Hull())
}

func TestEdge(t *testing.T) {
	a := assert.New(t)

	vertices := []math2d.Vector{
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
		a.True(math2d.Sub(nextEdge.End, nextEdge.Start).OnTheRight(math2d.Sub(edge.End, edge.Start)))
	}
}

func TestInHull(t *testing.T) {
	a := assert.New(t)

	vertices := []math2d.Vector{
		{0, 0},
		{100, 0},
		{0, 100},
		{100, 100}}

	c := New(vertices)

	a.True(c.InHull(math2d.Zero(), math2d.ZERO(), math2d.Vector{50, 50}))
	a.False(c.InHull(math2d.Zero(), math2d.ZERO(), math2d.Vector{50, -50}))
}

func TestSupport(t *testing.T) {
	a := assert.New(t)

	c := New(
		[]math2d.Vector{
			{0, 0},
			{100, 0},
			{0, 100},
			{100, 100}})

	a.Equal(c.Support(math2d.Vector{1, 1}, math2d.ZERO()), math2d.Vector{100, 100})
}