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
	Image           *ebiten.Image
	NewBehaviorTree func(actor task.Actor) *behavior_tree.BehaviorTree
}

var actors = map[int64]*Actor{
	1: {
		asset.CityHall,
		ai.NewCityHall,
	},
	2: {
		asset.Worker,
		ai.NewWorker,
	},
}
