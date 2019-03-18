package main

import (
	"fmt"
	"log"

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
	centerNode := world.GetCenterNode()
	cursor = ui.NewCursor(centerNode)
	centerNode.NewActor()

	ebiten.SetRunnableInBackground(true)
	err := ebiten.Run(tick, 800, 600, 1, "Marsettler")
	if err != nil {
		log.Fatalln(err)
	}
}

func tick(screen *ebiten.Image) error {
	world.Tick()

	world.ForEachNode(func(node *game.Node) {
		renderer.RenderNode(screen, node.Position())
	})

	world.ForEachNode(func(node *game.Node) {
		node.ForEachActor(func(actor *game.Actor) {
			renderer.RenderActor(screen, actor)
		})
	})

	renderer.RenderCursor(screen, cursor.Position())

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}
