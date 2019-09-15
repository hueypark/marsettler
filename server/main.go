package main

import (
	"flag"
	"log"
	"time"

	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/core/net"
	"github.com/hueypark/marsettler/core/physics"
	"github.com/hueypark/marsettler/server/ctx"
	"github.com/hueypark/marsettler/server/game"
	"github.com/hueypark/marsettler/server/game/handler"
	"github.com/hueypark/marsettler/server/user"
)

func main() {
	id := flag.Int("id", 0, "id")
	host := flag.String("host", "", "host")
	port := flag.Int("port", 208, "port")
	aoe := flag.String("aoe", "-1000,1000,-1000,1000", "Area of effect(Left,Right,Bottom,Top")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	id_generator.Init(*id)

	ctx.World = game.NewWorld(physics.NewAreaOfEffect(*aoe))
	ctx.World.NewActor(id_generator.Generate(), vector.Zero())

	go func() {
		delta := time.Second / 60
		floatDelta := delta.Seconds()
		ticker := time.NewTicker(delta)

		for {
			select {
			case <-ticker.C:
				ctx.World.Tick(floatDelta)
			}
		}
	}()

	server := net.NewServer(
		user.OnAccept,
		user.OnClose,
		handler.Handle)
	server.Listen(*host, *port)
}
