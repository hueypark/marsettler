package ai

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hueypark/marsettler/core/behavior_tree"
)

type Wait struct {
	behavior_tree.Node

	duration        float64
	currentDuration float64
}

func NewWait(params string) *Wait {
	duration, err := strconv.ParseFloat(strings.ReplaceAll(params, " ", ""), 64)
	if err != nil {
		log.Println(err)
	}

	return &Wait{duration: duration}
}

func (node *Wait) Init() {
	node.currentDuration = 0
}

func (node *Wait) Tick(delta float64) behavior_tree.State {
	node.currentDuration += delta
	if node.duration <= node.currentDuration {
		return node.SetState(behavior_tree.Success)
	}

	return node.SetState(behavior_tree.Running)
}

func (node *Wait) Wireframe() string {
	return fmt.Sprintf("Wait: %v", node.duration)
}
