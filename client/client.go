package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hueypark/marsettler/client/renderer"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hueypark/marsettler/server/game"
)

var world *game.World

func main() {
	world = game.NewWorld()

	ebiten.SetRunnableInBackground(true)
	err := ebiten.Run(tick, 800, 600, 1, "Marsettler")
	if err != nil {
		log.Fatalln(err)
	}
}

func tick(screen *ebiten.Image) error {
	now := time.Now()
	world.Tick(now)

	world.ForEachNode(func(n *game.Node) {
		renderer.RenderNode(screen, n.Position())
	})

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}
