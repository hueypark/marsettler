package net

import (
	"log"
	"net"
)

// Client represents network client
type Client struct {
	Conn    net.Conn
	handler func(conn net.Conn) error
}

// NewClient create new client
func NewClient(address string, handler func(conn net.Conn) error) *Client {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	client := &Client{
		Conn:    conn,
		handler: handler}

	go client.handle()

	return client
}

func (client *Client) handle() {
	for {
		err := client.handler(client.Conn)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
