package main

import (
	"log"
	"sync"

	"github.com/hueypark/marsettler/pkg/client"
	"github.com/hueypark/marsettler/pkg/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		s := server.NewServer()

		err := s.Run()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	runClient()

	wg.Wait()
}

func runClient() {
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
