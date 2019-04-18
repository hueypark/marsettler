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
	"github.com/hueypark/marsettler/core/physics/collision_check"
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
	centerNode.NewActor(1)
	centerNode.NewActor(100000)

	ebiten.SetRunnableInBackground(true)
	err := ebiten.Run(tick, 800, 600, 1, "Marsettler")
	if err != nil {
		log.Fatalln(err)
	}
}

func tick(screen *ebiten.Image) error {
	x, y := ebiten.CursorPosition()
	cursorPosition := vector.Vector{X: float64(x), Y: float64(y)}

	worldPosition := renderer.WorldPosition(cursorPosition)

	tickRenderer(cursorPosition)
	tickCollision(worldPosition)

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

func tickRenderer(cursorPosition vector.Vector) {
	_, dy := ebiten.Wheel()
	renderer.Zoom(dy*0.1, cursorPosition)

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		renderer.OnScrollStart(cursorPosition)
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		renderer.OnScrollEnd()
	}

	renderer.Tick(cursorPosition)
}

func tickCollision(worldPosition vector.Vector) {
	if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		return
	}

	world.ForEachNode(func(node *game.Node) {
		if collision_check.PointToAABB(worldPosition, node) {
			cursor.SetNode(node)
		}
	})
}
