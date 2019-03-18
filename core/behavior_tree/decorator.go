package behavior_tree

type Decorator struct {
	Node

	child *Node
}

func (decorator *Decorator) Init() {
	decorator.Child().Init()
}

func (decorator *Decorator) SetChild(n *Node) {
	decorator.child = n
}

func (decorator *Decorator) Child() *Node {
	return decorator.child
}
