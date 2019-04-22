package data

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/asset"
)

func GetActor(id int) *Actor {
	return actors[id]
}

type Actor struct {
	Name         string
	Image        *ebiten.Image
	BehaviorTree string
}

var actors = map[int]*Actor{
	1: {
		"CityHall",
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
	3: {
		"Blueberry",
		asset.Blueberry,
		``,
	},
	100000: {
		"Fairy",
		asset.Fairy,
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
- Name: Wait
  WaitTick: 60
  Tick: 0
- Name: CreateActor
  ActorID: 3
`,
	},
}
