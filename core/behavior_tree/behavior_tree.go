package behavior_tree

// BehaviorTree is tree for artificial intelligence.
type BehaviorTree struct {
	root       node
	blackboard *Blackboard
}

type node interface {
	Init()
	Tick() State
	SetState(state State)
	State() State
}

// NewBehaviorTree creates new BehaviorTree.
func NewBehaviorTree() *BehaviorTree {
	behaviorTree := &BehaviorTree{
		blackboard: NewBlackboard(),
	}

	return behaviorTree
}

// Blackboard returns blackboard.
func (bt *BehaviorTree) Blackboard() *Blackboard {
	return bt.blackboard
}

// SetRoot sets root node.
func (bt *BehaviorTree) SetRoot(root node) {
	bt.root = root
}

// Tick ticks behavior tree.
func (bt *BehaviorTree) Tick() {
	if bt.root == nil {
		return
	}

	Tick(bt.root)
}

// Update updates node.
func Tick(n node) State {
	if n.State() != Running {
		n.Init()
		n.SetState(Running)
	}

	state := n.Tick()

	n.SetState(state)

	return state
}
