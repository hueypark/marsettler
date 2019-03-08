package message

import "net"

type client interface {
	Conn() net.Conn
}
