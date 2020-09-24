package client

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/pkg/client/game"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
)

// Client is the client client.
type Client struct {
	conn      *net.Conn
	world     *game.World
	tickDelta float64
}

// NewClient creates new client.
func NewClient() (*Client, error) {
	c := &Client{}

	websocketConn, err := connect()
	if err != nil {
		return nil, err
	}

	conn, err := net.NewConn(websocketConn)
	if err != nil {
		return nil, err
	}

	c.conn = conn
	c.world = game.NewWorld()
	c.tickDelta = 1.0 / ebiten.DefaultTPS

	err = conn.SetHandlers(net.HandlerFuncs{
		message.ActorMovesPushID: func(conn *net.Conn, m *message.ActorMovesPush) error {
			return ActorMovesPushHandler(conn, m, c.world)
		},
		message.ActorsPushID: func(conn *net.Conn, m *message.ActorsPush) error {
			return ActorsPushHandler(conn, m, c.world)
		},
		message.SignInResponseID: func(conn *net.Conn, m *message.SignInResponse) error {
			return SignInResponseHandler(conn, m, c.world)
		},
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Close closes client.
func (c *Client) Close() {
	c.conn.Close()
}

// Run runs client.
func (c *Client) Run() error {
	go func() {
		err := c.conn.Run()
		if err != nil {
			log.Println(err)
		}
	}()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Marsettler")
	ebiten.SetRunnableOnUnfocused(true)
	return ebiten.RunGame(c)
}

// Draw implements ebiten.Game.Draw.
func (c *Client) Draw(screen *ebiten.Image) {
	c.world.Draw(screen)
}

// Layout implements ebiten.Game.Layout.
func (c *Client) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 320, 240
}

// Update implements ebiten.Game.Update.
func (c *Client) Update(_ *ebiten.Image) error {
	err := c.world.Tick(c.tickDelta)
	if err != nil {
		return err
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyEnter) {
		signIn := &message.SignInRequest{}
		err := c.conn.Write(signIn)
		if err != nil {
			return err
		}
	}

	err = c.updateMoveStickRequest()
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

		var direction math2d.Vector

		if ebiten.IsKeyPressed(ebiten.KeyW) {
			direction.Add(&math2d.Vector{X: 0, Y: 1})
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) {
			direction.Add(&math2d.Vector{X: -1, Y: 0})
		}

		if ebiten.IsKeyPressed(ebiten.KeyS) {
			direction.Add(&math2d.Vector{X: 0, Y: -1})
		}

		if ebiten.IsKeyPressed(ebiten.KeyD) {
			direction.Add(&math2d.Vector{X: 1, Y: 0})
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
