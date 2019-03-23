package decorator

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
)

// Blackboard is blackboard based decorator.
type Blackboard struct {
	behavior_tree.Decorator

	blackboard *behavior_tree.Blackboard
	conditions []blackboardCondition
}

// NewBlackboard creates new blackboard.
func NewBlackboard(blackboard *behavior_tree.Blackboard, conditions ...blackboardCondition) *Blackboard {
	decorator := &Blackboard{
		blackboard: blackboard,
		conditions: conditions,
	}

	return decorator
}

// Tick ticks task.
func (decorator *Blackboard) Tick() behavior_tree.State {
	for _, condition := range decorator.conditions {
		if !condition.valid(decorator.blackboard) {
			return behavior_tree.Failure
		}
	}

	return decorator.Child().Tick()
}

// BlackboardConditionHasKey is a conditional expression that checks for the presence of a key.
type BlackboardConditionHasKey struct {
	Key behavior_tree.BlackboardKey
}

func (condition *BlackboardConditionHasKey) valid(blackboard *behavior_tree.Blackboard) bool {
	if blackboard.Get(condition.Key) == nil {
		return false
	}

	return true
}

// BlackboardConditionNotHasKey is a conditional expression that checks for the presence of a key.
type BlackboardConditionNotHasKey struct {
	Key behavior_tree.BlackboardKey
}

func (condition *BlackboardConditionNotHasKey) valid(blackboard *behavior_tree.Blackboard) bool {
	if blackboard.Get(condition.Key) != nil {
		return false
	}

	return true
}

type blackboardCondition interface {
	valid(blackboard *behavior_tree.Blackboard) bool
}