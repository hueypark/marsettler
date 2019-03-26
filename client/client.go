package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/client/ui"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/server/game"
	"github.com/hueypark/marsettler/server/game/ai"
)

var (
	world  *game.World
	cursor *ui.Cursor
)

func main() {
	world = game.NewWorld()
	centerNode := world.GetCenterNode()
	cursor = ui.NewCursor(centerNode)
	actor := centerNode.NewActor()
	actor.SetBehaviorTree(ai.NewWorker(world.RandomPath, actor.Move))

	ebiten.SetRunnableInBackground(true)
	err := ebiten.Run(tick, 800, 600, 1, "Marsettler")
	if err != nil {
		log.Fatalln(err)
	}
}

func tick(screen *ebiten.Image) error {
	tickRenderer()

	world.Tick()

	world.ForEachNode(func(node *game.Node) {
		node.Render(screen)
	})

	world.ForEachNode(func(node *game.Node) {
		node.ForEachActor(func(actor *game.Actor) {
			actor.Render(screen)
		})
	})

	cursor.Render(screen)

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func tickRenderer() {
	_, dy := ebiten.Wheel()
	renderer.Zoom(dy * 0.1)

	x, y := ebiten.CursorPosition()
	cursorPosition := vector.Vector{X: float64(x), Y: float64(y)}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		renderer.OnScrollStart(cursorPosition)
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		renderer.OnScrollEnd()
	}

	renderer.Tick(cursorPosition)
}
