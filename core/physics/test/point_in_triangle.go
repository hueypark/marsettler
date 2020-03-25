package test

import "github.com/hueypark/marsettler/core/math/vector"

// Triangles a, b, and c must be ccw
func PointInTriangle(p, a, b, c vector.Vector) bool {
	if vector.Sub(p, a).OnTheRight(vector.Sub(b, a)) == false {
		return false
	}

	if vector.Sub(p, b).OnTheRight(vector.Sub(c, b)) == false {
		return false
	}

	if vector.Sub(p, c).OnTheRight(vector.Sub(a, c)) == false {
		return false
	}

	return true
}
