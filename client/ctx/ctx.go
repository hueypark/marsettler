package ctx

import (
	"github.com/hueypark/marsettler/client/core/ui"
	"github.com/hueypark/marsettler/client/game"
	"github.com/hueypark/marsettler/core/net"
)

func init() {
	Clients = make(map[string]*net.Client)
}

var (
	Clients map[string]*net.Client
	Cursor  ui.Cursor
	World   *game.World
)
