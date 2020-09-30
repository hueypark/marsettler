package client

import (
	"log"

	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
)

// ActResponseHandler handles message.ActResponse.
func ActResponseHandler(_ *net.Conn, m *message.ActResponse) error {
	log.Println(m)

	return nil
}
