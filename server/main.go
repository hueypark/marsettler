package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hueypark/marsettler/core/id_generator"
	"github.com/hueypark/marsettler/core/net"
	"github.com/hueypark/marsettler/core/physics"
	"github.com/hueypark/marsettler/server/ctx"
	"github.com/hueypark/marsettler/server/game"
	"github.com/hueypark/marsettler/server/game/handler"
	"github.com/hueypark/marsettler/server/game/node_handler"
	"github.com/hueypark/marsettler/server/user"
)

func main() {
	id := flag.Int("id", 0, "id")
	host := flag.String("host", "", "host")
	port := flag.Int("port", 208, "port")
	aoe := flag.String("aoe", "-1000,1000,-1000,1000", "area of effect(Left,Right,Bottom,Top")
	nodes := flag.String("nodes", "", "list of other nodes to connect")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	id_generator.Init(*id)

	ctx.World = game.NewWorld(physics.NewAreaOfEffect(*aoe))

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

	address := *host + ":" + strconv.Itoa(*port)
	for _, node := range strings.Split(*nodes, ",") {
		if node == address ||
			node == "" {
			continue
		}

		net.NewClient(node, node_handler.Handle)
	}

	server := net.NewServer(
		user.OnAccept,
		user.OnClose,
		handler.Handle)
	server.Listen(address)
}
