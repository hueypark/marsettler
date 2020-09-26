package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hueypark/marsettler/pkg/internal/game"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"golang.org/x/image/colornames"
)

// Actor is basic object in world.
type Actor struct {
	*game.Actor
	clientPosition *math2d.Vector
}

// NewActor Creates new actor.
func NewActor(id int64, position *math2d.Vector) *Actor {
	a := &Actor{
		Actor:          game.NewActor(id),
		clientPosition: &math2d.Vector{},
	}

	a.clientPosition.Set(position)
	a.SetPosition(position)

	return a
}

// Draw implements ebiten.Game.Draw.
func (a *Actor) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, a.Position().X, a.Position().Y, 10, 10, colornames.Saddlebrown)
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
