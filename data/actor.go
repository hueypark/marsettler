package data

type ActorID int

const (
	Hero ActorID = iota
	Legionary
)

func Actor(id ActorID) *ActorData {
	return actors[id]
}

type ActorData struct {
	Name         string
	Image        string
	Radius       float64
	BehaviorTree string
}

var actors = map[ActorID]*ActorData{
	Hero: {
		Name:   "Hero",
		Image:  "/asset/figures/hero.png",
		Radius: 10,
		BehaviorTree: `Sequence
	MoveTo: position
`,
	},
	Legionary: {
		Name:   "Legionary",
		Image:  "node/legionary",
		Radius: 10,
		BehaviorTree: `Sequence
	FindRandomPosition: patrolPosition
	MoveTo: patrolPosition
	Wait: 1
`,
	},
}
