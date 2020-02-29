package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/data"
)

type User struct {
	actor *Actor
}

func NewUser(world *World) *User {
	user := &User{
		actor: world.NewActor(data.Leader, vector.Zero(), vector.Zero()),
	}

	return user
}

func (u *User) Tick(position vector.Vector) {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		u.actor.behaviorTree.Blackboard().Set(-1, &position)
	}
}
