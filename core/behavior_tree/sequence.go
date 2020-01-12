package behavior_tree

import (
	"fmt"
	"strings"
)

// Sequence Execute their children from left to right.
// Stop when one of their children fails.
// If a child fails, then the sequence fails.
// If all the sequences's children succeed, then the sqeuence succeeds.
type Sequence struct {
	Composite

	index int
}

func NewSequence() *Sequence {
	return &Sequence{}
}

// Init inits sqeuence.
func (s *Sequence) Init() {
	s.index = 0

	for _, child := range s.children {
		child.Init()
	}
}

// Tick ticks sequnce.
func (s *Sequence) Tick(delta float64) State {
	childrenLen := len(s.children)
	if childrenLen == 0 {
		return s.SetState(Success)
	}

	for {
		node := s.children[s.index]

		state := node.Tick(delta)
		if state != Success {
			return s.SetState(state)
		}

		if childrenLen == s.index+1 {
			return s.SetState(Success)
		}

		s.index++
	}
}

func (s *Sequence) Wireframe() string {
	str := fmt.Sprintln("Sequence")

	for _, child := range s.children {
		wfs := strings.Split(child.Wireframe(), "\n")
		for _, wf := range wfs {
			str += fmt.Sprintln("\t" + wf)
		}
	}

	return str
}
