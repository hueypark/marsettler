package client

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/pkg/message"
	"github.com/hueypark/marsettler/pkg/shared"
)

// Client is the marsettler client.
type Client struct {
	conn *shared.Conn
}

// NewClient creates new client.
func NewClient() (*Client, error) {
	c := &Client{}

	websocketConn, err := connect()
	if err != nil {
		return nil, err
	}

	conn, err := shared.NewConn(
		websocketConn,
		shared.HandlerFuncs{
			message.PongID: func(conn *shared.Conn, m *message.Pong) error {
				log.Println("Pong")
				return nil
			},
			message.SignInResponseID: SignInResponseHandler,
		})
	if err != nil {
		return nil, err
	}

	c.conn = conn

	return c, nil
}

// Close closes client.
func (c *Client) Close() {
	c.conn.Close()
}

// Run runs client.
func (c *Client) Run() error {
	go c.conn.Run()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Marsettler")
	ebiten.SetRunnableOnUnfocused(true)
	return ebiten.RunGame(c)
}

// Draw implements ebiten.Game.Draw.
func (c *Client) Draw(screen *ebiten.Image) {
	_ = ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Draw implements ebiten.Game.Layout.
func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// Draw implements ebiten.Game.Update.
func (c *Client) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustReleased(ebiten.KeyEnter) {
		signIn := &message.SignIn{}
		err := c.conn.Write(signIn)
		if err != nil {
			return err
		}
	}

	c.conn.Consume()

	return nil
}

func connect() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
