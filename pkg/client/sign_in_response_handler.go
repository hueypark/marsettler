package client

import (
	"fmt"
	"log"

	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/shared"
)

// SignInResponseHandler handles message.SignInResponse.
func SignInResponseHandler(conn *shared.Conn, m *message.SignInResponse) error {
	log.Println(fmt.Sprintf("I am %v and I'm in (%v, %v)", m.Id, m.Position.X, m.Position.Y))

	return nil
}
