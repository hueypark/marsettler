package test

import "github.com/hueypark/marsettler/server/pkg/internal/math2d"

func PointToAABB(p math2d.Vector, box aabb) bool {
	return box.Left() <= p.X && p.X <= box.Right() &&
		box.Bottom() <= p.Y && p.Y <= box.Top()
}

type aabb interface {
	Left() float64
	Right() float64
	Bottom() float64
	Top() float64
}
