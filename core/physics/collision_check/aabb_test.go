package collision_check

import (
	"testing"

	"github.com/hueypark/marsettler/core/math/vector"
)

func TestPointToAABB(t *testing.T) {
	tests := []struct {
		lhs    testPoint
		rhs    testAABB
		result bool
	}{
		{
			testPoint{0, 0},
			testAABB{0, 10, 0, 10},
			true,
		},
		{
			testPoint{0, 0},
			testAABB{5, 10, 5, 10},
			false,
		},
	}

	for _, test := range tests {
		if PointToAABB(&test.lhs, &test.rhs) != test.result {
			t.Errorf("failed: %v", test)
		}
	}
}

type testPoint struct {
	x float64
	y float64
}

func (p *testPoint) Position() vector.Vector {
	return vector.Vector{X: p.x, Y: p.y}
}

type testAABB struct {
	left   float64
	right  float64
	bottom float64
	top    float64
}

func (aabb *testAABB) Left() float64 {
	return aabb.left
}

func (aabb *testAABB) Right() float64 {
	return aabb.right
}

func (aabb *testAABB) Bottom() float64 {
	return aabb.bottom
}

func (aabb *testAABB) Top() float64 {
	return aabb.top
}
