package decorator

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
)

// BlackboardCondition is blackboard based decorator.
type BlackboardCondition struct {
	behavior_tree.Decorator

	blackboard *behavior_tree.Blackboard
	conditions []blackboardCondition
}

// NewBlackboard creates new blackboard.
func NewBlackboard(blackboard *behavior_tree.Blackboard, conditions ...blackboardCondition) *BlackboardCondition {
	decorator := &BlackboardCondition{
		blackboard: blackboard,
		conditions: conditions,
	}

	return decorator
}

// Tick ticks task.
func (decorator *BlackboardCondition) Tick() behavior_tree.State {
	for _, condition := range decorator.conditions {
		if !condition.valid(decorator.blackboard) {
			return behavior_tree.Failure
		}
	}

	return decorator.Child().Tick()
}

func (decorator *BlackboardCondition) MarshalYAML() (interface{}, error) {
	type Blackboard struct {
		Conditions []blackboardCondition `yaml:"Conditions"`
		Child      behavior_tree.INode   `yaml:"Child"`
	}

	return struct {
		Blackboard Blackboard `yaml:"BlackboardCondition"`
	}{
		Blackboard: Blackboard{decorator.conditions, decorator.Child()},
	}, nil
}

// BlackboardConditionHasKey is a conditional expression that checks for the presence of a key.
type BlackboardConditionHasKey struct {
	Key behavior_tree.BlackboardKey `yaml:"Key"`
}

func (condition *BlackboardConditionHasKey) valid(blackboard *behavior_tree.Blackboard) bool {
	if blackboard.Get(condition.Key) == nil {
		return false
	}

	return true
}

func (condition *BlackboardConditionHasKey) MarshalYAML() (interface{}, error) {
	return struct {
		BlackboardConditionHasKey `yaml:"BlackboardConditionHasKey"`
	}{
		BlackboardConditionHasKey: *condition,
	}, nil
}

// BlackboardConditionNotHasKey is a conditional expression that checks for the presence of a key.
type BlackboardConditionNotHasKey struct {
	Key behavior_tree.BlackboardKey `yaml:"Key"`
}

func (condition *BlackboardConditionNotHasKey) valid(blackboard *behavior_tree.Blackboard) bool {
	if blackboard.Get(condition.Key) != nil {
		return false
	}

	return true
}

func (condition *BlackboardConditionNotHasKey) MarshalYAML() (interface{}, error) {
	return struct {
		BlackboardConditionNotHasKey `yaml:"BlackboardConditionNotHasKey"`
	}{
		BlackboardConditionNotHasKey: *condition,
	}, nil
}

type blackboardCondition interface {
	valid(blackboard *behavior_tree.Blackboard) bool
}
