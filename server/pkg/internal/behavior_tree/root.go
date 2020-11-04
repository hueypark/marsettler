package behavior_tree

type root struct {
	Node

	child INode
}

func (r *root) AddChild(node INode) {
	r.child = node
}

func (r *root) Init() {

}

// Tick ticks root.
func (r *root) Tick() State {
	if r.child.State() != Running {
		r.child.Init()
	}

	return r.SetState(r.child.Tick())
}

func (r *root) Wireframe() string {
	return r.child.Wireframe()
}
