package game

import (
	"log"

	"github.com/hueypark/marsettler/core/math/vector"
)

type Actor struct {
	id     int64
	nodeID int64
}

// Position returns position.
func (actor *Actor) Position() vector.Vector {
	if node, ok := nodes[actor.nodeID]; ok {
		return node.Position()
	}

	log.Println("node is nil", actor.nodeID)

	return vector.Zero()
}
