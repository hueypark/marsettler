package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/server/game"
)

type Actor struct {
	game.Actor
}

func (actor *Actor) Render(screen *ebiten.Image) {}
