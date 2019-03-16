package game

import (
	"time"

	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
)

// World represents game world.
type World struct {
	graph *graph.Graph
}

// NewWorld create new world.
func NewWorld() *World {
	world := &World{
		newGraph(),
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

func newGraph() *graph.Graph {
	offset := 10.0
	center := 30

	g := graph.NewGraph()
	for x := 0; x < (center*2)+1; x++ {
		for y := 0; y < (center*2)+1; y++ {
			g.AddNode(NewNode(id_generator.Generate(), vector.Vector{X: float64(x) * offset, Y: float64(y) * offset}))
		}
	}

	for _, startNode := range g.Nodes() {
		for _, endNode := range g.Nodes() {
			if startNode.Position().Sub(endNode.Position()).Size() <= offset {
				g.AddEdge(startNode.ID(), endNode.ID())
			}
		}
	}

	return g
}
