package net

import (
	"encoding/binary"
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/hueypark/marsettler/server/pkg/message/fbs"
)

// Conn reprents a connection.
type Conn struct {
	closeChan	chan bool
	closed		bool
	conn		*websocket.Conn
	handlers	*Handler
	messages	[]rawMessage
	mux		sync.Mutex
}

type rawMessage struct {
	ID	fbs.ID
	Bytes	[]byte
}

// NewConn creates new connection.
func NewConn(conn *websocket.Conn) (*Conn, error) {
	c := &Conn{
		conn:		conn,
		closeChan:	make(chan bool),
	}

	return c, nil
}

// Close closes connection.
func (c *Conn) Close() {
	c.closeChan <- true
	c.closed = true
	_ = c.conn.Close()
}

// Closed return close state.
func (c *Conn) Closed() bool {
	return c.closed
}

// Consume consumes messages.
func (c *Conn) Consume() error {
	c.mux.Lock()
	defer c.mux.Unlock()

	for _, m := range c.messages {
		err := c.handlers.Handle(c, m.ID, m.Bytes)
		if err != nil {
			// TODO(jaewan): Send error message to client.
			log.Printf("%+v", err)
		}
	}

	c.messages = nil

	return nil
}

// Run runs connection.
func (c *Conn) Run() error {
	for {
		select {
		case <-c.closeChan:
			return nil
		default:
			_, idBytes, err := c.conn.ReadMessage()
			if err != nil {
				return err
			}

			_, bytes, err := c.conn.ReadMessage()
			if err != nil {
				return err
			}

			id := fbs.ID(binary.LittleEndian.Uint32(idBytes))

			c.mux.Lock()
			c.messages = append(c.messages, rawMessage{id, bytes})
			c.mux.Unlock()
		}
	}
}

// SetHandlers sets handlers.
func (c *Conn) SetHandlers(handlerFuncs HandlerFuncs) error {
	handler, err := NewHandler(handlerFuncs)
	if err != nil {
		return err
	}

	c.handlers = handler

	return nil
}

// Send sends message.
func (c *Conn) Write(message fbs.Message) error {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	idBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(idBytes, uint32(message.ID()))

	err = c.conn.WriteMessage(websocket.BinaryMessage, idBytes)
	if err != nil {
		return err
	}

	return c.conn.WriteMessage(websocket.BinaryMessage, bytes)
}
