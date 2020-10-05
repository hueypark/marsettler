// This file was generated from `./pkg/cmd/generate/generate_message_handler.go`.

package net

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hueypark/marsettler/pkg/message"
)

// Handler is message handler.
type Handler struct {
	actResponseHandler func(*Conn, *message.ActResponse) error
	actRequestHandler func(*Conn, *message.ActRequest) error
	actorHandler func(*Conn, *message.Actor) error
	actorDisappearHandler func(*Conn, *message.ActorDisappear) error
	actorDisappearsPushHandler func(*Conn, *message.ActorDisappearsPush) error
	actorMoveHandler func(*Conn, *message.ActorMove) error
	actorMovesPushHandler func(*Conn, *message.ActorMovesPush) error
	actorsPushHandler func(*Conn, *message.ActorsPush) error
	moveStickRequestHandler func(*Conn, *message.MoveStickRequest) error
	moveToPositionRequestHandler func(*Conn, *message.MoveToPositionRequest) error
	signInRequestHandler func(*Conn, *message.SignInRequest) error
	signInResponseHandler func(*Conn, *message.SignInResponse) error
	vectorHandler func(*Conn, *message.Vector) error
}

// HandlerFuncs represents handler functions.
type HandlerFuncs map[message.ID]interface{}

// NewHandler creates new handler.
func NewHandler(handlers HandlerFuncs) (*Handler, error) {
	h := &Handler{}

	for id, handler := range handlers {
		switch id {
		case message.ActResponseID:
			v, ok := handler.(func(*Conn, *message.ActResponse) error)
			if !ok {
				return nil, errors.New("handler does not handles ActResponse")
			}

			h.actResponseHandler = v
		case message.ActRequestID:
			v, ok := handler.(func(*Conn, *message.ActRequest) error)
			if !ok {
				return nil, errors.New("handler does not handles ActRequest")
			}

			h.actRequestHandler = v
		case message.ActorID:
			v, ok := handler.(func(*Conn, *message.Actor) error)
			if !ok {
				return nil, errors.New("handler does not handles Actor")
			}

			h.actorHandler = v
		case message.ActorDisappearID:
			v, ok := handler.(func(*Conn, *message.ActorDisappear) error)
			if !ok {
				return nil, errors.New("handler does not handles ActorDisappear")
			}

			h.actorDisappearHandler = v
		case message.ActorDisappearsPushID:
			v, ok := handler.(func(*Conn, *message.ActorDisappearsPush) error)
			if !ok {
				return nil, errors.New("handler does not handles ActorDisappearsPush")
			}

			h.actorDisappearsPushHandler = v
		case message.ActorMoveID:
			v, ok := handler.(func(*Conn, *message.ActorMove) error)
			if !ok {
				return nil, errors.New("handler does not handles ActorMove")
			}

			h.actorMoveHandler = v
		case message.ActorMovesPushID:
			v, ok := handler.(func(*Conn, *message.ActorMovesPush) error)
			if !ok {
				return nil, errors.New("handler does not handles ActorMovesPush")
			}

			h.actorMovesPushHandler = v
		case message.ActorsPushID:
			v, ok := handler.(func(*Conn, *message.ActorsPush) error)
			if !ok {
				return nil, errors.New("handler does not handles ActorsPush")
			}

			h.actorsPushHandler = v
		case message.MoveStickRequestID:
			v, ok := handler.(func(*Conn, *message.MoveStickRequest) error)
			if !ok {
				return nil, errors.New("handler does not handles MoveStickRequest")
			}

			h.moveStickRequestHandler = v
		case message.MoveToPositionRequestID:
			v, ok := handler.(func(*Conn, *message.MoveToPositionRequest) error)
			if !ok {
				return nil, errors.New("handler does not handles MoveToPositionRequest")
			}

			h.moveToPositionRequestHandler = v
		case message.SignInRequestID:
			v, ok := handler.(func(*Conn, *message.SignInRequest) error)
			if !ok {
				return nil, errors.New("handler does not handles SignInRequest")
			}

			h.signInRequestHandler = v
		case message.SignInResponseID:
			v, ok := handler.(func(*Conn, *message.SignInResponse) error)
			if !ok {
				return nil, errors.New("handler does not handles SignInResponse")
			}

			h.signInResponseHandler = v
		case message.VectorID:
			v, ok := handler.(func(*Conn, *message.Vector) error)
			if !ok {
				return nil, errors.New("handler does not handles Vector")
			}

			h.vectorHandler = v
		}
	}

	return h, nil
}

// Handle handles message.
func (h *Handler) Handle(conn *Conn, id message.ID, bytes []byte) error {
	switch id {
	case message.ActResponseID:
		m := &message.ActResponse{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.actResponseHandler == nil {
			return nil
		}

		return h.actResponseHandler(conn, m)
	case message.ActRequestID:
		m := &message.ActRequest{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.actRequestHandler == nil {
			return nil
		}

		return h.actRequestHandler(conn, m)
	case message.ActorID:
		m := &message.Actor{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.actorHandler == nil {
			return nil
		}

		return h.actorHandler(conn, m)
	case message.ActorDisappearID:
		m := &message.ActorDisappear{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.actorDisappearHandler == nil {
			return nil
		}

		return h.actorDisappearHandler(conn, m)
	case message.ActorDisappearsPushID:
		m := &message.ActorDisappearsPush{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.actorDisappearsPushHandler == nil {
			return nil
		}

		return h.actorDisappearsPushHandler(conn, m)
	case message.ActorMoveID:
		m := &message.ActorMove{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.actorMoveHandler == nil {
			return nil
		}

		return h.actorMoveHandler(conn, m)
	case message.ActorMovesPushID:
		m := &message.ActorMovesPush{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.actorMovesPushHandler == nil {
			return nil
		}

		return h.actorMovesPushHandler(conn, m)
	case message.ActorsPushID:
		m := &message.ActorsPush{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.actorsPushHandler == nil {
			return nil
		}

		return h.actorsPushHandler(conn, m)
	case message.MoveStickRequestID:
		m := &message.MoveStickRequest{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.moveStickRequestHandler == nil {
			return nil
		}

		return h.moveStickRequestHandler(conn, m)
	case message.MoveToPositionRequestID:
		m := &message.MoveToPositionRequest{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.moveToPositionRequestHandler == nil {
			return nil
		}

		return h.moveToPositionRequestHandler(conn, m)
	case message.SignInRequestID:
		m := &message.SignInRequest{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.signInRequestHandler == nil {
			return nil
		}

		return h.signInRequestHandler(conn, m)
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
	case message.VectorID:
		m := &message.Vector{}
		err := proto.Unmarshal(bytes, m)
		if err != nil {
			return err
		}

		if h.vectorHandler == nil {
			return nil
		}

		return h.vectorHandler(conn, m)
	}

	return errors.New(fmt.Sprintf("unhandled id: %v", id))
}
