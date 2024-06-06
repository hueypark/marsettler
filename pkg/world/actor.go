package world

import "github.com/hajimehoshi/ebiten/v2"

type Actor interface {
	Draw(screen *ebiten.Image)
}
