package closest_point

import "github.com/hueypark/marsettler/server/pkg/internal/math2d"

func LineSegmentToPoint(point, lineA, lineB *math2d.Vector) *math2d.Vector {
	ab := math2d.Sub(lineB, lineA)

	t := math2d.Dot(math2d.Sub(point, lineA), ab) / math2d.Dot(ab, ab)
	if t < 0.0 {
		t = 0.0
	} else if t > 1.0 {
		t = 1
	}

	return math2d.Add(lineA, math2d.Mul(ab, t))
}
