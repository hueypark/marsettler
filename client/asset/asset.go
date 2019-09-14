package asset

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	Blueberry *ebiten.Image
	Building  *ebiten.Image
	Circle    *ebiten.Image
	CityHall  *ebiten.Image
	Cursor    *ebiten.Image
	Fairy     *ebiten.Image
	Grass     *ebiten.Image
	Menu      *ebiten.Image
	Worker    *ebiten.Image

	UiFont font.Face
)

func init() {
	Blueberry = newImageFromFileBytes(blueberry)
	Building = newImageFromFileBytes(building)
	Circle = newImageFromFileBytes(circle)
	CityHall = newImageFromFileBytes(city_hall)
	Cursor = newImageFromFileBytes(cursor)
	Fairy = newImageFromFileBytes(fairy)
	Grass = newImageFromFileBytes(grass)
	Menu = newImageFromFileBytes(menu)
	Worker = newImageFromFileBytes(worker)

	initFont()
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

func initFont() {
	tt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	UiFont = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}
