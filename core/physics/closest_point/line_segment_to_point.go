package closest_point

import "github.com/hueypark/marsettler/core/math/vector"

func LineSegmentToPoint(point, lineA, lineB vector.Vector) vector.Vector {
	ab := vector.Sub(lineB, lineA)

	t := vector.Dot(vector.Sub(point, lineA), ab) / vector.Dot(ab, ab)
	if t < 0.0 {
		t = 0.0
	} else if t > 1.0 {
		t = 1
	}

	return vector.Add(lineA, vector.Mul(ab, t))
}
