package ai

// Node is base struct for all node.
type Node struct {
	state State
}

// SetState sets state.
func (node *Node) SetState(state State) {
	node.state = state
}

// State returns state.
func (node Node) State() State {
	return node.state
}

func (node *Node) Init() {
}
