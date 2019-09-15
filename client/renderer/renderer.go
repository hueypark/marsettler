package renderer

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/physics"
)

const (
	// enum values of ScrollState.
	ScrollStart ScrollState = iota
	ScrollEnd

	minZoom = 0.2
)

var (
	op                  ebiten.DrawImageOptions
	zoom                float64 = 1
	cameraPosition      vector.Vector
	scrollState         = ScrollEnd
	scrollStartPosition vector.Vector
	scroll              vector.Vector
)

// ScrollState represents scroll state.
type ScrollState int

// Tick represents tick.
func Tick(position vector.Vector) {
	if scrollState != ScrollStart {
		return
	}

	scroll = position.Sub(scrollStartPosition)
}

// Render renders object.
func Render(screen *ebiten.Image, img *ebiten.Image, position vector.Vector) {
	position.X *= zoom
	position.Y *= zoom

	position.X += cameraPosition.X + scroll.X
	position.Y += cameraPosition.Y + scroll.Y

	op.GeoM.Reset()
	op.GeoM.Scale(zoom, zoom)
	op.GeoM.Translate(position.X, position.Y)
	err := screen.DrawImage(img, &op)
	if err != nil {
		log.Println(err)
	}
}

func RenderAOE(screen *ebiten.Image, aoe *physics.AreaOfEffect) {
	left := (aoe.Left * zoom) + (cameraPosition.X + scroll.X)
	right := (aoe.Right * zoom) + (cameraPosition.X + scroll.X)
	bottom := (aoe.Bottom * zoom) + (cameraPosition.Y + scroll.Y)
	top := (aoe.Top * zoom) + (cameraPosition.Y + scroll.Y)

	ebitenutil.DrawLine(screen, left, top, right, top, color.White)
	ebitenutil.DrawLine(screen, right, top, right, bottom, color.White)
	ebitenutil.DrawLine(screen, right, bottom, left, bottom, color.White)
	ebitenutil.DrawLine(screen, left, bottom, left, top, color.White)
}

func RenderUI(screen *ebiten.Image, img *ebiten.Image, x, y float64) {
	op.GeoM.Reset()
	op.GeoM.Translate(x, y)
	err := screen.DrawImage(img, &op)
	if err != nil {
		log.Println(err)
	}
}

// Zoom process zoom.
func Zoom(delta float64, cursorPosition vector.Vector) {
	oldPosition := WorldPosition(cursorPosition)

	zoom += delta
	if zoom <= minZoom {
		zoom = minZoom
	}

	newPosition := WorldPosition(cursorPosition)
	deltaPosition := newPosition.Sub(oldPosition)
	cameraPosition = cameraPosition.Add(deltaPosition.Mul(zoom))
}

// OnScrollStart process scroll start event.
func OnScrollStart(position vector.Vector) {
	scrollState = ScrollStart
	scrollStartPosition = position
	scroll = vector.Zero()
}

// OnScrollEnd process scroll end event.
func OnScrollEnd() {
	scrollState = ScrollEnd
	cameraPosition = cameraPosition.Add(scroll)
	scroll = vector.Zero()
}

// WorldPosition returns world position.
func WorldPosition(cursorPosition vector.Vector) vector.Vector {
	worldPosition := vector.Vector{
		X: (cursorPosition.X - cameraPosition.X - scroll.X) / zoom,
		Y: (cursorPosition.Y - cameraPosition.Y - scroll.Y) / zoom,
	}

	return worldPosition
}
