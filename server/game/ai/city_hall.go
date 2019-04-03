package ai

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/server/game/ai/task"
)

// NewCityHall creates new city hall.
func NewCityHall() *behavior_tree.BehaviorTree {
	cityHall := behavior_tree.NewBehaviorTree()

	waitAndCreateActor := behavior_tree.NewSequence()
	waitAndCreateActor.AddChild(task.NewWait(60))
	waitAndCreateActor.AddChild(task.NewCreateActor(2))

	cityHall.SetRoot(waitAndCreateActor)

	return cityHall
}
