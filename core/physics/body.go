package physics

import "github.com/hueypark/marsettler/core/math/vector"

// Body is rigidbody for physics engine.
type Body interface {
	ID() int64
	Shape() Shape
	Position() vector.Vector
	SetPosition(vector.Vector)
	OnCollision(other interface{}, normal vector.Vector, penetration float64)
}
