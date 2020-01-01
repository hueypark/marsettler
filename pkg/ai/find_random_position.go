package ai

import (
	"fmt"
	"strings"

	"github.com/hueypark/marsettler/core/behavior_tree"
)

type FindRandomPosition struct {
	behavior_tree.Node

	positionKey behavior_tree.BlackboardKey
}

func NewFindRandomPosition(params string) *FindRandomPosition {
	positionKey := strings.ReplaceAll(params, " ", "")

	return &FindRandomPosition{positionKey: Key(positionKey)}
}

func (node *FindRandomPosition) Wireframe() string {
	return fmt.Sprintf("FindRandomPosition: %v", String(node.positionKey))
}
