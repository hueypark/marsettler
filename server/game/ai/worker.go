package ai

import "github.com/hueypark/marsettler/core/behavior_tree"

// NewWorker creates new worker.
func NewWorker() *behavior_tree.BehaviorTree {
	worker := &behavior_tree.BehaviorTree{}
	worker.SetRoot(&behavior_tree.Node{})

	return worker
}
