package task

type Actor interface {
	CreateActor(id int)
	FindPath() *[]int64
	Move(nodeID int64)
}
