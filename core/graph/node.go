package graph

import "gitlab.com/legionary/legionary/core/math/vector"

type Node interface {
	ID() int64
	Position() vector.Vector
	Len(node Node) float64
}
