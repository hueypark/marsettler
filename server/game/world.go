package game

import (
	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
)

// World represents game world.
type World struct {
	graph      *graph.Graph
	centerNode *Node
	nodes      map[int64]*Node
}

// NewWorld create new world.
func NewWorld() *World {
	g, centerNode := newGraph()
	world := &World{
		g,
		centerNode,
		map[int64]*Node{},
	}

	return world
}

// Tick ticks world.
func (world *World) Tick() {
	for _, n := range world.graph.Nodes() {
		node := n.(*Node)
		node.Tick()
	}
}

// ForEachNode executes function to all nodes.
func (world *World) ForEachNode(f func(n *Node)) {
	for _, n := range world.graph.Nodes() {
		node := n.(*Node)
		f(node)
	}
}

// GetCenterNode returns center node.
func (world *World) GetCenterNode() *Node {
	return world.centerNode
}

// GetNode returns node.
func (world *World) GetNode(id int64) *Node {
	return world.nodes[id]
}

func newGraph() (g *graph.Graph, centerNode *Node) {
	offset := 10.0
	center := 30
	nodes := map[int64]*Node{}

	g = graph.NewGraph()
	for x := 0; x < (center*2)+1; x++ {
		for y := 0; y < (center*2)+1; y++ {
			node := NewNode(id_generator.Generate(), vector.Vector{X: float64(x) * offset, Y: float64(y) * offset})
			if center == x && center == y {
				centerNode = node
			}
			nodes[node.ID()] = node
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
