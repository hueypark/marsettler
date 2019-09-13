package ctx

import (
	"github.com/hueypark/marsettler/client/core/ui"
	"github.com/hueypark/marsettler/client/game"
	"github.com/hueypark/marsettler/core/net"
)

var (
	Client *net.Client
	Cursor ui.Cursor
	World  *game.World
)
