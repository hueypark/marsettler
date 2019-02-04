package net

import (
	"log"
	"net"
)

// Client represents network client
type Client struct {
	conn net.Conn
}

// NewClient create new client
func NewClient(address string) *Client {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	return &Client{conn}
}

// Tick called when every tick in game world
func (client Client) Tick() {
}
