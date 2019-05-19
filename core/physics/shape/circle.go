package shape

import "github.com/hueypark/marsettler/core/math/vector"

// Circle represents geometry circle.
type Circle interface {
	Position() vector.Vector
	SetPosition(vector.Vector)
	Radius() float64
}
