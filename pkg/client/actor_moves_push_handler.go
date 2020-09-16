package client

import (
	"log"

	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/shared"
)

// ActorMovesPushHandler handles message.ActorMovesPush.
func ActorMovesPushHandler(conn *shared.Conn, m *message.ActorMovesPush) error {
	for _, actorMove := range m.Moves {
		log.Println(actorMove.Id, actorMove.Position.X, actorMove.Position.Y)
	}

	return nil
}
