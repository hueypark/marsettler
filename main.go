package main

import "github.com/hueypark/marsettler/core/net"

func main() {
	server := net.NewServer()
	server.Listen(":8080")
}
