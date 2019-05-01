package ui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/asset"
	"github.com/hueypark/marsettler/client/config"
	"github.com/hueypark/marsettler/client/renderer"
)

type Layer struct {
}

func NewLayer() *Layer {
	layer := &Layer{}

	return layer
}

func (layer *Layer) Render(screen *ebiten.Image) {
	sizeWidth, sizeHeight := asset.Menu.Size()
	renderer.RenderUI(screen, asset.Menu, float64((config.ScreenWidth/2)-(sizeWidth/2)), float64(config.ScreenHeight-sizeHeight))
}
