package collision_check

import "github.com/hueypark/marsettler/core/math/vector"

func PointToAABB(point point, aabb aabb) bool {
	if point.Position().X < aabb.Left() {
		return false
	}

	if aabb.Right() < point.Position().X {
		return false
	}

	if aabb.Top() < point.Position().Y {
		return false
	}

	if point.Position().Y < aabb.Bottom() {
		return false
	}

	return true
}

type point interface {
	Position() vector.Vector
}

type aabb interface {
	Left() float64
	Right() float64
	Bottom() float64
	Top() float64
}
