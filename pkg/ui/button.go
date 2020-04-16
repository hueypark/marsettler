package ui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/asset"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics/test"
	"github.com/hueypark/marsettler/pkg/renderer"
)

type Button struct {
	image    *ebiten.Image
	position vector.Vector
	left     float64
	right    float64
	bottom   float64
	top      float64
	onClick  func()
	children []*Button
}

func NewButton(imgStr string, position vector.Vector, onCollision func()) *Button {
	img := asset.Image(imgStr)
	imgWidth, imgHeight := img.Size()
	imgWidthHalf := float64(imgWidth) * 0.5
	imgHeightHalf := float64(imgHeight) * 0.5

	layer := &Button{
		image:    img,
		position: position,
		left:     position.X - imgWidthHalf,
		right:    position.X + imgWidthHalf,
		bottom:   position.Y - imgHeightHalf,
		top:      position.Y + imgHeightHalf,
		onClick:  onCollision,
	}

	return layer
}

func (b *Button) AddChild(child *Button) {
	b.children = append(b.children, child)
}

func (b *Button) OnClick(cursorPosition vector.Vector) bool {
	for _, child := range b.children {
		if child.OnClick(cursorPosition) {
			return true
		}
	}

	result := test.PointToAABB(cursorPosition, b)

	if result && b.onClick != nil {
		b.onClick()
	}

	return result
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

func (b *Button) Left() float64 {
	return b.left
}

func (b *Button) Right() float64 {
	return b.right
}

func (b *Button) Bottom() float64 {
	return b.bottom
}

func (b *Button) Top() float64 {
	return b.top
}

func (b *Button) SetImage(image *ebiten.Image) {
	b.image = image
}
