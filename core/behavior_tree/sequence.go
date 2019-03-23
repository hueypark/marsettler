package behavior_tree

// Sequence Execute their children from left to right.
// Stop when one of their children fails.
// If a child fails, then the sequence fails.
// If all the sequences's children succeed, then the sqeuence succeeds.
type Sequence struct {
	Composite

	index int
}

// Init inits sqeuence.
func (s *Sequence) Init() {
	s.index = 0
}

// Tick ticks sequnce.
func (s *Sequence) Tick() State {
	s.Composite.Tick()

	if len(s.Children()) == 0 {
		return Success
	}

	for {
		node := s.Children()[s.index]

		state := Tick(node)
		if state != Success {
			return state
		}

		s.index++
		if s.index == len(s.Children()) {
			return Success
		}
	}
}