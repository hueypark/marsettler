package world

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hueypark/marsettler/pkg/resource"
)

func drawText(screen *ebiten.Image, s string, cx, cy int, clr color.Color) {
	bounds := text.BoundString(resource.Font, s)
	x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2
	text.Draw(screen, s, resource.Font, x, y, clr)
}
