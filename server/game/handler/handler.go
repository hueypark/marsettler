package handler

import (
	"fmt"
	"net"

	"github.com/hueypark/marsettler/server/game"
	"github.com/hueypark/marsettler/server/game/message"
	"github.com/hueypark/marsettler/server/game/message/fbs"
)

// Handle handle message
func Handle(userID int64, conn net.Conn) error {
	user := game.GetUser(userID)
	if user == nil {
		return fmt.Errorf("user is nil id: %d", userID)
	}

	messageID, body, err := message.ReadMessage(conn)
	if err != nil {
		return err
	}

	switch messageID {
	case fbs.LoginID:
		login := message.NewLogin(body)
		return handleLogin(user, login)
	}

	return fmt.Errorf("undefined message id: %d", messageID)
}
