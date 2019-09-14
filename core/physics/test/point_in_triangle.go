package test

import (
	"github.com/hueypark/marsettler/core/physics/math/vector"
)

// Triangles a, b, and c must be ccw
func PointInTriangle(p, a, b, c vector.Vector) bool {
	if vector.Subtract(p, a).OnTheRight(vector.Subtract(b, a)) == false {
		return false
	}

	if vector.Subtract(p, b).OnTheRight(vector.Subtract(c, b)) == false {
		return false
	}

	if vector.Subtract(p, c).OnTheRight(vector.Subtract(a, c)) == false {
		return false
	}

	return true
}
