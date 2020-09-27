package net

import (
	"encoding/binary"
	"errors"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/hueypark/marsettler/pkg/message"
)

// Conn reprents a connection.
type Conn struct {
	closeChan chan bool
	conn      *websocket.Conn
	handlers  *Handler
	messages  []rawMessage
	mux       sync.Mutex
}

type rawMessage struct {
	ID    message.ID
	Bytes []byte
}

// NewConn creates new connection.
func NewConn(conn *websocket.Conn) (*Conn, error) {
	c := &Conn{
		conn:      conn,
		closeChan: make(chan bool),
	}

	return c, nil
}

// Close closes connection.
func (c *Conn) Close() {
	c.closeChan <- true
	_ = c.conn.Close()
}

// Consume consumes messages.
func (c *Conn) Consume() error {
	c.mux.Lock()
	defer c.mux.Unlock()

	select {
	case <-c.closeChan:
		return errors.New("closed connection")
	default:
		for _, m := range c.messages {
			err := c.handlers.Handle(c, m.ID, m.Bytes)
			if err != nil {
				return err
			}
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

			id := message.ID(binary.LittleEndian.Uint32(idBytes))

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
func (c *Conn) Write(message message.Message) error {
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
