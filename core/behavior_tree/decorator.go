package behavior_tree

type Decorator struct {
	Node

	child node
}

func (decorator *Decorator) Init() {
	decorator.Child().Init()
}

func (decorator *Decorator) SetChild(n node) {
	decorator.child = n
}

func (decorator *Decorator) Child() node {
	return decorator.child
}
