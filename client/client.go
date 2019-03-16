package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/client/ui"
	"github.com/hueypark/marsettler/server/game"
)

var (
	world  *game.World
	cursor *ui.Cursor
)

func main() {
	world = game.NewWorld()
	cursor = ui.NewCursor(world.GetCenterNodeId())

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

	world.ForEachActor(func(a *game.Actor) {
		renderer.RenderActor(screen, a)
	})

	renderer.RenderCursor(screen, cursor.Position())

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}
