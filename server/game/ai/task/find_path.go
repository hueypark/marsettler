package task

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
)

// FindPath finds path.
type FindPath struct {
	behavior_tree.Node

	blackboard *behavior_tree.Blackboard
}

// NewFindPath creates new FindPath.
func NewFindPath(blackboard *behavior_tree.Blackboard) *FindPath {
	task := &FindPath{
		blackboard: blackboard,
	}

	return task
}

// Tick ticks task.
func (task *FindPath) Tick() behavior_tree.State {
	// TODO: Implement path find.
	// task.blackboard.SetInt64s(blackboard_key.Path, task.pathfinder())

	return behavior_tree.Success
}
