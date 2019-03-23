package renderer

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/asset"
	"golang.org/x/image/colornames"
)

var (
	actorImage  *ebiten.Image
	nodeImage   *ebiten.Image
	cursorImage *ebiten.Image
)

func init() {
	actorImage, _ = ebiten.NewImage(6, 6, ebiten.FilterDefault)
	if err := actorImage.Fill(colornames.Green); err != nil {
		log.Fatalln(err)
	}

	nodeImage, _ = ebiten.NewImage(9, 9, ebiten.FilterDefault)
	if err := nodeImage.Fill(colornames.Gray); err != nil {
		log.Fatalln(err)
	}

	if img, _, err := image.Decode(bytes.NewReader(asset.Cursor)); err != nil {
		log.Fatalln(err)
	} else {
		if ebitenImage, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault); err != nil {
			log.Fatalln(err)
		} else {
			cursorImage = ebitenImage
		}
	}
}
