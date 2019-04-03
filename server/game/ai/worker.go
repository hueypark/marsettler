package ai

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/server/game/ai/blackboard_key"
	"github.com/hueypark/marsettler/server/game/ai/decorator"
	"github.com/hueypark/marsettler/server/game/ai/task"
)

// NewWorker creates new worker.
func NewWorker() *behavior_tree.BehaviorTree {
	worker := behavior_tree.NewBehaviorTree()

	findAndMove := &behavior_tree.Sequence{}
	hasNotPath := decorator.NewBlackboard(worker.Blackboard(), &decorator.BlackboardConditionNotHasKey{Key: blackboard_key.Path})
	hasNotPath.SetChild(task.NewFindPath(worker.Blackboard()))
	findAndMove.AddChild(hasNotPath)
	findAndMove.AddChild(task.NewMoveTo(worker.Blackboard(), 60))

	worker.SetRoot(findAndMove)

	return worker
}
