package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/pkg/renderer"
)

type User struct {
	actor       *Actor
	world       *World
	clickedImg  *ebiten.Image
	clickedNode *Node
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

	node := u.world.NearestNode(pos)
	if node == nil {
		return
	}

	u.clickedNode = u.world.NearestNode(pos)
	u.actor.behaviorTree.Blackboard().SetInt64(behavior_tree.Key("node"), u.clickedNode.ID())
}
