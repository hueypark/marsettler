package client

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/asset"
	"github.com/hueypark/marsettler/client/config"
	"github.com/hueypark/marsettler/client/ui"
	"github.com/hueypark/marsettler/core/math/vector"
)

type Menu struct {
	layer *ui.Layer
}

func NewMenu() *Menu {
	_, sizeHeight := asset.Menu.Size()
	menu := &Menu{
		layer: ui.NewLayer(asset.Menu, vector.Vector{X: float64(config.ScreenWidth / 2), Y: float64(config.ScreenHeight - (sizeHeight / 2))}),
	}

	return menu
}

func (menu *Menu) Render(screen *ebiten.Image) {
	menu.layer.Render(screen)
}
