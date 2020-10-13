package game

import (
	"github.com/hajimehoshi/ebiten"
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
	return screen.DrawImage(
		a.image,
		&ebiten.DrawImageOptions{
			GeoM: cameraFunc(a),
		})
}

// Position is actor's position.
func (a *Actor) Position() *math2d.Vector {
	return a.clientPosition
}

// Tick updates actor periodically.
func (a *Actor) Tick(delta float64) error {
	serverPositionn := a.Actor.Position()

	a.clientPosition.X = a.clientPosition.X + (0.1 * (serverPositionn.X - a.clientPosition.X))
	a.clientPosition.Y = a.clientPosition.Y + (0.1 * (serverPositionn.Y - a.clientPosition.Y))

	return nil
}
