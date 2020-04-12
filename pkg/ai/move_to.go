package ai

import (
	"fmt"
	"log"
	"strings"

	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/core/graph"
	"github.com/hueypark/marsettler/pkg/game"
)

type moveToState int

const (
	findPath moveToState = iota
	moveTo
)

type MoveTo struct {
	behavior_tree.Node

	state moveToState

	actor      *game.Actor
	blackboard *behavior_tree.Blackboard
	nodeIDKey  behavior_tree.BlackboardKey
	path       graph.Path
}

func NewMoveTo(actor *game.Actor, blackboard *behavior_tree.Blackboard, params string) *MoveTo {
	nodeIDKey := strings.ReplaceAll(params, " ", "")

	return &MoveTo{
		state:      findPath,
		actor:      actor,
		blackboard: blackboard,
		nodeIDKey:  behavior_tree.Key(nodeIDKey),
	}
}

func (n *MoveTo) Init() {
	n.state = findPath
	n.path.Clear()
}

func (n *MoveTo) Tick() behavior_tree.State {
	switch n.state {
	case findPath:
		nodeID := n.blackboard.GetInt64(n.nodeIDKey)
		if nodeID == nil {
			return n.SetState(behavior_tree.Failure)
		}

		if n.actor.NodeID() == *nodeID {
			return n.SetState(behavior_tree.Success)
		}

		path, err := n.actor.FindPath(*nodeID)
		if err != nil {
			log.Println(err)
			return n.SetState(behavior_tree.Failure)
		}

		n.path = path
		n.state = moveTo
	case moveTo:
		if !n.actor.CanMove() {
			break
		}

		nextNodeID := n.path.Pop()
		if nextNodeID == nil {
			log.Println("Next node id is nil")
			return n.SetState(behavior_tree.Failure)
		}

		if *nextNodeID == n.actor.NodeID() {
			nextNodeID = n.path.Pop()
		}
		if nextNodeID == nil {
			log.Println("Next node id is nil")
			return n.SetState(behavior_tree.Failure)
		}

		err := n.actor.Move(*nextNodeID)
		if err != nil {
			return n.SetState(behavior_tree.Failure)
		}

		if n.path.Empty() {
			return n.SetState(behavior_tree.Success)
		}
	}

	return n.SetState(behavior_tree.Running)
}

func (n *MoveTo) Wireframe() string {
	return fmt.Sprintf("MoveTo: %v", behavior_tree.StringKey(n.nodeIDKey))
}
