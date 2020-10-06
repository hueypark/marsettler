package behavior_tree

// BehaviorTree is tree for artificial intelligence.
type BehaviorTree struct {
	root       INode
	blackboard *Blackboard
}

// NewBehaviorTree creates new BehaviorTree.
func NewBehaviorTree() *BehaviorTree {
	behaviorTree := &BehaviorTree{
		root:       &root{},
		blackboard: NewBlackboard(),
	}

	return behaviorTree
}

// Blackboard returns blackboard.
func (bt *BehaviorTree) Blackboard() *Blackboard {
	return bt.blackboard
}

func (bt *BehaviorTree) Root() INode {
	return bt.root
}

// Tick ticks behavior tree.
func (bt *BehaviorTree) Tick() {
	if bt.root == nil {
		return
	}

	if bt.root.State() != Running {
		bt.root.Init()
	}

	bt.root.Tick()
}

func (bt *BehaviorTree) MarshalYAML() (interface{}, error) {
	return &bt.root, nil
}

// Wireframe returns wireframe as string.
func (bt *BehaviorTree) Wireframe() string {
	if bt.root == nil {
		return ""
	}

	return bt.root.Wireframe()
}
