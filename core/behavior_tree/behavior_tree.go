package behavior_tree

// BehaviorTree is tree for artificial intelligence.
type BehaviorTree struct {
	root       INode
	blackboard *Blackboard
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

// SetRoot sets root INode.
func (bt *BehaviorTree) SetRoot(root INode) {
	bt.root = root
}

// Tick ticks behavior tree.
func (bt *BehaviorTree) Tick() {
	if bt.root == nil {
		return
	}

	Tick(bt.root)
}

func (bt *BehaviorTree) MarshalYAML() (interface{}, error) {
	return &bt.root, nil
}

// Update updates INode.
func Tick(n INode) State {
	if n.State() != Running {
		n.Init()
		n.SetState(Running)
	}

	state := n.Tick()

	n.SetState(state)

	return state
}

// Wireframe returns wireframe as string.
func (bt *BehaviorTree) Wireframe() string {
	if bt.root == nil {
		return ""
	}

	return bt.root.Wireframe()
}
