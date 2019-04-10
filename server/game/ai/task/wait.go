package task

import (
	"fmt"

	"github.com/hueypark/marsettler/core/behavior_tree"
)

// Wait represents wait task.
type Wait struct {
	behavior_tree.Node

	waitTick int
	tick     int
}

func NewWait(waitTick int) *Wait {
	task := &Wait{
		waitTick: waitTick,
		tick:     0,
	}

	return task
}

func (task *Wait) Init() {
	task.tick = 0
}

func (task *Wait) Tick() behavior_tree.State {
	task.tick++

	if task.tick < task.waitTick {
		return behavior_tree.Running
	}

	return behavior_tree.Success
}

func (task *Wait) Marshal() string {
	str := fmt.Sprintln("Wait:")
	str += behavior_tree.Indent("waitTick: %v", task.waitTick)
	str += behavior_tree.Indent("tick: %v", task.tick)

	return str
}
