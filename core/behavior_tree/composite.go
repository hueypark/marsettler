package behavior_tree

// Composite is composite base struct for behavior tree.
type Composite struct {
	Node

	children []*Node
	services []*Service
}

// AddChild add child to composite.
func (c *Composite) AddChild(n *Node) {
	c.children = append(c.children, n)
}

// Children return children fo composite.
func (c *Composite) Children() []*Node {
	return c.children
}

func (c *Composite) AddService(service *Service) {
	c.services = append(c.services, service)
}

func (c *Composite) Update(delta float64) {
	for _, service := range c.services {
		service.Update(delta)
	}
}
