package behavior_tree

// State of behavior tree node.
type State int

// Behavior tree States.
const (
	Invalid State = iota
	Success
	Failure
	Running
)