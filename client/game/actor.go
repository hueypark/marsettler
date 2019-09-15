package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/asset"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/message"
	"github.com/hueypark/marsettler/server/game"
)

type Actor struct {
	game.Actor
}

func NewActor(msgActor *message.Actor) *Actor {
	actor := &Actor{}
	actor.Init(vector.Vector{X: msgActor.Pos.X, Y: msgActor.Pos.Y})

	return actor
}

func (actor *Actor) Render(screen *ebiten.Image) {
	pos := actor.Position()
	radiusHalf := actor.Radius() * 0.5
	pos.X -= radiusHalf
	pos.Y -= radiusHalf

	renderer.Render(screen, asset.Circle, pos)
}
