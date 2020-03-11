package behavior_tree

// Node is base struct for all INode.
type Node struct {
	state State
}

// SetState sets state.
func (node *Node) SetState(state State) State {
	node.state = state

	return state
}

// State returns state.
func (node *Node) State() State {
	return node.state
}

// INode represents the interface of the node.
type INode interface {
	Init()
	SetState(state State) State
	State() State
	Tick() State
	Wireframe() string
}

type IAddChildNode interface {
	INode

	AddChild(INode)
}
