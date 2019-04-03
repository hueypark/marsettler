package task

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
)

// CreateActor represents create actor task.
type CreateActor struct {
	behavior_tree.Node

	actorID int64
}

func NewCreateActor(actorID int64) *CreateActor {
	task := &CreateActor{
		actorID: actorID,
	}

	return task
}

func (task *CreateActor) Init() {
}

func (task *CreateActor) Tick() behavior_tree.State {
	// TODO: Implement create actor.
	// task.createActorFunc()

	return behavior_tree.Success
}
