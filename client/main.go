package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/client/config"
	"github.com/hueypark/marsettler/client/ctx"
	"github.com/hueypark/marsettler/client/game"
	"github.com/hueypark/marsettler/client/handler"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/net"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ctx.Client = net.NewClient("127.0.0.1:8080", handler.Handle)
	ctx.World = game.NewWorld()

	ebiten.SetRunnableInBackground(true)
	err := ebiten.Run(tick, config.ScreenWidth, config.ScreenHeight, 1, "Marsettler")
	if err != nil {
		log.Fatalln(err)
	}

}

func tick(screen *ebiten.Image) error {
	x, y := ebiten.CursorPosition()
	cursorPosition := vector.Vector{X: float64(x), Y: float64(y)}

	worldPosition := renderer.WorldPosition(cursorPosition)

	tickRenderer(screen, cursorPosition)
	tickCollision(cursorPosition, worldPosition)

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func tickRenderer(screen *ebiten.Image, cursorPosition vector.Vector) {
	_, dy := ebiten.Wheel()
	renderer.Zoom(dy*0.1, cursorPosition)

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		renderer.OnScrollStart(cursorPosition)
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		renderer.OnScrollEnd()
	}

	renderer.Tick(cursorPosition)

	renderWorld(screen)

	ctx.Cursor.Render(screen, cursorPosition)
}

func renderWorld(screen *ebiten.Image) {
	//for _, actor := range ctx.World.Actors() {
	//	actor.Render(screen)
	//}
}

func tickCollision(cursorPosition, worldPosition vector.Vector) {
	if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		return
	}

	ctx.Cursor.OnClick(cursorPosition)

	//world.ForEachNode(func(node *game.Node) {
	//	if collision_check.PointToAABB(worldPosition, node) {
	//		cursor.SetNode(node)
	//	}
	//})
}
