package handler

import (
	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/message"
	"github.com/hueypark/marsettler/server/game"
	"github.com/hueypark/marsettler/server/game/message/fbs"
)

func handleLogin(user *game.User, login *fbs.Login) error {
	id := login.Id()
	if id == 0 {
		id = id_generator.Generate()
	}

	msgActorPush := &message.ActorsPush{}
	user.Send(msgActorPush)

	return nil
}
