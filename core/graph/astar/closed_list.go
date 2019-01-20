package astar

type ClosedList struct {
	scoreMap map[int64]*Score
}

func NewClosedList() *ClosedList {
	return &ClosedList{
		make(map[int64]*Score)}
}

func (cl ClosedList) Get(id int64) (score *Score, ok bool) {
	score, ok = cl.scoreMap[id]
	return score, ok
}

func (cl *ClosedList) Set(score *Score) {
	cl.scoreMap[score.ID] = score
}
