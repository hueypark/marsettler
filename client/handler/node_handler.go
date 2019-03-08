package handler

import (
	"github.com/hueypark/marsettler/client/game"
	"github.com/hueypark/marsettler/core/math/vector"
	"github.com/hueypark/marsettler/server/game/message/fbs"
)

func handleNode(node *fbs.Node) {
	position := node.Position(nil)

	var edgeMessage fbs.Edge
	var edges []*game.Edge
	edgesLength := node.EdgesLength()
	for i := 0; i < edgesLength; i++ {
		node.Edges(&edgeMessage, i)

		edge := game.GetEdge(0)
		if edge == nil {
			edge = game.NewEdge(0, edgeMessage.StartNodeID(), edgeMessage.EndNodeID())
		}
		edges = append(edges, edge)
	}

	game.NewNode(node.ID(), vector.Vector{X: position.X(), Y: position.Y()}, edges)
}
