package ai

import (
	"reflect"
	"testing"
)

func TestNewBehaviorTree(t *testing.T) {
	origin := `Sequence
	FindRandomPosition: patrolPosition
	MoveTo: patrolPosition
	Wait: 1
`

	bt, err := NewBehaviorTree(nil, origin)
	if err != nil {
		t.Error(err)
	}

	result := bt.Wireframe()

	if origin != result {
		t.Errorf("origin and result are different. [origin: %v, result: %v]", origin, result)
	}
}

func TestParse(t *testing.T) {
	origin := `Sequence
	FindRandomPosition: patrolPosition
	MoveTo: patrolPosition
	Wait: 1
`
	expected := nodeData{
		str: "Sequence",
		children: []nodeData{
			{
				str: "	FindRandomPosition: patrolPosition",
			},
			{
				str: "	MoveTo: patrolPosition",
			},
			{
				str: "	Wait: 1",
			},
		},
	}

	result, err := parse(origin)
	if err != nil {
		t.Fatalf("Parse failed. [origin: %v]", origin)
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected and result are different. [expected: %v, result: %v]", expected, result)
	}
}
