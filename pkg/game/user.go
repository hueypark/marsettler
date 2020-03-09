package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/core/asset"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/data"
	"github.com/hueypark/marsettler/pkg/renderer"
)

type User struct {
	actor       *Actor
	world       *World
	clickedImg  *ebiten.Image
	clickedNode *Node
}

func NewUser(world *World) *User {
	user := &User{
		actor:       world.NewActor(data.Hero),
		world:       world,
		clickedImg:  asset.Image("/asset/tiles/clicked.png"),
		clickedNode: nil,
	}

	return user
}

func (u *User) Render(screen *ebiten.Image) {
	if u.clickedNode == nil {
		return
	}

	pos := u.clickedNode.Position()
	x, _ := u.clickedImg.Size()
	radiusHalf := float64(x) * 0.5
	pos.X -= radiusHalf
	pos.Y -= radiusHalf

	renderer.Render(screen, u.clickedImg, pos)
}

func (u *User) Tick(pos vector.Vector) {
	if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		return
	}

	u.clickedNode = u.world.NearestNode(pos)
}
