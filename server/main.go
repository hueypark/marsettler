package main

import (
	"log"

	"github.com/hueypark/marsettler/server/pkg/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s := server.NewServer()

	err := s.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
