// This file was generated from `./pkg/cmd/generate/generate_message_handler.go`.

package shared

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hueypark/marsettler/pkg/message"
)

// Handler is message handler.
type Handler struct {
	pingHandler func(*Conn, *message.Ping) error
	pongHandler func(*Conn, *message.Pong) error
	signInHandler func(*Conn, *message.SignIn) error
	signInResponseHandler func(*Conn, *message.SignInResponse) error
}

// HandlerFuncs represents handler functions.
type HandlerFuncs map[message.ID]interface{}

// NewHandler creates new handler.
func NewHandler(handlers HandlerFuncs) (*Handler, error) {
	h := &Handler{}

	for id, handler := range handlers {
		switch id {
		case message.PingID:
			v, ok := handler.(func(*Conn, *message.Ping) error)
			if !ok {
				return nil, errors.New("handler does not handles Ping")
			}

			h.pingHandler = v
		case message.PongID:
			v, ok := handler.(func(*Conn, *message.Pong) error)
			if !ok {
				return nil, errors.New("handler does not handles Pong")
			}

			h.pongHandler = v
		case message.SignInID:
			v, ok := handler.(func(*Conn, *message.SignIn) error)
			if !ok {
				return nil, errors.New("handler does not handles SignIn")
			}

			h.signInHandler = v
		case message.SignInResponseID:
			v, ok := handler.(func(*Conn, *message.SignInResponse) error)
			if !ok {
				return nil, errors.New("handler does not handles SignInResponse")
			}

			h.signInResponseHandler = v
		}
	}

	return h, nil
}

// Handle handles message.
func (h *Handler) Handle(conn *Conn, id message.ID, bytes []byte) error {
	switch id {
	case message.PingID:
		m := &message.Ping{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.pingHandler == nil {
			return nil
		}

		return h.pingHandler(conn, m)
	case message.PongID:
		m := &message.Pong{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.pongHandler == nil {
			return nil
		}

		return h.pongHandler(conn, m)
	case message.SignInID:
		m := &message.SignIn{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.signInHandler == nil {
			return nil
		}

		return h.signInHandler(conn, m)
	case message.SignInResponseID:
		m := &message.SignInResponse{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.signInResponseHandler == nil {
			return nil
		}

		return h.signInResponseHandler(conn, m)
	}

	return errors.New(fmt.Sprintf("unhandled id: %v", id))
}
