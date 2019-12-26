package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/asset"
	"github.com/hueypark/marsettler/pkg/renderer"
)

type Layer struct {
	name        string
	image       *ebiten.Image
	position    vector.Vector
	onCollision func()
	parent      *Layer
	children    []*Layer
}

func NewLayer(name string, image *ebiten.Image, position vector.Vector, onCollision func(), parent *Layer) *Layer {
	layer := &Layer{
		name:        name,
		image:       image,
		position:    position,
		onCollision: onCollision,
		parent:      parent,
	}

	if parent != nil {
		parent.AddChild(layer)
	}

	return layer
}

func (layer *Layer) AddChild(child *Layer) {
	layer.children = append(layer.children, child)
}

func (layer *Layer) CheckCollision(cursorPosition vector.Vector) bool {
	for _, child := range layer.children {
		if child.CheckCollision(cursorPosition) {
			return true
		}
	}

	return false
	//result := collision_check.PointToAABB(cursorPosition, layer)
	//
	//if result && layer.onCollision != nil {
	//	layer.onCollision()
	//}
	//
	//return result
}

func (layer *Layer) Left() float64 {
	sizeWidth, _ := layer.image.Size()

	return layer.ScreenPosition().X - float64(sizeWidth/2)
}

func (layer *Layer) Right() float64 {
	sizeWidth, _ := layer.image.Size()

	return layer.ScreenPosition().X + float64(sizeWidth/2)
}

func (layer *Layer) Bottom() float64 {
	_, sizeHeight := layer.image.Size()

	return layer.ScreenPosition().Y - float64(sizeHeight/2)
}

func (layer *Layer) Top() float64 {
	_, sizeHeight := layer.image.Size()

	return layer.ScreenPosition().Y + float64(sizeHeight/2)
}

func (layer *Layer) Render(screen *ebiten.Image) {
	sizeWidth, sizeHeight := layer.image.Size()
	screenPosition := layer.ScreenPosition()
	posX, posY := screenPosition.X-float64(sizeWidth/2), screenPosition.Y-float64(sizeHeight/2)
	renderer.RenderUI(screen, layer.image, posX, posY)

	text.Draw(screen, layer.name, asset.UiFont, int(posX), int(posY), color.White)

	for _, child := range layer.children {
		child.Render(screen)
	}
}

func (layer *Layer) Position() vector.Vector {
	return layer.position
}

func (layer *Layer) SetImage(image *ebiten.Image) {
	layer.image = image
}

func (layer *Layer) ScreenPosition() vector.Vector {
	if layer.parent == nil {
		return layer.position
	}

	return layer.parent.ScreenPosition().Add(layer.position)
}
