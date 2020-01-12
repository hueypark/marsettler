package ai

import (
	"fmt"
	"log"
	"strings"

	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/core/math/vector"
)

type MoveTo struct {
	behavior_tree.Node

	actor       actor
	blackboard  *behavior_tree.Blackboard
	positionKey behavior_tree.BlackboardKey
}

func NewMoveTo(actor actor, blackboard *behavior_tree.Blackboard, params string) *MoveTo {
	positionKey := strings.ReplaceAll(params, " ", "")

	return &MoveTo{
		actor:       actor,
		blackboard:  blackboard,
		positionKey: Key(positionKey),
	}
}

func (node *MoveTo) Init() {
}

func (node *MoveTo) Tick(delta float64) behavior_tree.State {
	position, err := node.getPosition()
	if err != nil {
		log.Println(err)
		return node.SetState(behavior_tree.Failure)
	}

	if !node.actor.MoveTo(*position) {
		return node.SetState(behavior_tree.Running)
	}

	return node.SetState(behavior_tree.Success)
}

func (node *MoveTo) Wireframe() string {
	return fmt.Sprintf("MoveTo: %v", String(node.positionKey))
}

func (node *MoveTo) getPosition() (*vector.Vector, error) {
	position := node.blackboard.Get(node.positionKey).(*vector.Vector)
	if position == nil {
		return nil, fmt.Errorf("position is nil")
	}

	return position, nil
}
