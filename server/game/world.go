package game

import (
	"time"

	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
)

// World represents game world.
type World struct {
	graph      *graph.Graph
	centerNode *Node
	actors     map[int64]*Actor
}

// NewWorld create new world.
func NewWorld() *World {
	g, centerNode := newGraph()
	world := &World{
		g,
		centerNode,
		map[int64]*Actor{},
	}

	return world
}

// Tick ticks world.
func (world *World) Tick(now time.Time) {
	for _, n := range world.graph.Nodes() {
		node := n.(*Node)
		node.Tick(now)
	}
}

// ForEachNode executes function to all nodes.
func (world *World) ForEachNode(f func(n *Node)) {
	for _, n := range world.graph.Nodes() {
		node := n.(*Node)
		f(node)
	}
}

// ForEachActor executes function to all actors.
func (world *World) ForEachActor(f func(a *Actor)) {
	for _, a := range world.actors {
		f(a)
	}
}

// GetCenterNodeId returns center node ID.
func (world *World) GetCenterNodeId() int64 {
	return world.centerNode.ID()
}

func newGraph() (g *graph.Graph, centerNode *Node) {
	offset := 10.0
	center := 30

	g = graph.NewGraph()
	for x := 0; x < (center*2)+1; x++ {
		for y := 0; y < (center*2)+1; y++ {
			node := NewNode(id_generator.Generate(), vector.Vector{X: float64(x) * offset, Y: float64(y) * offset})
			if center == x && center == y {
				centerNode = node
			}
			g.AddNode(node)
		}
	}

	for _, startNode := range g.Nodes() {
		for _, endNode := range g.Nodes() {
			if startNode.Position().Sub(endNode.Position()).Size() <= offset {
				g.AddEdge(startNode.ID(), endNode.ID())
			}
		}
	}

	return g, centerNode
}
