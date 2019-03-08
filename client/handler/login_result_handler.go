package handler

import (
	"log"

	"github.com/hueypark/marsettler/server/game/message/fbs"
)

func handleLoginResult(loginResult *fbs.LoginResult) {
	log.Println(loginResult.Id())
}
