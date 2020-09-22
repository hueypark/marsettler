package test

import "github.com/hueypark/marsettler/pkg/internal/math"

func PointToAABB(p math.Vector, box aabb) bool {
	return box.Left() <= p.X && p.X <= box.Right() &&
		box.Bottom() <= p.Y && p.Y <= box.Top()
}

type aabb interface {
	Left() float64
	Right() float64
	Bottom() float64
	Top() float64
}
