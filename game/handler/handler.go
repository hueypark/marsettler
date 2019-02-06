package handler

import (
	"fmt"

	"github.com/hueypark/marsettler/game"
	"github.com/hueypark/marsettler/game/message"
	"github.com/hueypark/marsettler/game/message/fbs"
)

// Handle handle message
func Handle(iUser interface{}) error {
	user := iUser.(*game.User)

	id, body, err := message.ReadMessage(user)
	if err != nil {
		return err
	}
	switch id {
	case fbs.LoginID:
		login := message.MakeLogin(body)
		return handleLogin(user, login)
	}

	return fmt.Errorf("undefined message id: %d", id)
}
