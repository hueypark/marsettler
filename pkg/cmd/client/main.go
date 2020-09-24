package main

import (
	"log"

	"github.com/hueypark/marsettler/pkg/client"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c, err := client.NewClient()
	if err != nil {
		log.Fatalln(err)
	}

	err = c.Run()
	if err != nil {
		log.Fatalln(err)
	}

	c.Close()
}
