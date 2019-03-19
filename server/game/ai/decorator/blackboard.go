package decorator

import (
	"log"

	"github.com/hueypark/marsettler/core/behavior_tree"
)

// Blackboard is blackboard based decorator.
type Blackboard struct {
	behavior_tree.Decorator
}

// Tick ticks task.
func (decorator *Blackboard) Tick() behavior_tree.State {
	log.Println("Blackboard")

	return decorator.Child().Tick()
}
