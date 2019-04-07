package ai

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/server/game/ai/task"
)

func NewNil(actor task.Actor) *behavior_tree.BehaviorTree {
	return nil
}
