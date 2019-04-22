package ai

import (
	"log"

	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/server/game/ai/decorator"
	"github.com/hueypark/marsettler/server/game/ai/task"
	yaml "gopkg.in/yaml.v2"
)

// NewAI create new AI from string.
func NewAI(actor task.Actor, str string) *behavior_tree.BehaviorTree {
	data := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(str), data)
	if err != nil {
		log.Println(err)
		return nil
	}

	bt := behavior_tree.NewBehaviorTree()
	bt.SetRoot(newNode(actor, bt.Blackboard(), data))

	return bt
}

func newNode(actor task.Actor, blackboard *behavior_tree.Blackboard, data map[interface{}]interface{}) behavior_tree.INode {
	var node behavior_tree.INode
	name := data["Name"]
	switch name {
	case "BlackboardCondition":
		var blackboardConditions decorator.BlackboardConditions
		conditions := data["Conditions"].([]interface{})
		for _, condition := range conditions {
			conditionData := condition.(map[interface{}]interface{})
			name := conditionData["Name"]
			key := behavior_tree.BlackboardKey(conditionData["Key"].(int))
			switch name {
			case "HasKey":
				blackboardConditions = append(blackboardConditions, &decorator.BlackboardConditionHasKey{Key: key})
			case "NotHasKey":
				blackboardConditions = append(blackboardConditions, &decorator.BlackboardConditionNotHasKey{Key: key})
			}
		}

		decoratorBlackboardCondition := decorator.NewBlackboardCondition(blackboard, blackboardConditions...)

		childData := data["Child"].(map[interface{}]interface{})
		decoratorBlackboardCondition.SetChild(newNode(actor, blackboard, childData))

		node = decoratorBlackboardCondition
	case "CreateActor":
		actorID := data["ActorID"].(int)

		taskCreateActor := task.NewCreateActor(actor, actorID)

		node = taskCreateActor
	case "FindPath":
		taskFindPath := task.NewFindPath(blackboard, actor)

		node = taskFindPath
	case "MoveTo":
		moveWaitTime := data["MoveWaitTime"].(int)

		taskMoveTo := task.NewMoveTo(blackboard, actor, moveWaitTime)

		node = taskMoveTo
	case "Sequence":
		sequence := behavior_tree.NewSequence()

		children := data["Children"].([]interface{})
		for _, child := range children {
			childData := child.(map[interface{}]interface{})
			sequence.AddChild(newNode(actor, blackboard, childData))
		}

		node = sequence
	case "Wait":
		waitTick := data["WaitTick"].(int)

		taskWait := task.NewWait(waitTick)

		node = taskWait
	case nil:
		node = nil
	default:
		log.Println(name, data)
	}

	return node
}
