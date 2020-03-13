package game

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/asset"
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics/body"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/pkg/consts"
	"github.com/hueypark/marsettler/pkg/renderer"
)

// Actor represent actor.
type Actor struct {
	id           int64
	kingdomID    int64
	data         *data.ActorData
	behaviorTree *behavior_tree.BehaviorTree
	image        *ebiten.Image
	world        *World
	node         *Node

	// moveWaitTime represents the wait time for the move(millie seconds).
	// If value is less than 0 actor can move.
	moveWaitTime int
}

// NewBehaviorTree creates new bahavior tree.
var NewBehaviorTree func(actor *Actor, str string) (*behavior_tree.BehaviorTree, error)

// CanMove returns whether or not it can be moved.
func (a *Actor) CanMove() bool {
	return a.moveWaitTime <= 0
}

// FindPath finds path between nodes.
func (a *Actor) FindPath(toNodeID int64) (path graph.Path, err error) {
	return a.world.g.Path(a.node.ID(), toNodeID)
}

func (a *Actor) FindRandomPosition() vector.Vector {
	return vector.Vector{X: rand.Float64() * 500, Y: rand.Float64() * 500}
}

// Move moves actor to node.
func (a *Actor) Move(nodeID int64) error {
	if err := a.world.MoveActor(a, a.node.ID(), nodeID); err != nil {
		return err
	}

	a.moveWaitTime += a.data.MoveWaitTime

	return nil
}

// NodeID returns node id.
func (a *Actor) NodeID() int64 {
	return a.node.ID()
}

// SetBehaviorTree sets behavior tree.
func (a *Actor) SetBehaviorTree(behaviorTree *behavior_tree.BehaviorTree) {
	a.behaviorTree = behaviorTree
}

// SetNode sets node.
func (a *Actor) SetNode(node *Node) {
	a.node = node
}

func (a *Actor) Shape() body.Shape {
	return body.Circle
}

// ID returns id.
func (a *Actor) ID() int64 {
	return a.id
}

// Position returns position.
func (a *Actor) Position() vector.Vector {
	return a.node.Position()
}

func (a *Actor) Render(screen *ebiten.Image) {
	pos := a.Position()
	x, _ := a.image.Size()
	radiusHalf := float64(x) * 0.5
	pos.X -= radiusHalf
	pos.Y -= radiusHalf

	renderer.Render(screen, a.image, pos)
}

// Tick ticks actor.
func (a *Actor) Tick() {
	if a.behaviorTree != nil {
		a.behaviorTree.Tick()
	}

	if 0 < a.moveWaitTime {
		a.moveWaitTime -= consts.Delta
	}
}

// newActor creates new actor.
func newActor(kingdomID int64, actorID data.ActorID, world *World, node *Node) *Actor {
	actor := &Actor{
		id:        id_generator.Generate(),
		kingdomID: kingdomID,
		world:     world,
		node:      node,
	}

	if actorData := data.Actor(actorID); actorData != nil {
		actor.data = actorData

		if NewBehaviorTree != nil {
			if bt, err := NewBehaviorTree(actor, actorData.BehaviorTree); err == nil {
				actor.SetBehaviorTree(bt)
			} else {
				log.Println(err)
			}
		}

		actor.image = asset.Image(actorData.Image)
	}

	return actor
}
