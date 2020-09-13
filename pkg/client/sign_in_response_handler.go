package client

import (
	"log"

	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/shared"
)

// SignInResponseHandler handles message.SignInResponse.
func SignInResponseHandler(conn *shared.Conn, m *message.SignInResponse) error {
	log.Println(m.Id)

	return nil
}
