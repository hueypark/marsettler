package graph

import (
	"errors"

	"github.com/hueypark/marsettler/core/graph/astar"
	"github.com/hueypark/marsettler/core/id_generator"
)

// Graph represents graph.
type Graph struct {
	nodes map[int64]Node
	edges map[int64]map[int64]Edge
}

// NewGraph create new graph.
func NewGraph() *Graph {
	return &Graph{
		make(map[int64]Node),
		make(map[int64]map[int64]Edge),
	}
}

// AddNode adds node.
func (graph *Graph) AddNode(node Node) {
	graph.nodes[node.ID()] = node
}

// Nodes returns nodes.
func (graph Graph) Nodes() map[int64]Node {
	return graph.nodes
}

// AddEdge adds edge.
func (graph *Graph) AddEdge(from, to int64) {
	if _, ok := graph.edges[from]; !ok {
		graph.edges[from] = make(map[int64]Edge)
	}

	graph.edges[from][to] = Edge{id_generator.Generate(), from, to}
}

// Edges returns edges.
func (graph Graph) Edges() (edges []Edge) {
	for _, val := range graph.edges {
		for _, edge := range val {
			edges = append(edges, edge)
		}
	}

	return edges
}

// Path returns path between nodes.
func (graph Graph) Path(fromNodeID, toNodeID int64) (path Path, err error) {
	openList := astar.NewOpenList()
	closedList := astar.NewClosedList()

	openList.Push(&astar.Score{
		ID:       fromNodeID,
		F:        0,
		G:        0,
		H:        0,
		ParentID: 0})

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
				ID:       neighbor.ID(),
				F:        g + h,
				G:        g,
				H:        h,
				ParentID: openScore.ID})
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
