package closest_point

import "github.com/hueypark/heavycannon/math/vector"

func LineSegmentToPoint(point, lineA, lineB vector.Vector) vector.Vector {
	ab := vector.Subtract(lineB, lineA)

	t := vector.Dot(vector.Subtract(point, lineA), ab) / vector.Dot(ab, ab)
	if t < 0.0 {
		t = 0.0
	} else if t > 1.0 {
		t = 1
	}

	return vector.Add(lineA, vector.Multiply(ab, t))
}
