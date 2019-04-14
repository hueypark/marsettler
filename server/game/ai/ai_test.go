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
  Children:
  - Wait:
      WaitTick: 60
      Tick: 0
  - CreateActor:
      ActorID: 2
`,
			NewCityHall(actor),
		},
		{
			`Sequence:
  Children:
  - BlackboardCondition:
      Conditions:
      - BlackboardConditionNotHasKey:
          Key: 0
      Child: FindPath
  - MoveTo:
      Path: []
      MoveWaitTime: 60
      RemainMoveWaitTime: 0
  - Wait:
      WaitTick: 60
      Tick: 0
  - CreateActor:
      ActorID: 3
`,
			NewFairy(actor),
		},
		{
			`Sequence:
  Children:
  - BlackboardCondition:
      Conditions:
      - BlackboardConditionNotHasKey:
          Key: 0
      Child: FindPath
  - MoveTo:
      Path: []
      MoveWaitTime: 60
      RemainMoveWaitTime: 0
`,
			NewWorker(actor),
		},
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
