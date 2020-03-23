package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	_ "github.com/hueypark/asset"
	"github.com/hueypark/marsettler/core/math/vector"
	_ "github.com/hueypark/marsettler/pkg/ai"
	"github.com/hueypark/marsettler/pkg/config"
	"github.com/hueypark/marsettler/pkg/consts"
	"github.com/hueypark/marsettler/pkg/game"
	"github.com/hueypark/marsettler/pkg/renderer"
)

var (
	world *game.World
	user  *game.User
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	world = game.NewWorld()
	kingdom := game.NewKingdom()
	startNodeID, err := world.StartNodeID()
	if err != nil {
		panic(err)
	}
	user = world.NewUser(kingdom.ID(), startNodeID)

	ebiten.SetRunnableInBackground(true)
	ebiten.SetMaxTPS(consts.TPS)
	err = ebiten.Run(tick, config.ScreenWidth, config.ScreenHeight, 1, "Marsettler")
	if err != nil {
		log.Fatalln(err)
	}

}

func tick(screen *ebiten.Image) error {
	ebiten.CurrentFPS()
	x, y := ebiten.CursorPosition()
	cursorPosition := vector.Vector{X: float64(x), Y: float64(y)}

	worldPosition := renderer.WorldPosition(cursorPosition)

	user.Tick(worldPosition)
	world.Tick()
	tickRenderer(screen, cursorPosition)

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

	world.Render(screen)
	user.Render(screen)
}
