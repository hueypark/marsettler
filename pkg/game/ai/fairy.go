package ai

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/pkg/game/ai/blackboard_key"
	"github.com/hueypark/marsettler/pkg/game/ai/decorator"
	"github.com/hueypark/marsettler/pkg/game/ai/task"
)

func NewFairy(actor task.Actor) *behavior_tree.BehaviorTree {
	fairy := behavior_tree.NewBehaviorTree()

	sequence := behavior_tree.NewSequence()
	hasNotPath := decorator.NewBlackboardCondition(fairy.Blackboard(), &decorator.BlackboardConditionNotHasKey{Key: blackboard_key.Path})
	hasNotPath.SetChild(task.NewFindPath(fairy.Blackboard(), actor))
	sequence.AddChild(hasNotPath)
	sequence.AddChild(task.NewMoveTo(fairy.Blackboard(), actor, 60))
	sequence.AddChild(task.NewWait(60))
	sequence.AddChild(task.NewCreateActor(actor, 3))

	fairy.SetRoot(sequence)

	return fairy
}
