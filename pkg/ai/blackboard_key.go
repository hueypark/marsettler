package ai

import "github.com/hueypark/marsettler/core/behavior_tree"

const (
	patrolPosition behavior_tree.BlackboardKey = 0
)

func Key(key string) behavior_tree.BlackboardKey {
	switch key {
	case "patrolPosition":
		return patrolPosition
	default:
		return -1
	}
}

func String(key behavior_tree.BlackboardKey) string {
	switch key {
	case patrolPosition:
		return "patrolPosition"
	default:
		return ""
	}
}
