package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	_ "github.com/hueypark/marsettler/pkg/ai"
	"github.com/hueypark/marsettler/pkg/game"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	g := game.New()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
