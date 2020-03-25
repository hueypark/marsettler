package closest_point

import (
	"testing"

	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/stretchr/testify/assert"
)

func TestLineSegmentToPoint(t *testing.T) {
	a := assert.New(t)

	point, lineA, lineB := vector.Vector{100, 100}, vector.Vector{0, 0}, vector.Vector{200, 0}
	cp := LineSegmentToPoint(point, lineA, lineB)
	a.Equal(cp, vector.Vector{100, 0})

	point, lineA, lineB = vector.Vector{300, 100}, vector.Vector{0, 0}, vector.Vector{200, 0}
	cp = LineSegmentToPoint(point, lineA, lineB)
	a.Equal(cp, vector.Vector{200, 0})

	point, lineA, lineB = vector.Vector{-100, 100}, vector.Vector{0, 0}, vector.Vector{200, 0}
	cp = LineSegmentToPoint(point, lineA, lineB)
	a.Equal(cp, vector.Vector{0, 0})
}
