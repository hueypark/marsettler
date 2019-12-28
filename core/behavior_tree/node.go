package behavior_tree

// Node is base struct for all INode.
type Node struct {
	state State
}

// Init initializes INode.
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

// INode represents the interface of the node.
type INode interface {
	Init()
	Tick() State
	SetState(state State)
	State() State
	MarshalYAML() (interface{}, error)
	Wireframe() string
}
