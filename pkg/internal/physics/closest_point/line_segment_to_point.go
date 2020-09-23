package closest_point

import "github.com/hueypark/marsettler/pkg/internal/math"

func LineSegmentToPoint(point, lineA, lineB *math.Vector) *math.Vector {
	ab := math.Sub(lineB, lineA)

	t := math.Dot(math.Sub(point, lineA), ab) / math.Dot(ab, ab)
	if t < 0.0 {
		t = 0.0
	} else if t > 1.0 {
		t = 1
	}

	return math.Add(lineA, math.Mul(ab, t))
}
