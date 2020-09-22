package test

import "github.com/hueypark/marsettler/pkg/internal/math"

// Triangles a, b, and c must be ccw
func PointInTriangle(p, a, b, c math.Vector) bool {
	if math.Sub(p, a).OnTheRight(math.Sub(b, a)) == false {
		return false
	}

	if math.Sub(p, b).OnTheRight(math.Sub(c, b)) == false {
		return false
	}

	if math.Sub(p, c).OnTheRight(math.Sub(a, c)) == false {
		return false
	}

	return true
}
