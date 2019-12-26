package data

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/pkg/asset"
)

func Actor(id int) *ActorData {
	return actors[id]
}

type ActorData struct {
	Name         string
	Abbreviation string
	Image        *ebiten.Image
	Radius       float64
	BehaviorTree string
}

var actors = map[int]*ActorData{
	1: {
		"CityHall",
		"CH",
		asset.CityHall,
		10,
		`Name: Sequence
Children:
- Name: Wait
  WaitTick: 60
  Tick: 0
- Name: CreateActor
  ActorID: 2
`,
	},
	2: {
		"Worker",
		"W",
		asset.Worker,
		10,
		`Name: Sequence
Children:
- Name: BlackboardCondition
  Conditions:
  - Name: NotHasKey
    Key: 0
  Child:
    Name: FindPath
- Name: MoveTo
  Path: []
  MoveWaitTime: 60
  RemainMoveWaitTime: 0
`,
	},
}
