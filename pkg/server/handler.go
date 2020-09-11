package server

import (
	"github.com/golang/protobuf/proto"
	"github.com/hueypark/marsettler/pkg/message"
)

// Handle handles message.
func Handle(id message.ID, bytes []byte) error {
	switch id {
	case message.PingID:
		ping := &message.Ping{}
		err := proto.Unmarshal(bytes, ping)
		if err != nil {
			return err
		}
	}

	return nil
}
