package decorator

import (
	"fmt"

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

func (decorator *Blackboard) Marshal() string {
	str := fmt.Sprintln("Blackboard:")
	for _, condition := range decorator.conditions {
		str += behavior_tree.Indent(condition.marshal())
	}

	return str
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

func (condition *BlackboardConditionHasKey) marshal() string {
	str := fmt.Sprintln("BlackboardConditionHasKey:")
	str += behavior_tree.Indent("Key: %v", condition.Key)

	return str
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

func (condition *BlackboardConditionNotHasKey) marshal() string {
	str := fmt.Sprintln("BlackboardConditionNotHasKey:")
	str += behavior_tree.Indent("Key: %v", condition.Key)

	return str
}

type blackboardCondition interface {
	valid(blackboard *behavior_tree.Blackboard) bool
	marshal() string
}
