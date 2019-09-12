package main

import (
	"github.com/hueypark/marsettler/core/net"
	"github.com/hueypark/marsettler/server/game"
	"github.com/hueypark/marsettler/server/game/handler"
	"time"
)

func main() {
	world := game.NewWorld()

	go func() {
		ticker := time.NewTicker(time.Second / 60)

		for {
			select {
			case <- ticker.C:
			world.Tick()
			}
		}
	}()

	server := net.NewServer(
		game.OnAccept,
		game.OnClose,
		handler.Handle)
	server.Listen(":8080")
}
