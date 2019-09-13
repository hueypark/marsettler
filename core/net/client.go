package net

import (
	"log"
	"net"
)

// Client represents network client
type Client struct {
	conn    net.Conn
	address string
	handler func(conn net.Conn) error
}

// NewClient create new client
func NewClient(address string, handler func(conn net.Conn) error) *Client {
	client := &Client{
		address: address,
		handler: handler,
	}

	go client.handle()

	return client
}

// Conn returns network connection.
func (client *Client) Conn() net.Conn {
	return client.conn
}

func (client *Client) handle() {
	for {
		if client.conn == nil {
			conn, err := net.Dial("tcp", client.address)
			if err != nil {
				log.Println(err)
				continue
			}

			client.conn = conn
			continue
		}

		err := client.handler(client.conn)
		if err != nil {
			log.Println(err)
		}
	}
}
