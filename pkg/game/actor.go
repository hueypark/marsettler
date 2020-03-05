package game

import (
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/asset"
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics/body"
	"github.com/hueypark/marsettler/core/physics/body/circle"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/pkg/ai"
	"github.com/hueypark/marsettler/pkg/renderer"
)

// Actor represent actor.
type Actor struct {
	behaviorTree *behavior_tree.BehaviorTree
	body         *body.Body
	image        *ebiten.Image
}

const radius float64 = 16.0

// NewActor creates new actor.
func NewActor(actorID data.ActorID, position, velocity vector.Vector) *Actor {
	actor := &Actor{}

	actor.Init(id_generator.Generate(), position, velocity)
	if actorData := data.Actor(actorID); actorData != nil {
		if bt, err := ai.NewBehaviorTree(actor, actorData.BehaviorTree); err == nil {
			actor.SetBehaviorTree(bt)
		} else {
			log.Println(err)
		}

		actor.image = asset.Image(actorData.Image)
	}

	return actor
}

func (actor *Actor) Init(id int64, position, velocity vector.Vector) {
	b := body.New(id, position)
	b.Velocity = velocity
	b.SetMass(10)
	b.SetShape(circle.New(radius))
	actor.body = b
}

func (actor *Actor) FindRandomPosition() vector.Vector {
	return vector.Vector{X: rand.Float64() * 500, Y: rand.Float64() * 500}
}

func (actor *Actor) MoveTo(dest vector.Vector) (arrvie bool) {
	maxSpeed := 500.0
	acc := 500.0
	dec := 700.0
	lateralDec := 100.0
	arriveDistance := 10.0
	arriveSpeed := 10.0

	diff := dest.Sub(actor.Position())

	right := diff.Right().Normalize()
	rightSpeed := actor.body.Velocity.Dot(right)
	lateralDec = math.Min(math.Abs(rightSpeed), lateralDec)
	if 0 <= rightSpeed {
		actor.body.AddForce(right.Invert().Mul(lateralDec * actor.body.Mass()))
	} else {
		actor.body.AddForce(right.Mul(lateralDec * actor.body.Mass()))
	}

	remainingDistance := diff.Size()
	dir := diff.Normalize()
	curSpeed := math.Abs(actor.body.Velocity.Dot(dir))
	breakingDistance := curSpeed * curSpeed / dec
	if remainingDistance <= breakingDistance {
		actor.body.AddForce(actor.body.Velocity.Normalize().Invert().Mul(dec * actor.body.Mass()))
	} else if curSpeed <= maxSpeed {
		actor.body.AddForce(dir.Mul(acc * actor.body.Mass()))
	}

	if arriveDistance < remainingDistance {
		return false
	}

	if arriveSpeed < curSpeed {
		actor.body.AddForce(actor.body.Velocity.Normalize().Invert().Mul(dec * actor.body.Mass()))
		return false
	}

	actor.body.Clear()
	return true
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
	x, _ := actor.image.Size()
	radiusHalf := float64(x) * 0.5
	pos.X -= radiusHalf
	pos.Y -= radiusHalf

	renderer.Render(screen, actor.image, pos)
}

func (actor *Actor) Velocity() vector.Vector {
	return actor.body.Velocity
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

func (actor *Actor) Body() *body.Body {
	return actor.body
}
