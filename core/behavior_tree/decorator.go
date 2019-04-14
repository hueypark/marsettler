package behavior_tree

type Decorator struct {
	Node

	child INode
}

func (decorator *Decorator) Init() {
	decorator.Child().Init()
}

func (decorator *Decorator) SetChild(n INode) {
	decorator.child = n
}

func (decorator *Decorator) Child() INode {
	return decorator.child
}
