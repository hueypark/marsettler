package closest_point

import (
	"testing"

	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/stretchr/testify/assert"
)

func TestLineSegmentToPoint(t *testing.T) {
	a := assert.New(t)

	point, lineA, lineB := math2d.Vector{100, 100}, math2d.Vector{0, 0}, math2d.Vector{200, 0}
	cp := LineSegmentToPoint(point, lineA, lineB)
	a.Equal(cp, math2d.Vector{100, 0})

	point, lineA, lineB = math2d.Vector{300, 100}, math2d.Vector{0, 0}, math2d.Vector{200, 0}
	cp = LineSegmentToPoint(point, lineA, lineB)
	a.Equal(cp, math2d.Vector{200, 0})

	point, lineA, lineB = math2d.Vector{-100, 100}, math2d.Vector{0, 0}, math2d.Vector{200, 0}
	cp = LineSegmentToPoint(point, lineA, lineB)
	a.Equal(cp, math2d.Vector{0, 0})
}
