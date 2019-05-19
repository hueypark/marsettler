package shape

import "github.com/hueypark/marsettler/core/math/vector"

type Bullet interface {
	Position() vector.Vector
	SetPosition(vector.Vector)
}
