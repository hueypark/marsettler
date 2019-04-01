package asset

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	CityHall *ebiten.Image
	Grass    *ebiten.Image
	Cursor   *ebiten.Image
	Worker   *ebiten.Image
)

func init() {
	CityHall = newImageFromFileBytes(city_hall)
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
