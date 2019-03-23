package game

import (
	"log"

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
func NewActor(node *Node) *Actor {
	actor := &Actor{
		id:   id_generator.Generate(),
		node: node,
	}

	return actor
}

// SetBehaviorTree sets behavior tree.
func (actor *Actor) SetBehaviorTree(behaviorTree *behavior_tree.BehaviorTree) {
	actor.behaviorTree = behaviorTree
}

// ID returns id.
func (actor *Actor) ID() int64 {
	return actor.id
}

// Position returns position.
func (actor *Actor) Position() vector.Vector {
	return actor.node.Position()
}

// Tick ticks actor.
func (actor *Actor) Tick() {
	actor.behaviorTree.Tick()
}

// Move moves actor another node.
func (actor *Actor) Move(nodeID int64) {
	node := GetNode(nodeID)
	if node == nil {
		log.Println("node is nil", nodeID)
		return
	}

	actor.node.DeleteActor(actor.id)

	actor.node = node
	node.AddActor(actor)
}