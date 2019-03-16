package game

import (
	"sync"
	"time"

	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/id_generator"
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
}

// NewNode create new node.
func NewNode(id int64, position vector.Vector) *Node {
	nodeMux.Lock()
	defer nodeMux.Unlock()

	node := &Node{
		id,
		position,
	}

	nodes[id] = node

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
}

// NewActor creates new actor.
func (node *Node) NewActor() *Actor {
	actor := &Actor{
		id_generator.Generate(),
		node.id,
	}

	return actor
}

// ForEachNode executes a function for all nodes.
func ForEachNode(f func(node *Node)) {
	nodeMux.RLock()
	defer nodeMux.RUnlock()
	for _, node := range nodes {
		f(node)
	}
}
