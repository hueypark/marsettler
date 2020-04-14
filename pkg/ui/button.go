package ui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/asset"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/renderer"
)

type Button struct {
	image    *ebiten.Image
	position vector.Vector
	onClick  func()
	children []*Button
}

func NewButton(image string, position vector.Vector, onCollision func()) *Button {
	_ = asset.Image(image)
	layer := &Button{
		image:    asset.Image(image),
		position: position,
		onClick:  onCollision,
	}

	return layer
}

func (b *Button) AddChild(child *Button) {
	b.children = append(b.children, child)
}

// TODO(hueypark): Implement below.
func (b *Button) OnClick(cursorPosition vector.Vector) {
	//for _, child := range b.children {
	//	if child.OnClick(cursorPosition) {
	//		return
	//	}
	//}

	//return false
	//result := collision_check.PointToAABB(cursorPosition, b)
	//
	//if result && b.onClick != nil {
	//	b.onClick()
	//}
	//
	//return result
}

func (b *Button) Render(screen *ebiten.Image, relativePos *vector.Vector) {
	imgWidth, imgHeight := b.image.Size()
	posX, posY := b.position.X-float64(imgWidth/2), b.position.Y-float64(imgHeight/2)
	if relativePos != nil {
		posX += relativePos.X
		posY += relativePos.Y
	}
	renderer.RenderUI(screen, b.image, posX, posY)

	for _, child := range b.children {
		child.Render(screen, relativePos)
	}
}

func (b *Button) Position() vector.Vector {
	return b.position
}

func (b *Button) SetImage(image *ebiten.Image) {
	b.image = image
}
