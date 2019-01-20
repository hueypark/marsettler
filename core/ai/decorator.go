package ai

type Decorator struct {
	Node

	child INode
}

func (decorator *Decorator) SetChild(n INode) {
	decorator.child = n
}

func (decorator *Decorator) Child() INode {
	return decorator.child
}

func (decorator *Decorator) Init() {
	decorator.Child().Init()
}
