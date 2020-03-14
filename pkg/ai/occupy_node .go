package ai

import (
	"fmt"

	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/pkg/game"
)

// OccupyNode occupies node.
type OccupyNode struct {
	behavior_tree.Node

	actor *game.Actor
}

// NewOccupyNode creates occupy.
func NewOccupyNode(actor *game.Actor) *OccupyNode {
	return &OccupyNode{
		actor: actor,
	}
}

// Init implements behavior_tree.INode interface.
func (n *OccupyNode) Init() {
}

// Tick implements behavior_tree.INode interface.
func (n *OccupyNode) Tick() behavior_tree.State {
	n.actor.OccupyNode()

	return n.SetState(behavior_tree.Success)
}

// Wireframe implements behavior_tree.INode interface.
func (n *OccupyNode) Wireframe() string {
	return fmt.Sprintf("OccupyNode: %v")
}
