package server

import (
	"github.com/hueypark/marsettler/core/net"
	"github.com/hueypark/marsettler/server/game"
	"github.com/hueypark/marsettler/server/game/handler"
)

func Run() {
	game.NewWorld()

	server := net.NewServer(
		game.OnAccept,
		game.OnClose,
		handler.Handle)
	server.Listen(":8080")
}
