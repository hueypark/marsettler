package net

import (
	"encoding/binary"
	"log"
	"net"

	"github.com/gogo/protobuf/proto"
	"github.com/hueypark/marsettler/message"
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

// Send sends message.
func (client *Client) Send(msg message.Msg) {
	if client.conn == nil {
		log.Println("conn is nil")
		return
	}

	id := msg.MsgID()
	size := msg.Size()
	buffer, err := proto.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	head := make([]byte, message.HeadSize)
	binary.LittleEndian.PutUint32(head[0:], uint32(id))
	binary.LittleEndian.PutUint32(head[4:], uint32(size))

	_, err = client.conn.Write(head)
	if err != nil {
		log.Println(err)
	}

	_, err = client.conn.Write(buffer)
	if err != nil {
		log.Println(err)
	}
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
			client.conn = nil
		}
	}
}
