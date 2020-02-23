package data

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/pkg/asset"
)

type ActorID int

const (
	Leader ActorID = iota
	Legionary
)

func Actor(id ActorID) *ActorData {
	return actors[id]
}

type ActorData struct {
	Name         string
	Image        *ebiten.Image
	Radius       float64
	BehaviorTree string
}

var actors = map[ActorID]*ActorData{
	Leader: {
		Name:   "Leader",
		Image:  asset.Circle,
		Radius: 10,
		BehaviorTree: `Sequence
	MoveTo: position
`,
	},
	Legionary: {
		Name:   "Legionary",
		Image:  asset.Circle,
		Radius: 10,
		BehaviorTree: `Sequence
	FindRandomPosition: patrolPosition
	MoveTo: patrolPosition
	Wait: 1
`,
	},
}
