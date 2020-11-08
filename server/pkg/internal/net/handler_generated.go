// This file was generated from `./pkg/cmd/generate/generate_message_handler.go`.

package net

import (
	"errors"
	"fmt"
	"github.com/hueypark/marsettler/server/pkg/message/fbs"
)

// Handler is message handler.
type Handler struct {
	signInRequestHandler func(*Conn, *fbs.SignInRequest) error
}

// HandlerFuncs represents handler functions.
type HandlerFuncs map[fbs.ID]interface{}

// NewHandler creates new handler.
func NewHandler(handlers HandlerFuncs) (*Handler, error) {
	h := &Handler{}

	for id, handler := range handlers {
		switch id {
		case fbs.SignInRequestID:
			v, ok := handler.(func(*Conn, *fbs.SignInRequest) error)
			if !ok {
				return nil, errors.New("handler does not handles SignInRequest")
			}

			h.signInRequestHandler = v
		}
	}

	return h, nil
}

// Handle handles message.
func (h *Handler) Handle(conn *Conn, id fbs.ID, bytes []byte) error {
	switch id {
	case fbs.SignInRequestID:
		if h.signInRequestHandler == nil {
			return nil
		}

		return h.signInRequestHandler(conn, fbs.GetRootAsSignInRequest(bytes, 0))
	}

	return errors.New(fmt.Sprintf("unhandled id: %v", id))
}
