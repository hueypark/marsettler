package ai

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hueypark/marsettler/core/behavior_tree"
	"github.com/hueypark/marsettler/pkg/consts"
)

type Wait struct {
	behavior_tree.Node

	duration        int
	currentDuration int
}

func NewWait(params string) *Wait {
	duration, err := strconv.ParseInt(strings.ReplaceAll(params, " ", ""), 10, 64)
	if err != nil {
		log.Println(err)
	}

	return &Wait{duration: int(duration)}
}

func (node *Wait) Init() {
	node.currentDuration = 0
}

func (node *Wait) Tick() behavior_tree.State {
	node.currentDuration += consts.Delta
	if node.duration <= node.currentDuration {
		return node.SetState(behavior_tree.Success)
	}

	return node.SetState(behavior_tree.Running)
}

func (node *Wait) Wireframe() string {
	return fmt.Sprintf("Wait: %v", node.duration)
}
