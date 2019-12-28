package ai

import "testing"

func TestNewBehaviorTree(t *testing.T) {
	origin := `Sequence
	FindRandomPosition: patrolPosition
	MoveTo: patrolPosition
	Wait: 1
`

	bt := NewBehaviorTree(origin)

	result := bt.Wireframe()

	if origin != result {
		t.Errorf("Origin and result are different. [origin: %v, result: %s]", origin, result)
	}
}
