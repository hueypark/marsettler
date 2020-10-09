package handler

import "github.com/hueypark/marsettler/pkg/client/game"

type client interface {
	SetMyActor(actor *game.Actor)
}
