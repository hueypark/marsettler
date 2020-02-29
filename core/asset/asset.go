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
	images      map[string]*ebiten.Image
	fallbackImg *ebiten.Image

	UiFont font.Face
)

// Image returns image. If not exists it returns a fallback image.
func Image(str string) *ebiten.Image {
	if img, ok := images[str]; ok {
		return img
	}

	return fallbackImg
}

// Set image sets image.
func SetImage(str string, img *ebiten.Image) {
	images[str] = img
}

func init() {
	images = make(map[string]*ebiten.Image)

	fallbackImg = newImageFromFileBytes(fallback)

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
