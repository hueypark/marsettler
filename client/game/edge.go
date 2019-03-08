package game

// Edge represents edge between nodes.
type Edge struct {
	From int64
	To   int64
}

// NewEdge create new edge.
func NewEdge(id int64, from, to int64) *Edge {
	edge := &Edge{
		from,
		to,
	}

	return edge
}

// GetEdge returns edge.
func GetEdge(id int64) *Edge {
	return nil
}
