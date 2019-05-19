package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/client"
	"github.com/hueypark/marsettler/client/config"
	"github.com/hueypark/marsettler/client/ctx"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/server/game"
)

var (
	world *game.World
	menu  *client.Menu
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	world = game.NewWorld()
	menu = client.NewMenu()

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

	world.Tick()

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

	ctx.Cursor.Render(screen, cursorPosition)

	menu.Render(screen)

	renderer.Tick(cursorPosition)
}

func tickCollision(cursorPosition, worldPosition vector.Vector) {
	if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		return
	}

	if menu.CheckCollision(cursorPosition) {
		return
	}

	ctx.Cursor.OnClick(cursorPosition)

	//world.ForEachNode(func(node *game.Node) {
	//	if collision_check.PointToAABB(worldPosition, node) {
	//		cursor.SetNode(node)
	//	}
	//})
}
