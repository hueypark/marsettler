package decorator

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
)

// BlackboardCondition is blackboard based decorator.
type BlackboardCondition struct {
	behavior_tree.Decorator

	blackboard *behavior_tree.Blackboard
	conditions BlackboardConditions
}

type BlackboardConditions []blackboardCondition

// NewBlackboardCondition creates new blackboard condition.
func NewBlackboardCondition(
	blackboard *behavior_tree.Blackboard, conditions ...blackboardCondition,
) *BlackboardCondition {
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
	return struct {
		Name       string                `yaml:"Name"`
		Conditions []blackboardCondition `yaml:"Conditions"`
		Child      behavior_tree.INode   `yaml:"Child"`
	}{
		Name:       "BlackboardCondition",
		Conditions: decorator.conditions,
		Child:      decorator.Child(),
	}, nil
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

func (condition *BlackboardConditionHasKey) MarshalYAML() (interface{}, error) {
	return struct {
		Name string                      `yaml:"Name"`
		Key  behavior_tree.BlackboardKey `yaml:"Key"`
	}{
		Name: "HasKey",
		Key:  condition.Key,
	}, nil
}

// BlackboardConditionNotHasKey is a conditional expression that checks for the presence of a key.
type BlackboardConditionNotHasKey struct {
	Key behavior_tree.BlackboardKey `yaml:"NotHasKey"`
}

func (condition *BlackboardConditionNotHasKey) valid(blackboard *behavior_tree.Blackboard) bool {
	if blackboard.Get(condition.Key) != nil {
		return false
	}

	return true
}

func (condition *BlackboardConditionNotHasKey) MarshalYAML() (interface{}, error) {
	return struct {
		Name string                      `yaml:"Name"`
		Key  behavior_tree.BlackboardKey `yaml:"Key"`
	}{
		Name: "NotHasKey",
		Key:  condition.Key,
	}, nil
}

type blackboardCondition interface {
	valid(blackboard *behavior_tree.Blackboard) bool
}
