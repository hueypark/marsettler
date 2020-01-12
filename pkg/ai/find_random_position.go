package ai

import (
	"fmt"
	"strings"

	"github.com/hueypark/marsettler/core/behavior_tree"
)

type FindRandomPosition struct {
	behavior_tree.Node

	actor       actor
	blackboard  *behavior_tree.Blackboard
	positionKey behavior_tree.BlackboardKey
}

func NewFindRandomPosition(actor actor, blackboard *behavior_tree.Blackboard, params string) *FindRandomPosition {
	positionKey := strings.ReplaceAll(params, " ", "")

	return &FindRandomPosition{
		actor:       actor,
		blackboard:  blackboard,
		positionKey: Key(positionKey),
	}
}

func (node *FindRandomPosition) Init() {
}

func (node *FindRandomPosition) Tick(delta float64) behavior_tree.State {
	position := node.actor.FindRandomPosition()
	node.blackboard.Set(node.positionKey, &position)

	return node.SetState(behavior_tree.Success)
}

func (node *FindRandomPosition) Wireframe() string {
	return fmt.Sprintf("FindRandomPosition: %v", String(node.positionKey))
}
