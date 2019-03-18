package behavior_tree

// BehaviorTree is tree for artificial intelligence.
type BehaviorTree struct {
	root *Node
}

// SetRoot sets root node.
func (bt *BehaviorTree) SetRoot(root *Node) {
	bt.root = root
}

// Tick ticks behavior tree.
func (bt *BehaviorTree) Tick() {
	Tick(bt.root)
}

// Update updates node.
func Tick(n *Node) State {
	if n.State() != Running {
		n.Init()
		n.SetState(Running)
	}

	state := n.Tick()

	n.SetState(state)

	return state
}
