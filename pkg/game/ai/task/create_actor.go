package task

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
)

// CreateActor represents create actor task.
type CreateActor struct {
	behavior_tree.Node

	actor   Actor
	actorID int
}

func NewCreateActor(actor Actor, actorID int) *CreateActor {
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

func (task *CreateActor) MarshalYAML() (interface{}, error) {
	return struct {
		Name    string `yaml:"Name"`
		ActorID int    `yaml:"ActorID"`
	}{
		Name:    "CreateActor",
		ActorID: task.actorID,
	}, nil
}