package composite

import "gitlab.com/legionary/legionary/core/ai"

// Selector Execute their children from left to right.
// Stop when one of their children succeeds, then selector succeeds.
// If all the selector's children fail, then the selector fails.
type Selector struct {
	ai.Composite

	index int
}

// Init inits.
func (s *Selector) Init() {
	s.index = 0
}

// Update updates.
func (s *Selector) Update(delta float64) ai.State {
	s.Composite.Update(delta)

	if len(s.Children()) == 0 {
		return ai.Success
	}

	for {
		node := s.Children()[s.index]

		state := ai.Update(node, delta)
		if state != ai.Failure {
			return state
		}

		s.index++
		if s.index == len(s.Children()) {
			return ai.Failure
		}
	}
}
