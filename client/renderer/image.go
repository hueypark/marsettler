package renderer

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"golang.org/x/image/colornames"
)

var (
	actorImage *ebiten.Image
	nodeImage  *ebiten.Image
)

func init() {
	actorImage, _ = ebiten.NewImage(30, 30, ebiten.FilterDefault)
	if err := actorImage.Fill(colornames.White); err != nil {
		log.Fatal(err)
	}

	nodeImage, _ = ebiten.NewImage(9, 9, ebiten.FilterDefault)
	if err := nodeImage.Fill(colornames.Gray); err != nil {
		log.Fatal(err)
	}
}
