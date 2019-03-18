package game

import (
	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/math/vector"
)

// Node represents major hub of the world.
type Node struct {
	id       int64
	position vector.Vector
	actors   []*Actor
}

// NewNode create new node.
func NewNode(id int64, position vector.Vector) *Node {
	node := &Node{
		id,
		position,
		[]*Actor{},
	}

	return node
}

// ID returns id.
func (node Node) ID() int64 {
	return node.id
}

// Position returns position.
func (node Node) Position() vector.Vector {
	return node.position
}

// Distance returns distance between another node.
func (node Node) Distance(o graph.Node) float64 {
	return o.Position().Sub(node.Position()).Size()
}

// Tick ticks node.
func (node *Node) Tick() {
	for _, actor := range node.actors {
		actor.Tick()
	}
}

// NewActor creates new actor.
func (node *Node) NewActor() *Actor {
	actor := NewActor(node)

	node.actors = append(node.actors, actor)

	return actor
}

// ForEachActor executes function to all actors.
func (node *Node) ForEachActor(f func(a *Actor)) {
	for _, a := range node.actors {
		f(a)
	}
}
