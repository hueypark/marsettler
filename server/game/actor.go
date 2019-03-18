package game

import (
	"log"

	"github.com/hueypark/marsettler/core/ai"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
)

// Actor represent actor.
type Actor struct {
	id           int64
	node         *Node
	behaviorTree ai.BehaviorTree
}

// NewActor creates new actor.
func NewActor(node *Node) *Actor {
	actor := &Actor{
		id:   id_generator.Generate(),
		node: node,
	}

	return actor
}

// Position returns position.
func (actor *Actor) Position() vector.Vector {
	return actor.node.Position()
}

// Tick ticks actor.
func (actor *Actor) Tick() {
	log.Println(actor)
}
