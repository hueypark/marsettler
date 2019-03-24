package asset

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	Grass  *ebiten.Image
	Cursor *ebiten.Image
	Worker *ebiten.Image
)

func init() {
	Grass = newImageFromFileBytes(grass)
	Cursor = newImageFromFileBytes(cursor)
	Worker = newImageFromFileBytes(worker)
}

func newImageFromFileBytes(src []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(src))
	if err != nil {
		log.Fatalln(err)
	}

	ebitenImage, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatalln(err)
	}

	return ebitenImage
}
