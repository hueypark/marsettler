package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/server/game/ai"
)

// Actor represent actor.
type Actor struct {
	id           int64
	behaviorTree *behavior_tree.BehaviorTree
	image        *ebiten.Image
	position     vector.Vector
	radius       float64
}

// NewActor creates new actor.
func NewActor(actorData *data.ActorData, position vector.Vector) *Actor {
	actor := &Actor{
		id:       id_generator.Generate(),
		image:    actorData.Image,
		position: position,
		radius:   actorData.Radius,
	}

	actor.SetBehaviorTree(ai.NewAI(actor, actorData.BehaviorTree))

	return actor
}

func (actor *Actor) OnCollision(other interface{}, normal vector.Vector, penetration float64) {

}

// Render renders actor in screen.
func (actor *Actor) Render(screen *ebiten.Image) {
	renderer.Render(screen, actor.image, actor.Position())
}

// SetBehaviorTree sets behavior tree.
func (actor *Actor) SetBehaviorTree(behaviorTree *behavior_tree.BehaviorTree) {
	actor.behaviorTree = behaviorTree
}

func (actor *Actor) SetPosition(position vector.Vector) {
	actor.position = position
}

func (actor *Actor) Shape() physics.Shape {
	return physics.Circle
}

// ID returns id.
func (actor *Actor) ID() int64 {
	return actor.id
}

// Position returns position.
func (actor *Actor) Position() vector.Vector {
	return actor.position
}

func (actor *Actor) Radius() float64 {
	return actor.radius
}

// Tick ticks actor.
func (actor *Actor) Tick() {
	if actor.behaviorTree != nil {
		actor.behaviorTree.Tick()
	}
}

func (actor *Actor) CreateActor(id int) {
	//actor.node.NewActor(id)
}

func (actor *Actor) FindPath() *[]int64 {
	//return actor.node.World().RandomPath(3)
	return nil
}

// Move moves actor another node.
func (actor *Actor) Move(nodeID int64) {
	//node := GetNode(nodeID)
	//if node == nil {
	//	log.Println("node is nil", nodeID)
	//	return
	//}
	//
	//node.AddActor(actor)
}
