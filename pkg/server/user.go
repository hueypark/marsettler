package server

import "github.com/hueypark/marsettler/pkg/game"

// User represents user.
type User struct {
	actor *game.Actor
}

// NewUser creates new user.
func NewUser() *User {
	u := &User{}

	return u
}

// SetActor sets actor.
func (u *User) SetActor(actor *game.Actor) {
	u.actor = actor
}
