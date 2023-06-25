package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hueypark/marsettler/pkg/marsettler"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Marsettler")
	game := marsettler.NewGame()
	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
