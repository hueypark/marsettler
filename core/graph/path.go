package graph

type Path struct {
	ids []int64
}

func (path Path) Empty() bool {
	return len(path.ids) == 0
}

func (path *Path) Push(id int64) {
	path.ids = append(path.ids, id)
}

func (path *Path) Pop() *int64 {
	l := len(path.ids)
	if l < 1 {
		return nil
	}

	var id *int64
	id, path.ids = &path.ids[l-1], path.ids[:l-1]

	return id
}
