package types

import (
	"errors"
	. "ladderGame/pkg/helpers"
	"ladderGame/pkg/queue"
)

type GraphNode struct {
	Value    int
	Children []*GraphNode
}

func (n *GraphNode) AddChild(child *GraphNode) {
	if !n.HasChild(child) {
		n.Children = append(n.Children, child)
	}
}

func (n *GraphNode) RemoveChild(child *GraphNode) {
	if n.HasChild(child) {
		childIndex := IndexOf(n.Children, child)
		n.Children = append(n.Children[:childIndex], n.Children[childIndex+1:]...)
	}
}

func (n *GraphNode) HasChild(child *GraphNode) bool {
	for _, nChild := range n.Children {
		if nChild == child {
			return true
		}
	}
	return false
}

type Graph struct {
	Nodes []*GraphNode
}

func NewGraph(start, end int, ladders, snakes map[int]int) (graph *Graph, err error) {
	graph = &Graph{}
	for num := start; num <= end; num++ {
		graph.Nodes = append(graph.Nodes, &GraphNode{Value: num, Children: []*GraphNode{}})
	}

	for n1 := start; n1 < end; n1++ {

		if _, ok := ladders[n1]; ok {
			continue
		} else if _, ok := snakes[n1]; ok {
			continue
		}

		for dice := 1; dice <= 6; dice++ {
			n2 := n1 + dice

			if n2 > end {
				continue
			}

			if destLadder, ok := ladders[n2]; ok {
				n2 = destLadder
			} else if destSnake, ok := snakes[n2]; ok {
				n2 = destSnake
			}

			err = graph.AddEdge(graph.Nodes[n1], graph.Nodes[n2])
			if err != nil {
				return
			}

		}
	}

	// should this return pointer to graph??
	return
}

func (g *Graph) AddEdge(node1, node2 *GraphNode) error {
	if IndexOf(g.Nodes, node1) != -1 && IndexOf(g.Nodes, node2) != -1 {
		node1.AddChild(node2)
	} else {
		return errors.New("nodes do not exist on graph")
	}
	return nil
}

func (g *Graph) RemoveEdge(node1, node2 *GraphNode) error {
	if IndexOf(g.Nodes, node1) != -1 && IndexOf(g.Nodes, node2) != -1 {
		node1.RemoveChild(node2)
	} else {
		return errors.New("nodes do not exist on graph")
	}

	return nil
}

func (g *Graph) ShortestPathLadderGame() ([]*GraphNode, []int) {
	targetNode := g.Nodes[len(g.Nodes)-1]
	q := queue.Queue{}
	q.Enq(g.Nodes[0])
	var visited []*GraphNode
	ways := map[int][]*GraphNode{
		0: []*GraphNode{},
	}

	for !q.IsEmpty() {
		deq, ok := q.Deq()
		if !ok {
			continue
		}
		currentNode := deq.(*GraphNode)
		visited = append(visited, currentNode)

		for _, adjacent := range currentNode.Children {
			adjacentWay, adjacentHasWay := ways[adjacent.Value]
			if len(adjacentWay) > len(ways[currentNode.Value])+1 || !adjacentHasWay {
				ways[adjacent.Value] = append(ways[currentNode.Value], currentNode)
			}
			if IndexOf(visited, adjacent) == -1 {
				q.Enq(adjacent)
			}
		}

	}
	// add end node itself to its way
	result := append(ways[targetNode.Value], targetNode)
	return result, nil
}
