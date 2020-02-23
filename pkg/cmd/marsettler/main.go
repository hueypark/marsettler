package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/config"
	"github.com/hueypark/marsettler/pkg/ctx"
	"github.com/hueypark/marsettler/pkg/game"
	"github.com/hueypark/marsettler/pkg/renderer"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ctx.World = game.NewWorld()
	ctx.User = game.NewUser(ctx.World)

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

	sec := (time.Second / 60).Seconds()
	ctx.User.Tick(sec, worldPosition)
	ctx.World.Tick(sec)
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

	ctx.World.Render(screen)
}
