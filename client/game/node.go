package game

import (
	"sync"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/math/vector"
)

var (
	nodeMux sync.RWMutex
	nodes   = map[int64]*Node{}
)

// Node represents major hub of the world.
type Node struct {
	id       int64
	position vector.Vector
	edges    []*Edge
}

// NewNode create new node.
func NewNode(id int64, position vector.Vector, edges []*Edge) *Node {
	nodeMux.Lock()
	defer nodeMux.Unlock()

	node := &Node{
		id,
		position,
		edges,
	}

	nodes[id] = node

	return node
}

// GetNode returns node.
func GetNode(id int64) *Node {
	nodeMux.RLock()
	defer nodeMux.RUnlock()

	if node, ok := nodes[id]; ok {
		return node
	} else {
		return nil
	}
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

// Image returns image.
func (Node) Image() *ebiten.Image {
	return nodeImage
}

// ForEachEdge executes a function for node edges.
func (node *Node) ForEachEdge(f func(edge *Edge)) {
	for _, edge := range node.edges {
		f(edge)
	}
}

// ForEachNode executes a function for all nodes.
func ForEachNode(f func(node *Node)) {
	nodeMux.RLock()
	defer nodeMux.RUnlock()
	for _, node := range nodes {
		f(node)
	}
}
