package task

import (
	"log"

	"github.com/hueypark/marsettler/core/behavior_tree"
)

// FindPath finds path.
type FindPath struct {
	behavior_tree.Node
}

// Tick ticks task.
func (task *FindPath) Tick() behavior_tree.State {
	log.Println("FindPath")

	return behavior_tree.Success
}
