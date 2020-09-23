package bot

import (
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
)

// Bot is an automated bot.
type Bot struct {
	conn *net.Conn
}

// NewBot creates new bot.
func NewBot() (*Bot, error) {
	b := &Bot{}
	conn, err := connect()
	if err != nil {
		return nil, err
	}

	b.conn = conn

	return b, nil
}

// Run rus bot.
func (b *Bot) Run() error {
	signIn := &message.SignInRequest{}
	err := b.conn.Write(signIn)
	if err != nil {
		return err
	}

	delta := time.Second / 10
	deltaSeconds := delta.Seconds()
	ticker := time.NewTicker(delta)

	for range ticker.C {
		err := b.Tick(deltaSeconds)
		if err != nil {
			return err
		}
	}

	return nil
}

// Tick updates bot periodically.
func (b *Bot) Tick(delta float64) error {
	return nil
}

func connect() (*net.Conn, error) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	websocketConn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	conn, err := net.NewConn(websocketConn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
