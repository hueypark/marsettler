package game

import (
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics"
)

// Actor represent actor.
type Actor struct {
	id           int64
	behaviorTree *behavior_tree.BehaviorTree
	position     vector.Vector
}

// NewActor creates new actor.
func NewActor(position vector.Vector) *Actor {
	actor := &Actor{
		id:       id_generator.Generate(),
		position: position,
	}

	return actor
}

func (actor *Actor) OnCollision(other interface{}, normal vector.Vector, penetration float64) {

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
	return 16.0
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
