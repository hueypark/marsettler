package client

import (
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/core/math/vector"
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

	conn, err := shared.NewConn(websocketConn)
	if err != nil {
		return nil, err
	}

	err = conn.SetHandlers(shared.HandlerFuncs{
		message.ActorMovesPushID: ActorMovesPushHandler,
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
		signIn := &message.SignInRequest{}
		err := c.conn.Write(signIn)
		if err != nil {
			return err
		}
	}

	err := c.updateMoveStickRequest()
	if err != nil {
		return err
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

func (c *Client) updateMoveStickRequest() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) ||
		inpututil.IsKeyJustPressed(ebiten.KeyA) ||
		inpututil.IsKeyJustPressed(ebiten.KeyS) ||
		inpututil.IsKeyJustPressed(ebiten.KeyD) ||
		inpututil.IsKeyJustReleased(ebiten.KeyW) ||
		inpututil.IsKeyJustReleased(ebiten.KeyA) ||
		inpututil.IsKeyJustReleased(ebiten.KeyS) ||
		inpututil.IsKeyJustReleased(ebiten.KeyD) {

		var direction vector.Vector

		if ebiten.IsKeyPressed(ebiten.KeyW) {
			direction.Add(&vector.Vector{X: 0, Y: 1})
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) {
			direction.Add(&vector.Vector{X: -1, Y: 0})
		}

		if ebiten.IsKeyPressed(ebiten.KeyS) {
			direction.Add(&vector.Vector{X: 0, Y: -1})
		}

		if ebiten.IsKeyPressed(ebiten.KeyD) {
			direction.Add(&vector.Vector{X: 1, Y: 0})
		}

		direction.Normalize()

		moveStick := &message.MoveStickRequest{}
		moveStick.Direction = &message.Vector{
			X: direction.X,
			Y: direction.Y,
		}
		err := c.conn.Write(moveStick)
		if err != nil {
			return err
		}
	}

	return nil
}
