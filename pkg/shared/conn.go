package shared

import (
	"encoding/binary"
	"fmt"
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/hueypark/marsettler/pkg/message"
)

// Conn reprents a connection.
type Conn struct {
	closeChan chan bool
	conn      *websocket.Conn
	handler   *message.Handler
	messages  []rawMessage
	mux       sync.Mutex
}

type rawMessage struct {
	ID    message.ID
	Bytes []byte
}

// NewConn creates new connection.
func NewConn(conn *websocket.Conn, handlerFuncs message.HandlerFuncs) (*Conn, error) {
	c := &Conn{
		conn:      conn,
		closeChan: make(chan bool),
	}

	handler, err := message.NewHandler(handlerFuncs)
	if err != nil {
		return nil, err
	}

	c.handler = handler

	return c, nil
}

// Close closes connection.
func (c *Conn) Close() {
	c.closeChan <- true
	_ = c.conn.Close()
}

// Consume consumes messages.
func (c *Conn) Consume() {
	c.mux.Lock()
	defer c.mux.Unlock()

	for _, m := range c.messages {
		err := c.handler.Handle(m.ID, m.Bytes)
		if err != nil {
			log.Println(fmt.Sprintf("id: %v, err: %v", m.ID, err))
		}
	}

	c.messages = nil
}

// Run runs connection.
func (c *Conn) Run() {
	for {
		select {
		case <-c.closeChan:
			return
		default:
			func() {
				_, idBytes, err := c.conn.ReadMessage()
				if err != nil {
					log.Println(err)
					return
				}

				_, bytes, err := c.conn.ReadMessage()
				if err != nil {
					log.Println(err)
					return
				}

				id := message.ID(binary.LittleEndian.Uint32(idBytes))

				c.mux.Lock()
				defer c.mux.Unlock()
				c.messages = append(c.messages, rawMessage{id, bytes})
			}()
		}
	}
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
