package game

import "github.com/jakecoffman/cp"

// Space represents physics space of game.
var Space *cp.Space

func init() {
	Space = cp.NewSpace()
	Space.Iterations = 1
}

// EachActor iterates all actors.
func EachActor(f func(actor *Actor)) {
	Space.EachBody(func(body *cp.Body) {
		actor := body.UserData.(*Actor)

		f(actor)
	})
}
