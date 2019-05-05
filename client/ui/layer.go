package ui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/core/math/vector"
)

type Layer struct {
	image    *ebiten.Image
	position vector.Vector
}

func NewLayer(image *ebiten.Image, position vector.Vector) *Layer {
	layer := &Layer{
		image:    image,
		position: position,
	}

	return layer
}

func (layer *Layer) Render(screen *ebiten.Image) {
	sizeWidth, sizeHeight := layer.image.Size()
	renderer.RenderUI(screen, layer.image, layer.position.X-float64(sizeWidth/2), layer.position.Y-float64(sizeHeight/2))
}

func (layer *Layer) Position() vector.Vector {
	return layer.position
}

func (layer *Layer) SetImage(image *ebiten.Image) {
	layer.image = image
}
