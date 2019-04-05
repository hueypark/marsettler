package task

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/server/game/ai/blackboard_key"
)

// FindPath finds path.
type FindPath struct {
	behavior_tree.Node

	blackboard *behavior_tree.Blackboard
	actor      Actor
}

// NewFindPath creates new FindPath.
func NewFindPath(blackboard *behavior_tree.Blackboard, actor Actor) *FindPath {
	task := &FindPath{
		blackboard: blackboard,
		actor:      actor,
	}

	return task
}

// Tick ticks task.
func (task *FindPath) Tick() behavior_tree.State {
	task.blackboard.SetInt64s(blackboard_key.Path, task.actor.FindPath())

	return behavior_tree.Success
}
