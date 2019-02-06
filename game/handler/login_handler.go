package handler

import (
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/game"
	"github.com/hueypark/marsettler/game/message/fbs"
)

func handleLogin(user *game.User, login *fbs.Login) error {
	id := login.Id()
	if id == 0 {
		id = id_generator.Generate()
	}

	user.SendLoginResult(id)

	return nil
}
