package task

import (
	"log"

	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/server/game/ai/blackboard_key"
)

// MoveTo represents move to task.
type MoveTo struct {
	behavior_tree.Node

	blackboard         *behavior_tree.Blackboard
	path               []int64
	moveWaitTime       int
	remainMoveWaitTime int
}

// NewMoveTo creates MoveTo task.
func NewMoveTo(blackboard *behavior_tree.Blackboard, moveWaitTime int) *MoveTo {
	task := &MoveTo{
		blackboard:   blackboard,
		moveWaitTime: moveWaitTime,
	}

	return task
}

// Init initializes task.
func (task *MoveTo) Init() {
	task.path = *task.blackboard.GetInt64s(blackboard_key.Path)
	task.remainMoveWaitTime = 0
}

// Tick ticks task.
func (task *MoveTo) Tick() behavior_tree.State {
	if len(task.path) == 0 {
		return behavior_tree.Success
	}

	task.remainMoveWaitTime--

	if 0 < task.remainMoveWaitTime {
		return behavior_tree.Running
	}

	task.remainMoveWaitTime += task.moveWaitTime

	nextNodeID := task.path[len(task.path)-1]
	task.path = task.path[:len(task.path)-1]

	// TODO: Implement move.
	//task.move(nextNodeID)
	log.Println(nextNodeID)

	return behavior_tree.Running
}
