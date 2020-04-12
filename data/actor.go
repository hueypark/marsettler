package data

import "github.com/hueypark/marsettler/data/ai"

type ActorID int

const (
	Hero ActorID = iota
	Tree
)

var actors map[ActorID]*ActorData

func init() {
	actors = map[ActorID]*ActorData{
		Hero: newActor(
			"Hero",
			"/asset/figures/hero.png",
			ai.Hero,
			300,
		),
		Tree: newActor(
			"Tree",
			"/asset/tiles_forest_conifer_dense_clear_green/0.png",
			ai.Tree,
			0),
	}

}

func Actor(id ActorID) *ActorData {
	return actors[id]
}

type ActorData struct {
	Name         string
	Image        string
	BehaviorTree string

	// MoveWaitTime represents the wait time for the move(millie seconds).
	MoveWaitTime int
}

func newActor(name string, image string, bt string, moveWaitTime int) *ActorData {
	data := &ActorData{
		Name:         name,
		Image:        image,
		BehaviorTree: bt,
		MoveWaitTime: moveWaitTime,
	}

	return data
}
