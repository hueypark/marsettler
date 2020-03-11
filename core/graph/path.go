package graph

// Path represents a path between nodes.
type Path struct {
	ids []int64
}

// Clear clears path.
func (path *Path) Clear() {
	path.ids = nil
}

// Empty returns true if the path is empty.
func (path *Path) Empty() bool {
	return len(path.ids) == 0
}

// Push pushes the node ID to the path
func (path *Path) Push(id int64) {
	path.ids = append(path.ids, id)
}

// Pop pops the node ID from the path
func (path *Path) Pop() *int64 {
	l := len(path.ids)
	if l < 1 {
		return nil
	}

	var id *int64
	id, path.ids = &path.ids[l-1], path.ids[:l-1]

	return id
}
