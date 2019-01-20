package astar

type Score struct {
	ID       int64
	F        float64
	G        float64
	H        float64
	ParentID int64
}
