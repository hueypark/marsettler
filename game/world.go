package game

import (
	"time"

	"github.com/hueypark/marsettler/core/graph"
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

// Run runs world.
func (world *World) Run() {
	ticker := time.NewTicker(time.Second / 30)

	go func() {
		for {
			select {
			case <-ticker.C:
				world.Tick(time.Now())
			}
		}
	}()
}

// Tick ticks world.
func (world *World) Tick(now time.Time) {
	for _, n := range world.graph.Nodes() {
		node := n.(*Node)
		node.Tick(now)
	}
}

func newGraph() *graph.Graph {
	offset := 200.0

	g := graph.NewGraph()
	g.AddNode(NewNode(1, vector.Vector{X: -3 * offset, Y: -3 * offset}))
	g.AddNode(NewNode(2, vector.Vector{X: -1 * offset, Y: -3 * offset}))
	g.AddNode(NewNode(3, vector.Vector{X: +1 * offset, Y: -3 * offset}))
	g.AddNode(NewNode(4, vector.Vector{X: +3 * offset, Y: -3 * offset}))
	g.AddNode(NewNode(5, vector.Vector{X: -4 * offset, Y: -2 * offset}))
	g.AddNode(NewNode(6, vector.Vector{X: -2 * offset, Y: -2 * offset}))
	g.AddNode(NewNode(7, vector.Vector{X: +0 * offset, Y: -2 * offset}))
	g.AddNode(NewNode(8, vector.Vector{X: +2 * offset, Y: -2 * offset}))
	g.AddNode(NewNode(9, vector.Vector{X: +4 * offset, Y: -2 * offset}))
	g.AddNode(NewNode(10, vector.Vector{X: -5 * offset, Y: -1 * offset}))
	g.AddNode(NewNode(11, vector.Vector{X: -3 * offset, Y: -1 * offset}))
	g.AddNode(NewNode(12, vector.Vector{X: -1 * offset, Y: -1 * offset}))
	g.AddNode(NewNode(13, vector.Vector{X: +1 * offset, Y: -1 * offset}))
	g.AddNode(NewNode(14, vector.Vector{X: +3 * offset, Y: -1 * offset}))
	g.AddNode(NewNode(15, vector.Vector{X: +5 * offset, Y: -1 * offset}))
	g.AddNode(NewNode(16, vector.Vector{X: -6 * offset, Y: +0 * offset}))
	g.AddNode(NewNode(17, vector.Vector{X: -4 * offset, Y: +0 * offset}))
	g.AddNode(NewNode(18, vector.Vector{X: -2 * offset, Y: +0 * offset}))
	g.AddNode(NewNode(19, vector.Vector{X: +0 * offset, Y: +0 * offset}))
	g.AddNode(NewNode(20, vector.Vector{X: +2 * offset, Y: +0 * offset}))
	g.AddNode(NewNode(21, vector.Vector{X: +4 * offset, Y: +0 * offset}))
	g.AddNode(NewNode(22, vector.Vector{X: +6 * offset, Y: +0 * offset}))
	g.AddNode(NewNode(23, vector.Vector{X: -5 * offset, Y: +1 * offset}))
	g.AddNode(NewNode(24, vector.Vector{X: -3 * offset, Y: +1 * offset}))
	g.AddNode(NewNode(25, vector.Vector{X: -1 * offset, Y: +1 * offset}))
	g.AddNode(NewNode(26, vector.Vector{X: +1 * offset, Y: +1 * offset}))
	g.AddNode(NewNode(27, vector.Vector{X: +3 * offset, Y: +1 * offset}))
	g.AddNode(NewNode(28, vector.Vector{X: +5 * offset, Y: +1 * offset}))
	g.AddNode(NewNode(29, vector.Vector{X: -4 * offset, Y: +2 * offset}))
	g.AddNode(NewNode(30, vector.Vector{X: -2 * offset, Y: +2 * offset}))
	g.AddNode(NewNode(31, vector.Vector{X: +0 * offset, Y: +2 * offset}))
	g.AddNode(NewNode(32, vector.Vector{X: +2 * offset, Y: +2 * offset}))
	g.AddNode(NewNode(33, vector.Vector{X: +4 * offset, Y: +2 * offset}))
	g.AddNode(NewNode(34, vector.Vector{X: -3 * offset, Y: +3 * offset}))
	g.AddNode(NewNode(35, vector.Vector{X: -1 * offset, Y: +3 * offset}))
	g.AddNode(NewNode(36, vector.Vector{X: +1 * offset, Y: +3 * offset}))
	g.AddNode(NewNode(37, vector.Vector{X: +3 * offset, Y: +3 * offset}))

	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(1, 6)

	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(2, 6)
	g.AddEdge(2, 7)

	g.AddEdge(3, 2)
	g.AddEdge(3, 4)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)

	g.AddEdge(4, 3)
	g.AddEdge(4, 8)
	g.AddEdge(4, 9)

	g.AddEdge(5, 1)
	g.AddEdge(5, 6)
	g.AddEdge(5, 10)
	g.AddEdge(5, 11)

	g.AddEdge(6, 1)
	g.AddEdge(6, 2)
	g.AddEdge(6, 5)
	g.AddEdge(6, 7)
	g.AddEdge(6, 11)
	g.AddEdge(6, 12)

	g.AddEdge(7, 2)
	g.AddEdge(7, 3)
	g.AddEdge(7, 6)
	g.AddEdge(7, 8)
	g.AddEdge(7, 12)
	g.AddEdge(7, 13)

	g.AddEdge(8, 3)
	g.AddEdge(8, 4)
	g.AddEdge(8, 7)
	g.AddEdge(8, 9)
	g.AddEdge(8, 13)
	g.AddEdge(8, 14)

	g.AddEdge(9, 4)
	g.AddEdge(9, 8)
	g.AddEdge(9, 14)
	g.AddEdge(9, 15)

	g.AddEdge(10, 5)
	g.AddEdge(10, 11)
	g.AddEdge(10, 16)
	g.AddEdge(10, 17)

	g.AddEdge(11, 5)
	g.AddEdge(11, 6)
	g.AddEdge(11, 10)
	g.AddEdge(11, 12)
	g.AddEdge(11, 17)
	g.AddEdge(11, 18)

	g.AddEdge(12, 6)
	g.AddEdge(12, 7)
	g.AddEdge(12, 11)
	g.AddEdge(12, 13)
	g.AddEdge(12, 18)
	g.AddEdge(12, 19)

	g.AddEdge(13, 7)
	g.AddEdge(13, 8)
	g.AddEdge(13, 12)
	g.AddEdge(13, 14)
	g.AddEdge(13, 19)
	g.AddEdge(13, 20)

	g.AddEdge(14, 8)
	g.AddEdge(14, 9)
	g.AddEdge(14, 13)
	g.AddEdge(14, 15)
	g.AddEdge(14, 20)
	g.AddEdge(14, 21)

	g.AddEdge(15, 9)
	g.AddEdge(15, 14)
	g.AddEdge(15, 21)
	g.AddEdge(15, 22)

	g.AddEdge(16, 10)
	g.AddEdge(16, 17)
	g.AddEdge(16, 23)

	g.AddEdge(17, 10)
	g.AddEdge(17, 11)
	g.AddEdge(17, 16)
	g.AddEdge(17, 18)
	g.AddEdge(17, 23)
	g.AddEdge(17, 24)

	g.AddEdge(18, 11)
	g.AddEdge(18, 12)
	g.AddEdge(18, 17)
	g.AddEdge(18, 19)
	g.AddEdge(18, 24)
	g.AddEdge(18, 25)

	g.AddEdge(19, 12)
	g.AddEdge(19, 13)
	g.AddEdge(19, 18)
	g.AddEdge(19, 20)
	g.AddEdge(19, 25)
	g.AddEdge(19, 26)

	g.AddEdge(20, 13)
	g.AddEdge(20, 14)
	g.AddEdge(20, 19)
	g.AddEdge(20, 21)
	g.AddEdge(20, 26)
	g.AddEdge(20, 27)

	g.AddEdge(21, 14)
	g.AddEdge(21, 15)
	g.AddEdge(21, 20)
	g.AddEdge(21, 22)
	g.AddEdge(21, 27)
	g.AddEdge(21, 28)

	g.AddEdge(22, 15)
	g.AddEdge(22, 21)
	g.AddEdge(22, 28)

	g.AddEdge(23, 16)
	g.AddEdge(23, 17)
	g.AddEdge(23, 24)
	g.AddEdge(23, 29)

	g.AddEdge(24, 17)
	g.AddEdge(24, 18)
	g.AddEdge(24, 23)
	g.AddEdge(24, 25)
	g.AddEdge(24, 29)
	g.AddEdge(24, 30)

	g.AddEdge(25, 18)
	g.AddEdge(25, 19)
	g.AddEdge(25, 24)
	g.AddEdge(25, 26)
	g.AddEdge(25, 30)
	g.AddEdge(25, 31)

	g.AddEdge(26, 19)
	g.AddEdge(26, 20)
	g.AddEdge(26, 25)
	g.AddEdge(26, 27)
	g.AddEdge(26, 31)
	g.AddEdge(26, 32)

	g.AddEdge(27, 20)
	g.AddEdge(27, 21)
	g.AddEdge(27, 26)
	g.AddEdge(27, 28)
	g.AddEdge(27, 32)
	g.AddEdge(27, 33)

	g.AddEdge(28, 21)
	g.AddEdge(28, 22)
	g.AddEdge(28, 27)
	g.AddEdge(28, 33)

	g.AddEdge(29, 23)
	g.AddEdge(29, 24)
	g.AddEdge(29, 30)
	g.AddEdge(29, 34)

	g.AddEdge(30, 24)
	g.AddEdge(30, 25)
	g.AddEdge(30, 29)
	g.AddEdge(30, 31)
	g.AddEdge(30, 34)
	g.AddEdge(30, 35)

	g.AddEdge(31, 25)
	g.AddEdge(31, 26)
	g.AddEdge(31, 30)
	g.AddEdge(31, 32)
	g.AddEdge(31, 35)
	g.AddEdge(31, 36)

	g.AddEdge(32, 26)
	g.AddEdge(32, 27)
	g.AddEdge(32, 31)
	g.AddEdge(32, 33)
	g.AddEdge(32, 36)
	g.AddEdge(32, 37)

	g.AddEdge(33, 27)
	g.AddEdge(33, 28)
	g.AddEdge(33, 32)
	g.AddEdge(33, 37)

	g.AddEdge(34, 29)
	g.AddEdge(34, 30)
	g.AddEdge(34, 35)

	g.AddEdge(35, 30)
	g.AddEdge(35, 31)
	g.AddEdge(35, 34)
	g.AddEdge(35, 36)

	g.AddEdge(36, 31)
	g.AddEdge(36, 32)
	g.AddEdge(36, 35)
	g.AddEdge(36, 37)

	g.AddEdge(37, 32)
	g.AddEdge(37, 33)
	g.AddEdge(37, 36)

	return g
}
