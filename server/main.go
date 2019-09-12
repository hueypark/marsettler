package main

import (
	"time"

	"github.com/hueypark/marsettler/core/math/vector"

	"github.com/hueypark/marsettler/core/id_generator"

	"github.com/hueypark/marsettler/core/net"
	"github.com/hueypark/marsettler/server/game"
	"github.com/hueypark/marsettler/server/game/handler"
)

func main() {
	world := game.NewWorld()
	world.NewActor(id_generator.Generate(), vector.Zero())

	go func() {
		ticker := time.NewTicker(time.Second / 60)

		for {
			select {
			case <-ticker.C:
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
