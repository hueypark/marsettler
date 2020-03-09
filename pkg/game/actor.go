package game

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/asset"
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics/body"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/pkg/renderer"
)

// Actor represent actor.
type Actor struct {
	id           int64
	behaviorTree *behavior_tree.BehaviorTree
	image        *ebiten.Image
}

// NewBehaviorTree creates new bahavior tree.
var NewBehaviorTree func(actor *Actor, str string) (*behavior_tree.BehaviorTree, error)

const radius float64 = 16.0

// NewActor creates new actor.
func NewActor(actorID data.ActorID) *Actor {
	actor := &Actor{
		id: id_generator.Generate(),
	}

	if actorData := data.Actor(actorID); actorData != nil {
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

func (actor *Actor) FindRandomPosition() vector.Vector {
	return vector.Vector{X: rand.Float64() * 500, Y: rand.Float64() * 500}
}

func (actor *Actor) MoveTo(dest vector.Vector) (arrvie bool) {
	return true
}

func (actor *Actor) OnCollision(other interface{}, normal vector.Vector, penetration float64) {
}

// SetBehaviorTree sets behavior tree.
func (actor *Actor) SetBehaviorTree(behaviorTree *behavior_tree.BehaviorTree) {
	actor.behaviorTree = behaviorTree
}

func (actor *Actor) Shape() body.Shape {
	return body.Circle
}

// ID returns id.
func (actor *Actor) ID() int64 {
	return actor.id
}

// Position returns position.
func (actor *Actor) Position() vector.Vector {
	return vector.Zero()
}

func (actor *Actor) Render(screen *ebiten.Image) {
	pos := actor.Position()
	x, _ := actor.image.Size()
	radiusHalf := float64(x) * 0.5
	pos.X -= radiusHalf
	pos.Y -= radiusHalf

	renderer.Render(screen, actor.image, pos)
}

func (actor *Actor) Radius() float64 {
	return radius
}

// Tick ticks actor.
func (actor *Actor) Tick(delta float64) {
	if actor.behaviorTree != nil {
		actor.behaviorTree.Tick(delta)
	}
}

func (actor *Actor) CreateActor(id int) {
	//actor.node.NewActor(id)
}
