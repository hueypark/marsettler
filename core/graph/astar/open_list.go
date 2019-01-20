package astar

import "sort"

type OpenList struct {
	scores   Scores
	scoreMap map[int64]*Score
}

func NewOpenList() *OpenList {
	return &OpenList{
		Scores{},
		make(map[int64]*Score)}
}

func (openList OpenList) Len() int {
	return openList.scores.Len()
}

func (openList OpenList) Empty() bool {
	return openList.Len() == 0
}

func (openList *OpenList) Push(ps *Score) {
	openList.scores = append(openList.scores, ps)
	sort.Sort(openList.scores)

	openList.scoreMap[ps.ID] = ps
}

func (openList *OpenList) Pop() *Score {
	old := openList.scores
	n := len(old)
	ps := old[n-1]
	openList.scores = old[0 : n-1]

	delete(openList.scoreMap, ps.ID)

	return ps
}

func (openList OpenList) Get(id int64) *Score {
	return openList.scoreMap[id]
}

type Scores []*Score

func (scores Scores) Len() int {
	return len(scores)
}

func (scores Scores) Less(i, j int) bool {
	return scores[i].F > scores[j].F
}

func (scores Scores) Swap(i, j int) {
	scores[i], scores[j] = scores[j], scores[i]
}
