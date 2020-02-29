package graph

import "github.com/hueypark/marsettler/core/math/vector"

// Node represents graph node.
type Node interface {
	ID() int64
	Position() vector.Vector
}
