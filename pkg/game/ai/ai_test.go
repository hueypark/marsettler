package ai

import (
	"testing"

	"github.com/hueypark/marsettler/core/behavior_tree"
	yaml "gopkg.in/yaml.v2"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		str string
		bt  *behavior_tree.BehaviorTree
	}{
		{
			`Name: Sequence
Children:
- Name: Wait
  WaitTick: 60
  Tick: 0
- Name: CreateActor
  ActorID: 2
`,
			NewCityHall(nil),
		},
		{
			`Name: Sequence
Children:
- Name: BlackboardCondition
  Conditions:
  - Name: NotHasKey
    Key: 0
  Child:
    Name: FindPath
- Name: MoveTo
  Path: []
  MoveWaitTime: 60
  RemainMoveWaitTime: 0
- Name: Wait
  WaitTick: 60
  Tick: 0
- Name: CreateActor
  ActorID: 3
`,
			NewFairy(nil),
		},
		{
			`Name: Sequence
Children:
- Name: BlackboardCondition
  Conditions:
  - Name: NotHasKey
    Key: 0
  Child:
    Name: FindPath
- Name: MoveTo
  Path: []
  MoveWaitTime: 60
  RemainMoveWaitTime: 0
`,
			NewWorker(nil),
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

	for _, test := range tests {
		newAI := NewAI(nil, test.str)
		bytes, err := yaml.Marshal(newAI)
		if err != nil {
			t.Errorf("marshal failed: %v", err)
		}

		str := string(bytes)

		if test.str != str {
			t.Errorf("expected: %v, got:%v", test.str, str)
		}
	}
}
