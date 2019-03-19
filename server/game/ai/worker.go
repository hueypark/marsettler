package ai

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/server/game/ai/decorator"
	"github.com/hueypark/marsettler/server/game/ai/task"
)

// NewWorker creates new worker.
func NewWorker() *behavior_tree.BehaviorTree {
	worker := &behavior_tree.BehaviorTree{}

	findAndMove := &behavior_tree.Sequence{}
	hasNotPath := &decorator.Blackboard{}
	hasNotPath.SetChild(&task.FindPath{})
	findAndMove.AddChild(hasNotPath)
	findAndMove.AddChild(&task.MoveTo{})

	worker.SetRoot(findAndMove)

	return worker
}
