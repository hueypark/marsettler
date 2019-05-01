package ui

import (
	"github.com/hajimehoshi/ebiten"
)

type Menu struct {
	layer *Layer
}

func NewMenu() *Menu {
	menu := &Menu{
		layer: NewLayer(),
	}

	return menu
}

func (menu *Menu) Render(screen *ebiten.Image) {
	menu.layer.Render(screen)
}
