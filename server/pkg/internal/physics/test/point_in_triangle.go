package test

import "github.com/hueypark/marsettler/server/pkg/internal/math2d"

// Triangles a, b, and c must be ccw
func PointInTriangle(p, a, b, c math2d.Vector) bool {
	if math2d.Sub(p, a).OnTheRight(math2d.Sub(b, a)) == false {
		return false
	}

	if math2d.Sub(p, b).OnTheRight(math2d.Sub(c, b)) == false {
		return false
	}

	if math2d.Sub(p, c).OnTheRight(math2d.Sub(a, c)) == false {
		return false
	}

	return true
}
