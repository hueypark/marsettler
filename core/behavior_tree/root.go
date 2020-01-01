package behavior_tree

type root struct {
	Node

	child INode
}

func (r *root) AddChild(node INode) {
	r.child = node
}

func (r *root) Wireframe() string {
	return r.child.Wireframe()
}
