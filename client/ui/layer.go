package ui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics/collision_check"
)

type Layer struct {
	image       *ebiten.Image
	position    vector.Vector
	onCollision func()
	parent      *Layer
	children    []*Layer
}

func NewLayer(image *ebiten.Image, position vector.Vector, onCollision func(), parent *Layer) *Layer {
	layer := &Layer{
		image:       image,
		position:    position,
		onCollision: onCollision,
		parent:      parent,
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

	result := collision_check.PointToAABB(cursorPosition, layer)

	if result && layer.onCollision != nil {
		layer.onCollision()
	}

	return result
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
	renderer.RenderUI(screen, layer.image, layer.position.X-float64(sizeWidth/2), layer.position.Y-float64(sizeHeight/2))

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
