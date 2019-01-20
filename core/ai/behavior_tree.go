package ai

// BehaviorTree is tree for artificial intelligence.
type BehaviorTree struct {
	root INode
}

// INode is interface for Behavior Tree INode.
type INode interface {
	Init()
	Update(delta float64) State
	SetState(state State)
	State() State
}

// Init inits behavior tree.
func (bt *BehaviorTree) Init(n INode) {
	bt.root = n
}

// Update updates behavior tree.
func (bt *BehaviorTree) Update(delta float64) {
	Update(bt.root, delta)
}

// Update updates node.
func Update(n INode, delta float64) State {
	if n.State() != Running {
		n.Init()
		n.SetState(Running)
	}

	state := n.Update(delta)

	n.SetState(state)

	return state
}
