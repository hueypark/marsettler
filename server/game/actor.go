package game

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/server/game/ai"
)

// Actor represent actor.
type Actor struct {
	id           int64
	node         *Node
	behaviorTree *behavior_tree.BehaviorTree
	image        *ebiten.Image
}

// NewActor creates new actor.
func NewActor(node *Node, actorData *data.Actor) *Actor {
	actor := &Actor{
		id:    id_generator.Generate(),
		node:  node,
		image: actorData.Image,
	}

	actor.SetBehaviorTree(ai.NewAI(actor, actorData.BehaviorTree))

	return actor
}

// SetBehaviorTree sets behavior tree.
func (actor *Actor) SetBehaviorTree(behaviorTree *behavior_tree.BehaviorTree) {
	actor.behaviorTree = behaviorTree
}

// Render renders actor in screen.
func (actor *Actor) Render(screen *ebiten.Image) {
	renderer.Render(screen, actor.image, actor.Position())
}

// ID returns id.
func (actor *Actor) ID() int64 {
	return actor.id
}

// Position returns position.
func (actor *Actor) Position() vector.Vector {
	return actor.node.Position()
}

// Tick ticks actor.
func (actor *Actor) Tick() {
	if actor.behaviorTree != nil {
		actor.behaviorTree.Tick()
	}
}

func (actor *Actor) CreateActor(id int) {
	actor.node.NewActor(id)
}

func (actor *Actor) FindPath() *[]int64 {
	return actor.node.World().RandomPath(3)
}

// Move moves actor another node.
func (actor *Actor) Move(nodeID int64) {
	node := GetNode(nodeID)
	if node == nil {
		log.Println("node is nil", nodeID)
		return
	}

	actor.node.DeleteActor(actor.id)

	actor.node = node
	node.AddActor(actor)
}
