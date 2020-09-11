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
		c, err := client.NewClient()
		if err != nil {
			log.Fatalln(err)
		}

		err = c.Run()
		if err != nil {
			log.Fatalln(err)
		}

		c.Close()
	}()

	wg.Add(1)
	go func() {
		s := server.NewServer()

		err := s.Run()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	wg.Wait()
}
