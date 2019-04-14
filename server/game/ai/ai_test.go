package ai

import (
	"testing"

	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/server/game/ai/task"
	yaml "gopkg.in/yaml.v2"
)

func TestMarshal(t *testing.T) {
	var actor task.Actor

	tests := []struct {
		str string
		bt  *behavior_tree.BehaviorTree
	}{
		{
			`Sequence:
- Wait:
    waitTick: 60
    tick: 0
- CreateActor:
    actorID: 2
`,
			NewCityHall(actor),
		},
		{
			`Sequence:
			Blackboard:
				BlackboardConditionNotHasKey:
					Key: 0
			MoveTo:
				path: []
				moveWaitTime: 60
				remainMoveWaitTime: 0
			Wait:
				waitTick: 60
				tick: 0
			CreateActor:
				actorID: 3
		`,
			NewFairy(actor),
		},
		//{
		//	``,
		//	NewNil(actor),
		//},
		//{
		//	`Sequence:
		//	Blackboard:
		//		BlackboardConditionNotHasKey:
		//			Key: 0
		//	MoveTo:
		//		path: []
		//		moveWaitTime: 60
		//		remainMoveWaitTime: 0
		//`,
		//	NewWorker(actor),
		//},
	}

	for _, test := range tests {
		bytes, err := yaml.Marshal(test.bt)
		if err != nil {
			t.Errorf("marshal failed: %v", err)
		}

		str := string(bytes)

		if test.str != str {
			t.Errorf("expected: %v, got:%v", test.str, str)
		}
	}
}
