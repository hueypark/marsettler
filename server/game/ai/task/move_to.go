package task

import (
	"log"

	"github.com/hueypark/marsettler/core/behavior_tree"
)

// MoveTo represents move to task.
type MoveTo struct {
	behavior_tree.Node
}

// Tick ticks task.
func (task *MoveTo) Tick() behavior_tree.State {
	log.Println("MoveTo")

	return behavior_tree.Success
}
