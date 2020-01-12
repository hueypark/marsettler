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

// Update updates INode.
func (r *root) Tick(delta float64) State {
	if r.child.State() != Running {
		r.child.Init()
	}

	return r.SetState(r.child.Tick(delta))
}

func (r *root) Wireframe() string {
	return r.child.Wireframe()
}
