package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/client/conn"
	"github.com/hueypark/marsettler/client/game"
	"github.com/hueypark/marsettler/client/renderer"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
)

var (
	ball *game.Actor
)

func main() {
	conn.SendLogin(0)

	ball = game.NewActor(cp.Vector{})

	ebiten.SetRunnableInBackground(true)
	err := ebiten.Run(tick, 800, 600, 1, "Marsettler")
	if err != nil {
		log.Fatalln(err)
	}
}

func tick(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	x, y := ebiten.CursorPosition()
	cursorPosition := vector.Vector{X: float64(x), Y: float64(y)}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		renderer.OnScrollStart(cursorPosition)
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		renderer.OnScrollEnd()

		worldPosition := renderer.WorldPosition(cursorPosition)
		ball.SetDesiredPosition(cp.Vector{X: worldPosition.X, Y: worldPosition.Y})
	}

	_, dy := ebiten.Wheel()
	renderer.OnZoom(dy * 0.05)

	renderer.Tick(cursorPosition)

	game.Space.Step(1.0 / float64(ebiten.MaxTPS()))

	err := screen.Fill(color.Black)
	if err != nil {
		return err
	}

	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(200.0/255.0, 200.0/255.0, 200.0/255.0, 1)

	game.ForEachNode(func(node *game.Node) {
		renderer.Render(screen, node.Image(), node.Position())

		node.ForEachEdge(func(edge *game.Edge) {
			if toNode := game.GetNode(edge.To); toNode != nil {
				ebitenutil.DrawLine(
					screen,
					node.Position().X,
					node.Position().Y,
					toNode.Position().X,
					toNode.Position().Y,
					colornames.White)
			}
		})
	})

	game.EachActor(func(actor *game.Actor) {
		actor.Tick()

		renderer.Render(screen, actor.Image(), vector.Vector{X: actor.Position().X, Y: actor.Position().Y})
	})

	return ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}
