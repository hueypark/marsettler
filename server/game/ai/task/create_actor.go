package task

import (
	"fmt"

	"github.com/hueypark/marsettler/core/behavior_tree"
)

// CreateActor represents create actor task.
type CreateActor struct {
	behavior_tree.Node

	actor   Actor
	actorID int64
}

func NewCreateActor(actor Actor, actorID int64) *CreateActor {
	task := &CreateActor{
		actor:   actor,
		actorID: actorID,
	}

	return task
}

func (task *CreateActor) Init() {
}

func (task *CreateActor) Tick() behavior_tree.State {
	task.actor.CreateActor(task.actorID)

	return behavior_tree.Success
}

func (task *CreateActor) Marshal() string {
	str := fmt.Sprintln("CreateActor:")
	str += behavior_tree.Indent("actorID: %v", task.actorID)

	return str
}
