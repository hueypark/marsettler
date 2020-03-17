package data

import "github.com/hueypark/marsettler/data/ai"

type ActorID int

const (
	Hero ActorID = iota
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
