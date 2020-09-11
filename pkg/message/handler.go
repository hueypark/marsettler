package message

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
)

// Handler is message handler.
type Handler struct {
	pingHandler func(*Ping) error
	pongHandler func(*Pong) error
}

// HandlerFuncs represents handler functions.
type HandlerFuncs map[ID]interface{}

// NewHandler creates new handler.
func NewHandler(handlers HandlerFuncs) (*Handler, error) {
	h := &Handler{}

	for id, handler := range handlers {
		switch id {
		case PingID:
			v, ok := handler.(func(*Ping) error)
			if !ok {
				return nil, errors.New("handler not handles ping")
			}

			h.pingHandler = v
		case PongID:
			v, ok := handler.(func(*Pong) error)
			if !ok {
				return nil, errors.New("handler not handles pong")
			}

			h.pongHandler = v
		}
	}

	return h, nil
}

// Handle handles message.
func (h *Handler) Handle(id ID, bytes []byte) error {
	switch id {
	case PingID:
		m := &Ping{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.pingHandler == nil {
			return nil
		}

		return h.pingHandler(m)
	case PongID:
		m := &Pong{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.pongHandler == nil {
			return nil
		}

		return h.pongHandler(m)
	}

	return errors.New(fmt.Sprintf("unhandled id: %v", id))
}
