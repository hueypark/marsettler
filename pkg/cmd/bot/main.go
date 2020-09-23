package main

import (
	"log"

	"github.com/hueypark/marsettler/pkg/bot"
)

func main() {
	b, err := bot.NewBot()
	if err != nil {
		log.Fatalln(err)
	}

	err = b.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
