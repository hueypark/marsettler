package main

import (
	"log/slog"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hueypark/marsettler/pkg/lumberjack"
)

func main() {
	ebiten.SetWindowTitle("Lumberjack")

	err := lumberjack.NewGame().Run()
	if err != nil {
		slog.Error("Game failed.", "error", err)
		return
	}

	logger.Info("Game finished.")
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
