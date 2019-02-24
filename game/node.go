package game

import (
	"time"

	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/math/vector"
)

// Node represents major hub of the world.
type Node struct {
	id       int64
	position vector.Vector
}

// NewNode create new node.
func NewNode(id int64, position vector.Vector) *Node {
	node := &Node{
		id,
		position,
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
func (node *Node) Tick(now time.Time) {
	// 모든 유저에게 내 정보 보내기

}
