package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/asset"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/math/vector"
)

// Node represents major hub of the world.
type Node struct {
	id       int64
	position vector.Vector
	actors   map[int64]*Actor
}

// NewNode create new node.
func NewNode(id int64, position vector.Vector) *Node {
	node := &Node{
		id,
		position,
		map[int64]*Actor{},
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

// Render renders node in screen
func (node *Node) Render(screen *ebiten.Image) {
	renderer.Render(screen, asset.Grass, node.Position())
}

// NewActor creates new actor.
func (node *Node) NewActor() *Actor {
	actor := NewActor(node)

	node.actors[actor.ID()] = actor

	return actor
}

// AddActor adds actor
func (node *Node) AddActor(actor *Actor) {
	node.actors[actor.ID()] = actor
}

// DeleteActor deletes actor.
func (node *Node) DeleteActor(id int64) {
	delete(node.actors, id)
}

// ForEachActor executes function to all actors.
func (node *Node) ForEachActor(f func(a *Actor)) {
	for _, a := range node.actors {
		f(a)
	}
}
