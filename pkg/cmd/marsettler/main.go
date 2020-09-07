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

	go func() {
		wg.Add(1)

		c := client.NewClient()
		c.Run()
	}()

	go func() {
		wg.Add(1)

		s := server.NewServer()
		s.Run()
	}()

	wg.Wait()
}
