package graph

import (
	"testing"

	"github.com/hueypark/marsettler/core/math/vector"
)

type benchmarkNode struct {
	id  int64
	pos vector.Vector
}

func (n *benchmarkNode) ID() int64 {
	return n.id
}

func (n *benchmarkNode) Position() vector.Vector {
	return n.pos
}

func newNode(id int64, pos vector.Vector) *benchmarkNode {
	return &benchmarkNode{
		id:  id,
		pos: pos,
	}
}

func BenchmarkPath(b *testing.B) {
	startNode := newNode(1, vector.Vector{X: 0, Y: 0})
	endNode := newNode(2, vector.Vector{X: 0, Y: 0})
	g := NewGraph()
	g.AddNode(startNode)
	g.AddNode(endNode)
	g.AddEdge(startNode.ID(), endNode.ID())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := g.Path(startNode.ID(), endNode.ID())
		if err != nil {
			b.Fatal(err)
		}
	}
}
