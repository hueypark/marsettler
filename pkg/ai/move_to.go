package ai

import (
	"fmt"
	"strings"

	"github.com/hueypark/marsettler/core/behavior_tree"
)

type MoveTo struct {
	behavior_tree.Node

	positionKey behavior_tree.BlackboardKey
}

func NewMoveTo(params string) *MoveTo {
	positionKey := strings.ReplaceAll(params, " ", "")

	return &MoveTo{positionKey: Key(positionKey)}
}

func (node *MoveTo) Wireframe() string {
	return fmt.Sprintf("MoveTo: %v", String(node.positionKey))
}
