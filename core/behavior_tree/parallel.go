package behavior_tree

import (
	"fmt"
	"strings"
)

// Parallel executes all children at once.
// It returns first-child result.
// If parallel does not have a child, it fails.
type Parallel struct {
	Composite
}

// NewParallel creates parallel node.
func NewParallel() *Parallel {
	return &Parallel{}
}

// Init inits parallel.
func (s *Parallel) Init() {
	for _, child := range s.children {
		child.Init()
	}
}

// Tick ticks parallel.
func (s *Parallel) Tick() State {
	state := Failure

	for i, child := range s.children {
		if i == 0 {
			state = child.Tick()
		} else {
			child.Tick()
		}

	}

	return state
}

// Wireframe make wireframe of node.
func (s *Parallel) Wireframe() string {
	str := fmt.Sprintln("Parallel")

	for _, child := range s.children {
		wfs := strings.Split(child.Wireframe(), "\n")
		for _, wf := range wfs {
			str += fmt.Sprintln("\t" + wf)
		}
	}

	return str
}
