package game

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/pkg/consts"
)

// World represents game world.
type World struct {
	actors map[int64]*Actor
	g      *graph.Graph
}

// NewWorld create new world.
func NewWorld() *World {
	world := &World{
		actors: make(map[int64]*Actor),
	}

	world.newGraph()

	return world
}

func (w *World) NewActor(actorID data.ActorID, position, velocity vector.Vector) *Actor {
	actor := NewActor(actorID, position, velocity)

	w.actors[actor.ID()] = actor

	return actor
}

func (w *World) Actor(id int64) *Actor {
	if actor, ok := w.actors[id]; ok {
		return actor
	}

	return nil
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

// Tick ticks world.
func (w *World) Tick() {
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
