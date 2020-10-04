package handler

import (
	"log"

	"github.com/hueypark/marsettler/pkg/client/game"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
)

// ActorDisappearsPushHandler handles message.ActorDisappearsPush.
func ActorDisappearsPushHandler(
	conn *net.Conn, m *message.ActorDisappearsPush, world *game.World,
) error {
	for _, disappear := range m.Disappears {
		err := world.DeleteActor(disappear.Id)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
