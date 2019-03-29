package collision_check

import "github.com/hueypark/marsettler/core/math/vector"

func PointToAABB(point vector.Vector, aabb aabb) bool {
	if point.X < aabb.Left() {
		return false
	}

	if aabb.Right() < point.X {
		return false
	}

	if aabb.Top() < point.Y {
		return false
	}

	if point.Y < aabb.Bottom() {
		return false
	}

	return true
}

type aabb interface {
	Left() float64
	Right() float64
	Bottom() float64
	Top() float64
}
