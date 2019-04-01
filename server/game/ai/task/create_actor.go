package task

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
)

// CreateActor represents create actor task.
type CreateActor struct {
	behavior_tree.Node

	createActorFunc func()
}

func NewCreateActor(createActorFunc func()) *CreateActor {
	task := &CreateActor{
		createActorFunc: createActorFunc,
	}

	return task
}

func (task *CreateActor) Init() {
}

func (task *CreateActor) Tick() behavior_tree.State {
	task.createActorFunc()

	return behavior_tree.Success
}
