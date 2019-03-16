package renderer

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/math/vector"
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
	width, height := img.Size()
	op.GeoM.Scale(zoom, zoom)
	op.GeoM.Translate(position.X-(float64(width)*0.5), position.Y-(float64(height)*0.5))
	err := screen.DrawImage(img, &op)
	if err != nil {
		log.Println(err)
	}
}

// RenderNode renders node.
func RenderNode(screen *ebiten.Image, position vector.Vector) {
	Render(screen, nodeImage, position)
}

// Zoom returns zoom.
func Zoom() float64 {
	return zoom
}

// OnZoom process zoom.
func OnZoom(delta float64) {
	zoom += delta
	if zoom <= minZoom {
		zoom = minZoom
	}
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

	log.Println(cameraPosition.X)
	log.Println(zoom)
	log.Println(worldPosition.X)
	return worldPosition
}
