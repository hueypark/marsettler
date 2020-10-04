package main

import (
	"flag"
	"log"

	"github.com/hueypark/marsettler/pkg/client"
)

func main() {
	useRenderer := flag.Bool("renderer", true, "use renderer")

	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c, err := client.NewClient(*useRenderer)
	if err != nil {
		log.Fatalln(err)
	}

	err = c.Run()
	if err != nil {
		log.Fatalln(err)
	}

	c.Close()
}
