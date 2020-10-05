package client

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/pkg/client/game"
	"github.com/hueypark/marsettler/pkg/client/handler"
	"github.com/hueypark/marsettler/pkg/internal/math2d"
	"github.com/hueypark/marsettler/pkg/internal/net"
	"github.com/hueypark/marsettler/pkg/message"
)

// Client is the client client.
type Client struct {
	useRenderer bool

	id        int64
	conn      *net.Conn
	world     *game.World
	tickDelta time.Duration
	geoM      ebiten.GeoM
}

// NewClient creates new client.
func NewClient(useRenderer bool) (*Client, error) {
	c := &Client{}

	c.useRenderer = useRenderer

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
	c.tickDelta = time.Second / ebiten.DefaultTPS

	err = conn.SetHandlers(handler.Generate(c, c.world))
	if err != nil {
		return nil, err
	}

	c.geoM.Scale(1, -1)

	return c, nil
}

// Act does act.
func (c *Client) Act() error {
	m := &message.ActRequest{}

	var err error
	m.TargetId, err = c.world.NearestActorId(c.id)
	if err != nil {
		return err
	}

	err = c.conn.Write(m)
	if err != nil {
		return err
	}

	return nil
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

	if c.useRenderer {
		return ebiten.RunGame(c)
	} else {
		ticker := time.NewTicker(c.tickDelta)

		for range ticker.C {
			err := c.Tick(c.tickDelta.Seconds())
			if err != nil {
				return err
			}
		}
		return nil
	}
}

// Draw implements ebiten.Game.Draw.
func (c *Client) Draw(screen *ebiten.Image) {
	err := c.world.Draw(
		screen,
		func(a *game.Actor) ebiten.GeoM {
			m := ebiten.GeoM{}

			m.Scale(1, -1)

			m.Translate(a.Position().X, a.Position().Y)

			m.Scale(1, -1)

			return m
		})
	if err != nil {
		log.Println(err)
	}
}

// SetID sets id.
func (c *Client) SetID(id int64) {
	c.id = id
}

// Layout implements ebiten.Game.Layout.
func (c *Client) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 320, 240
}

// Tick updates actor periodically.
func (c *Client) Tick(delta float64) error {
	err := c.world.Tick(delta)
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

	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		err := c.Act()
		if err != nil {
			return err
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		moveToPosition := &message.MoveToPositionRequest{
			Position: &message.Vector{X: float64(x), Y: -float64(y)},
		}
		err := c.conn.Write(moveToPosition)
		if err != nil {
			return err
		}
	}

	err = c.updateMoveStickRequest()
	if err != nil {
		return err
	}

	err = c.conn.Consume()
	if err != nil {
		return err
	}

	return nil
}

// Update implements ebiten.Game.Update.
func (c *Client) Update(_ *ebiten.Image) error {
	return c.Tick(c.tickDelta.Seconds())
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
