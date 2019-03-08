package handler

import (
	"fmt"

	"github.com/hueypark/marsettler/core/net"
	"github.com/hueypark/marsettler/server/game/message"
	"github.com/hueypark/marsettler/server/game/message/fbs"
)

// Handle handles message.
func Handle(iClient interface{}) error {
	client := iClient.(*net.Client)

	id, body, err := message.ReadMessage(client.Conn())
	if err != nil {
		return err
	}

	switch id {
	case fbs.LoginResultID:
		loginResult := message.NewLoginResult(body)
		handleLoginResult(loginResult)
	case fbs.NodeID:
		node := message.NewNode(body)
		handleNode(node)
	default:
		return fmt.Errorf("unhandled message id: %d", id)
	}

	return nil
}
