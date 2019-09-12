package ctx

import (
	"github.com/hueypark/marsettler/client/core/ui"
	"github.com/hueypark/marsettler/client/game"
)

func init() {
	World = game.NewWorld()
}

var (
	Cursor ui.Cursor
	World  *game.World
)
