package main

import (
	_ "github.com/hueypark/marsettler/pkg/ai"
	"github.com/hueypark/marsettler/pkg/game"
)

func main() {
	g := game.New()
	g.Run()
}
