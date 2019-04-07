package asset

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	Blueberry *ebiten.Image
	CityHall  *ebiten.Image
	Fairy     *ebiten.Image
	Grass     *ebiten.Image
	Cursor    *ebiten.Image
	Worker    *ebiten.Image
)

func init() {
	Blueberry = newImageFromFileBytes(blueberry)
	CityHall = newImageFromFileBytes(city_hall)
	Fairy = newImageFromFileBytes(fairy)
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
