package composite

import "gitlab.com/legionary/legionary/core/ai"

// Sequence Execute their children from left to right.
// Stop when one of their children fails.
// If a child fails, then the sequence fails.
// If all the sequences's children succeed, then the sqeuence succeeds.
type Sequence struct {
	ai.Composite

	index int
}

// Init inits sqeuence.
func (s *Sequence) Init() {
	s.index = 0
}

// Update updates sequnce.
func (s *Sequence) Update(delta float64) ai.State {
	s.Composite.Update(delta)

	if len(s.Children()) == 0 {
		return ai.Success
	}

	for {
		node := s.Children()[s.index]

		state := ai.Update(node, delta)
		if state != ai.Success {
			return state
		}

		s.index++
		if s.index == len(s.Children()) {
			return ai.Success
		}
	}
}
