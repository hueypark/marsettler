package ai

import (
	"fmt"
	"strings"

	"github.com/hueypark/marsettler/core/behavior_tree"
)

type nodeData struct {
	str      string
	children []nodeData
}

func (n *nodeData) depth() int {
	return strings.Count(n.str, "\t")
}

func NewBehaviorTree(actor actor, str string) (*behavior_tree.BehaviorTree, error) {
	bt := behavior_tree.NewBehaviorTree()

	nodeData, err := parse(str)
	if err != nil {
		return nil, err
	}

	err = addNode(actor, bt.Root(), nodeData, bt.Blackboard())
	if err != nil {
		return bt, err
	}

	return bt, nil
}

func addNode(actor actor, parent behavior_tree.INode, nodeData nodeData, blackboard *behavior_tree.Blackboard) error {
	newNode, err := newNode(actor, nodeData.str, blackboard)
	if err != nil {
		return err
	}

	addChildNode, ok := parent.(behavior_tree.IAddChildNode)
	if !ok {
		return fmt.Errorf("parent is not add child node [parent: %v]", parent)
	}

	addChildNode.AddChild(newNode)

	for _, child := range nodeData.children {
		err := addNode(actor, newNode, child, blackboard)
		if err != nil {
			return err
		}
	}

	return nil
}

func newNode(actor actor, str string, blackboard *behavior_tree.Blackboard) (behavior_tree.INode, error) {
	str = strings.ReplaceAll(str, "\t", "")
	strs := strings.SplitN(str, ":", 2)

	name := strs[0]
	var params string
	if 1 < len(strs) {
		params = strs[1]
	}

	switch name {
	case "Sequence":
		return behavior_tree.NewSequence(), nil
	case "FindRandomPosition":
		return NewFindRandomPosition(actor, blackboard, params), nil
	case "MoveTo":
		return NewMoveTo(actor, blackboard, params), nil
	case "Wait":
		return NewWait(params), nil
	default:
		return nil, fmt.Errorf("no node named %v", name)
	}
}

func parse(str string) (nodeData, error) {
	nodeStrs := strings.Split(str, "\n")

	var root *nodeData
	var parents []*nodeData
	for _, nodeStr := range nodeStrs {
		if nodeStr == "" {
			continue
		}

		n := nodeData{
			str:      nodeStr,
			children: nil,
		}

		for {
			var parent *nodeData
			if 1 <= len(parents) {
				parent = parents[len(parents)-1]
			} else {
				root = &n
				break
			}

			diffDepth := parent.depth() - n.depth()
			if 0 < diffDepth {
				parents = parents[:len(parents)-1]
			} else if diffDepth == -1 {
				parent.children = append(parent.children, n)
				break
			} else {
				return nodeData{}, fmt.Errorf("depth is invalid. [str: %v, nodeData: %v]", str, n)
			}
		}

		parents = append(parents, root)
	}

	if root == nil {
		return nodeData{}, fmt.Errorf("nodeData is empty. [str: %v]", str)
	}

	return *root, nil
}
