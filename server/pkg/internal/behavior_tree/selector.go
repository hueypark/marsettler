package behavior_tree

// Selector Execute their children from left to right.
// Stop when one of their children succeeds, then selector succeeds.
// If all the selector's children fail, then the selector fails.
type Selector struct {
	Composite

	index int
}

// Init inits.
func (s *Selector) Init() {
	s.index = 0
}

// Tick ticks selector.
func (s *Selector) Tick(delta float64) State {
	s.Composite.Update(delta)

	if len(s.Children()) == 0 {
		return Success
	}

	for {
		node := s.Children()[s.index]

		state := node.Tick()
		if state != Failure {
			return state
		}

		s.index++
		if s.index == len(s.Children()) {
			return Failure
		}
	}
}
