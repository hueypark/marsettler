package ai

import "testing"

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
