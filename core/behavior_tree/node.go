package behavior_tree

// Node is base struct for all node.
type Node struct {
	state State
}

// Init initializes node.
func (node *Node) Init() {
}

// Tick ticks nodes.
func (node *Node) Tick() State {
	return Invalid
}

// SetState sets state.
func (node *Node) SetState(state State) {
	node.state = state
}

// State returns state.
func (node *Node) State() State {
	return node.state
}
