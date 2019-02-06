package net

import (
	"log"
	"net"
)

// Client represents network client
type Client struct {
	conn    net.Conn
	address string
	handler func(client interface{}) error
}

// NewClient create new client
func NewClient(address string, handler func(client interface{}) error) *Client {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
	}

	client := &Client{
		conn:    conn,
		address: address,
		handler: handler}

	go client.handle()

	return client
}

// Conn returns network connection.
func (client *Client) Conn() net.Conn {
	return client.conn
}

func (client *Client) handle() {
	for {
		err := client.handler(client)
		if err != nil {
			log.Println(err)
		}
	}
}
