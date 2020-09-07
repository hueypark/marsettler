package client

import (
	"bufio"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

// Client is the marsettler client.
type Client struct {
	conn *websocket.Conn
}

// NewClient creates new client.
func NewClient() *Client {
	c := &Client{
		conn: connect(),
	}

	return c
}

// Close closes client.
func (c *Client) Close() error {
	return c.conn.Close()
}

// Run runs client.
func (c *Client) Run() error {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		err = c.conn.WriteMessage(websocket.BinaryMessage, []byte(text))
		if err != nil {
			return err
		}
	}
}

func connect() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			log.Println(string(message))
		}
	}()

	return conn
}
