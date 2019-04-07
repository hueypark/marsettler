package data

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hueypark/marsettler/client/asset"
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/server/game/ai"
	"github.com/hueypark/marsettler/server/game/ai/task"
)

func GetActor(id int64) *Actor {
	return actors[id]
}

type Actor struct {
	Name            string
	Image           *ebiten.Image
	NewBehaviorTree func(actor task.Actor) *behavior_tree.BehaviorTree
}

var actors = map[int64]*Actor{
	1: {
		"CityHall",
		asset.CityHall,
		ai.NewCityHall,
	},
	2: {
		"Worker",
		asset.Worker,
		ai.NewWorker,
	},
	3: {
		"Blueberry",
		asset.Blueberry,
		ai.NewNil,
	},
	100000: {
		"Fairy",
		asset.Fairy,
		ai.NewFairy,
	},
}
