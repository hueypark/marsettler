package game

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
)

// Actor represent actor.
type Actor struct {
	id           int64
	node         *Node
	behaviorTree *behavior_tree.BehaviorTree
}

// NewActor creates new actor.
func NewActor(node *Node, behaviorTree *behavior_tree.BehaviorTree) *Actor {
	actor := &Actor{
		id_generator.Generate(),
		node,
		behaviorTree,
	}

	return actor
}

// Position returns position.
func (actor *Actor) Position() vector.Vector {
	return actor.node.Position()
}

// Tick ticks actor.
func (actor *Actor) Tick() {
	actor.behaviorTree.Tick()
}
