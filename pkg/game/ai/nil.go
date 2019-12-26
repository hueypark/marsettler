package ai

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/pkg/game/ai/task"
)

func NewNil(actor task.Actor) *behavior_tree.BehaviorTree {
	nil := behavior_tree.NewBehaviorTree()

	return nil
}
