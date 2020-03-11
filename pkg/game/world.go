package game

import (
	"fmt"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/asset"
	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/pkg/consts"
)

// World represents game world.
type World struct {
	actors      map[int64]*Actor
	g           *graph.Graph
	startNodeID int64
}

// NewWorld create new world.
func NewWorld() *World {
	world := &World{
		actors: make(map[int64]*Actor),
	}

	world.newGraph()

	return world
}

// NewActor creates new actor.
func (w *World) NewActor(nodeID int64, actorID data.ActorID) *Actor {
	node := w.Node(nodeID)
	if node == nil {
		return nil
	}

	actor := node.NewActor(actorID, w)

	w.actors[actor.ID()] = actor

	return actor
}

// NewUser creates new user.
func (w *World) NewUser(nodeID int64) *User {
	actor := w.NewActor(nodeID, data.Hero)
	if actor == nil {
		return nil
	}

	user := &User{
		actor:       actor,
		world:       w,
		clickedImg:  asset.Image("/asset/tiles/clicked.png"),
		clickedNode: nil,
	}

	return user
}

func (w *World) Actor(id int64) *Actor {
	if actor, ok := w.actors[id]; ok {
		return actor
	}

	return nil
}

func (w *World) MoveActor(actor *Actor, fromNodeID, toNodeID int64) error {
	from := w.Node(fromNodeID)
	if from == nil {
		return fmt.Errorf("from node is nil")
	}
	to := w.Node(toNodeID)
	if to == nil {
		return fmt.Errorf("to node is nil")
	}

	from.DeleteActor(actor.ID())
	to.AddActor(actor)
	actor.SetNode(to)

	return nil
}

// Node returns node.
func (w *World) Node(id int64) *Node {
	n := w.g.Node(id)
	if n == nil {
		return nil
	}

	node := n.(*Node)
	return node
}

// NearestNode returns nearest node.
func (w *World) NearestNode(pos vector.Vector) *Node {
	minDistanceSQ := math.MaxFloat64
	var nearestNode graph.Node
	for _, n := range w.g.Nodes() {
		distanceSQ := pos.Sub(n.Position()).SizeSquare()
		if distanceSQ <= minDistanceSQ {
			minDistanceSQ = distanceSQ
			nearestNode = n
		}
	}

	if nearestNode == nil {
		return nil
	}
	if consts.NodeSizeHalfSq < minDistanceSQ {
		return nil
	}

	return nearestNode.(*Node)
}

func (w *World) StartNodeID() int64 {
	return w.startNodeID
}

// Tick ticks world.
func (w *World) Tick() {
	for _, actor := range w.actors {
		actor.Tick()
	}
}

func (w *World) Render(screen *ebiten.Image) {
	for _, iter := range w.g.Nodes() {
		node, ok := iter.(*Node)
		if !ok {
			log.Print("Node is not game.node.")
		}

		node.Render(screen)
	}

	for _, actor := range w.actors {
		actor.Render(screen)
	}
}

func (w *World) newGraph() {
	w.g = graph.NewGraph()
	node := NewNode(vector.Zero())
	w.startNodeID = node.ID()
	w.g.AddNode(node)

	var nodes, newNodes []*Node
	newNodes = append(newNodes, node)

	for i := 0; i < 3; i++ {
		nodes = newNodes
		newNodes = nil
		for _, node := range nodes {
			newNodes = append(newNodes, w.newNodes(node)...)
		}
	}

	for _, lhs := range w.g.Nodes() {
		for _, rhs := range w.g.Nodes() {
			if lhs.Position().Sub(rhs.Position()).SizeSquare() <= consts.NodeSizeSq {
				w.g.AddEdge(lhs.ID(), rhs.ID())
			}
		}
	}
}

func (w *World) newNodes(node *Node) []*Node {
	var newNodes []*Node

	for _, pos := range node.GetNeighborNodePositions() {
		nodeExist := w.NearestNode(pos)

		if nodeExist != nil {
			continue
		}

		newNode := NewNode(pos)
		w.g.AddNode(newNode)

		newNodes = append(newNodes, newNode)
	}

	return newNodes
}
