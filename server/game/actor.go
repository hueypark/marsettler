package game

import (
	"github.com/hueypark/marsettler/core/math/vector"
)

type Actor struct {
	id   int64
	node *Node
}

// Position returns position.
func (actor *Actor) Position() vector.Vector {
	return actor.node.Position()
}
