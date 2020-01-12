package ai

import "github.com/hueypark/marsettler/core/math/vector"

type actor interface {
	FindRandomPosition() vector.Vector
	MoveTo(position vector.Vector) (arrvie bool)
}
