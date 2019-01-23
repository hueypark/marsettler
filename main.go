package main

import (
	"github.com/hueypark/marsettler/core/net"
	"github.com/hueypark/marsettler/game"
	"github.com/hueypark/marsettler/game/handler"
)

func main() {
	server := net.NewServer(
		game.NewUser,
		handler.Handle)
	server.Listen(":8080")
}
