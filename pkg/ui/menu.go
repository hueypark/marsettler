package ui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/pkg/asset"
	"github.com/hueypark/marsettler/pkg/config"
	"github.com/hueypark/marsettler/pkg/ctx"
	"github.com/hueypark/marsettler/pkg/renderer"
)

type Menu struct {
	layer *Layer
}

func NewMenu() *Menu {
	_, sizeHeight := asset.Menu.Size()

	layer := NewLayer(
		"",
		asset.Menu,
		vector.Vector{X: float64(config.ScreenWidth / 2), Y: float64(config.ScreenHeight - (sizeHeight / 2))},
		nil,
		nil)

	buildingCount := len(data.Buildings())
	space := 50.0
	halfSpace := space * 0.5
	left := halfSpace - float64(buildingCount)*halfSpace
	for i, building := range data.Buildings() {
		actorData := data.Actor(building.ActorID)
		NewLayer(
			actorData.Abbreviation,
			actorData.Image,
			vector.Vector{X: left + space*float64(i), Y: 0},
			func() {
				ctx.Cursor.Set(
					func(cursorPosition vector.Vector) {
						//ctx.World.NewActor(id_generator.Generate(), renderer.WorldPosition(cursorPosition))

						ctx.Cursor.Clear()
					},
					func(screen *ebiten.Image, cursorPosition vector.Vector) {
						renderer.RenderUI(screen, actorData.Image, cursorPosition.X, cursorPosition.Y)
					},
				)
			},
			layer)
	}

	menu := &Menu{layer}

	return menu
}

func (menu *Menu) CheckCollision(position vector.Vector) bool {
	return menu.layer.CheckCollision(position)
}

func (menu *Menu) Render(screen *ebiten.Image) {
	menu.layer.Render(screen)
}
