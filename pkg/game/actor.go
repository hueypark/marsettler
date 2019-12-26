package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics/body"
	"github.com/hueypark/marsettler/core/physics/body/circle"
	"github.com/hueypark/marsettler/pkg/asset"
	"github.com/hueypark/marsettler/pkg/renderer"
)

// Actor represent actor.
type Actor struct {
	behaviorTree *behavior_tree.BehaviorTree
	body         *body.Body
}

const radius float64 = 16.0

// NewActor creates new actor.
func NewActor(id int64, position, velocity vector.Vector) *Actor {
	actor := &Actor{}

	actor.Init(id, position, velocity)

	return actor
}

func (actor *Actor) Init(id int64, position, velocity vector.Vector) {
	b := body.New(id, position)
	b.Velocity = velocity
	b.SetMass(10)
	b.SetShape(circle.New(radius))
	actor.body = b
}

func (actor *Actor) OnCollision(other interface{}, normal vector.Vector, penetration float64) {
}

// SetBehaviorTree sets behavior tree.
func (actor *Actor) SetBehaviorTree(behaviorTree *behavior_tree.BehaviorTree) {
	actor.behaviorTree = behaviorTree
}

func (actor *Actor) SetPosition(position vector.Vector) {
	actor.body.SetPosition(position)
}

func (actor *Actor) SetVelocity(velocity vector.Vector) {
	actor.body.SetVelocity(velocity)
}

func (actor *Actor) Shape() body.Shape {
	return body.Circle
}

// ID returns id.
func (actor *Actor) ID() int64 {
	return actor.body.ID()
}

// Position returns position.
func (actor *Actor) Position() vector.Vector {
	return actor.body.Position()
}

func (actor *Actor) Render(screen *ebiten.Image) {
	pos := actor.Position()
	radiusHalf := actor.Radius() * 0.5
	pos.X -= radiusHalf
	pos.Y -= radiusHalf

	renderer.Render(screen, asset.Circle, pos)
}

func (actor *Actor) Velocity() vector.Vector {
	return actor.body.Velocity
}

func (actor *Actor) Radius() float64 {
	return radius
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

func (actor *Actor) Body() *body.Body {
	return actor.body
}
