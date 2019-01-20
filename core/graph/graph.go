package graph

import (
	"errors"

	"gitlab.com/legionary/legionary/core/graph/astar"
	"gitlab.com/legionary/legionary/core/id_generator"
)

type Graph struct {
	nodes map[int64]Node
	edges map[int64]map[int64]Edge
}

func NewGraph() *Graph {
	return &Graph{
		make(map[int64]Node),
		make(map[int64]map[int64]Edge),
	}
}

func (graph *Graph) AddNode(node Node) {
	graph.nodes[node.ID()] = node
}

func (graph Graph) Nodes() map[int64]Node {
	return graph.nodes
}

func (graph *Graph) AddEdge(from, to int64) {
	if _, ok := graph.edges[from]; !ok {
		graph.edges[from] = make(map[int64]Edge)
	}

	graph.edges[from][to] = Edge{id_generator.Generate(), from, to}
}

func (graph Graph) Edges() (edges []Edge) {
	for _, val := range graph.edges {
		for _, edge := range val {
			edges = append(edges, edge)
		}
	}

	return edges
}

func (graph Graph) Path(fromNodeID, toNodeID int64) (path Path, err error) {
	openList := astar.NewOpenList()
	closedList := astar.NewClosedList()

	openList.Push(&astar.Score{
		fromNodeID,
		0,
		0,
		0,
		0})

	toNode := graph.nodes[toNodeID]

	for !openList.Empty() {
		openScore := openList.Pop()

		if openScore.ID == toNodeID {
			return createPath(closedList, openScore), nil
		}

		closedList.Set(openScore)

		for neighborID := range graph.edges[openScore.ID] {
			if _, ok := closedList.Get(neighborID); ok {
				continue
			}

			neighbor := graph.nodes[neighborID]
			openNode := graph.nodes[openScore.ID]

			g := neighbor.Len(openNode) + openScore.G
			h := toNode.Len(neighbor)

			openList.Push(&astar.Score{
				neighbor.ID(),
				g + h,
				g,
				h,
				openScore.ID})
		}
	}

	return path, errors.New("has no path")
}

func createPath(closedList *astar.ClosedList, score *astar.Score) (path Path) {
	for {
		path.Push(score.ID)

		if score.ParentID == 0 {
			break
		}

		score, _ = closedList.Get(score.ParentID)
	}

	return path
}
