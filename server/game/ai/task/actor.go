package task

type Actor interface {
	CreateActor(id int64)
	FindPath() *[]int64
	Move(nodeID int64)
}
