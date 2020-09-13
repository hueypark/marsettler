package client

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/shared"
)

// Client is the marsettler client.
type Client struct {
	conn *shared.Conn
}

// NewClient creates new client.
func NewClient() (*Client, error) {
	c := &Client{}

	websocketConn, err := connect()
	if err != nil {
		return nil, err
	}

	conn, err := shared.NewConn(
		websocketConn,
		shared.HandlerFuncs{
			message.PongID: func(conn *shared.Conn, m *message.Pong) error {
				log.Println("Pong")
				return nil
			},
		})
	if err != nil {
		return nil, err
	}

	c.conn = conn

	return c, nil
}

// Close closes client.
func (c *Client) Close() {
	c.conn.Close()
}

// Run runs client.
func (c *Client) Run() error {
	go c.conn.Run()

	for {
		ping := &message.Ping{}
		err := c.conn.Write(ping)
		if err != nil {
			return err
		}

		c.conn.Consume()

		time.Sleep(time.Second)
	}
}

func connect() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
