package data

type ActorID int32

const (
	UserID ActorID = iota
	SwordSkillID
)

var actors map[ActorID]*ActorData

func init() {
	actors = map[ActorID]*ActorData{
		UserID: {
			"User",
			"circle",
			25,
		},
		SwordSkillID: {
			"SwordSkill",
			"circle",
			2,
		},
	}

}

func Actor(id ActorID) *ActorData {
	return actors[id]
}

type ActorData struct {
	Name   string
	Image  string
	Radius float64
}
