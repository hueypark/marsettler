package main

import (
	"flag"
	"log"
	"time"

	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/net"
	"github.com/hueypark/marsettler/server/game"
	"github.com/hueypark/marsettler/server/game/handler"
)

func main() {
	id := flag.Int("id", 0, "id")
	host := flag.String("host", "", "host")
	port := flag.Int("port", 208, "port")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	id_generator.Init(*id)

	world := game.NewWorld()
	world.NewActor(id_generator.Generate(), vector.Zero())

	go func() {
		delta := time.Second / 60
		floatDelta := delta.Seconds()
		ticker := time.NewTicker(delta)

		for {
			select {
			case <-ticker.C:
				world.Tick(floatDelta)
			}
		}
	}()

	server := net.NewServer(
		game.OnAccept,
		game.OnClose,
		handler.Handle)
	server.Listen(*host, *port)
}
