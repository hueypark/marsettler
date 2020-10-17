package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hueypark/marsettler/pkg/asset"
	"github.com/hueypark/marsettler/pkg/data"
	"github.com/hueypark/marsettler/pkg/internal/game"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
)

// Actor is basic object in world.
type Actor struct {
	*game.Actor
	clientPosition *math2d.Vector
	image          *ebiten.Image

	hp    int32
	maxHP int32
}

// NewActor Creates new actor.
func NewActor(id int64, dataID data.ActorID, position *math2d.Vector) (*Actor, error) {
	internalActor, err := game.NewActor(id, dataID, position, nil)
	if err != nil {
		return nil, err
	}

	a := &Actor{
		Actor:          internalActor,
		clientPosition: &math2d.Vector{},
	}
	a.clientPosition.Set(position)
	a.SetPosition(position)

	a.image = asset.Image("circle")

	return a, nil
}

// Draw implements ebiten.Game.Draw.
func (a *Actor) Draw(screen *ebiten.Image, cameraFunc func(*Actor) ebiten.GeoM) error {
	text.Draw(screen, fmt.Sprintf("%v/%v", a.hp, a.maxHP), asset.UiFont, int(a.Position().X), int(-a.Position().Y), color.White)

	return screen.DrawImage(
		a.image,
		&ebiten.DrawImageOptions{
			GeoM: cameraFunc(a),
		})
}

// HP returns hp.
func (a *Actor) HP() int32 {
	return a.hp
}

// MaxHP returns max hp.
func (a *Actor) MaxHP() int32 {
	return a.maxHP
}

// Position is actor's position.
func (a *Actor) Position() *math2d.Vector {
	return a.clientPosition
}

// SetHP sets hp.
func (a *Actor) SetHP(hp int32) {
	a.hp = hp
}

// SetMaxHP sets max hp.
func (a *Actor) SetMaxHP(maxHP int32) {
	a.maxHP = maxHP
}

// Tick updates actor periodically.
func (a *Actor) Tick(delta float64) error {
	serverPositionn := a.Actor.Position()

	a.clientPosition.X = a.clientPosition.X + (0.1 * (serverPositionn.X - a.clientPosition.X))
	a.clientPosition.Y = a.clientPosition.Y + (0.1 * (serverPositionn.Y - a.clientPosition.Y))

	return nil
}
