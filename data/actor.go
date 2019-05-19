package data

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/asset"
)

func Actor(id int) *ActorData {
	return actors[id]
}

type ActorData struct {
	Name         string
	Abbreviation string
	Image        *ebiten.Image
	BehaviorTree string
}

var actors = map[int]*ActorData{
	1: {
		"CityHall",
		"CH",
		asset.CityHall,
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
