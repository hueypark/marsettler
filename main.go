package main

import (
	"github.com/hueypark/marsettler/client"
	"github.com/hueypark/marsettler/server"
)

func main() {
	go func(){server.Run()}()

	client.Run()
}
