package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGraphNode_AddChild(t *testing.T) {
	parentNode := GraphNode{
		Value:    1,
		Children: []*GraphNode{},
	}

	childNode := GraphNode{
		Value:    2,
		Children: []*GraphNode{},
	}

	parentNode.AddChild(&childNode)

	if len(parentNode.Children) != 1 {
		t.Error("Parent has no children at all.")
	}

	if parentNode.Children[0].Value != childNode.Value {
		t.Error("Child node differs")
	}

	if !parentNode.HasChild(&childNode) {
		t.Error("parent has not the childNode it should have.")
	}
}

func TestGraphNode_RemoveChild(t *testing.T) {
	childNode := GraphNode{
		Value:    0,
		Children: []*GraphNode{},
	}
	childNode1 := GraphNode{
		Value:    1,
		Children: []*GraphNode{},
	}
	childNode2 := GraphNode{
		Value:    2,
		Children: []*GraphNode{},
	}
	childNode3 := GraphNode{
		Value:    3,
		Children: []*GraphNode{},
	}

	t.Run("remove child from start", func(t *testing.T) {
		parentNode := GraphNode{
			Value:    99,
			Children: []*GraphNode{},
		}

		parentNode.AddChild(&childNode)
		parentNode.AddChild(&childNode1)

		if !parentNode.HasChild(&childNode) && !parentNode.HasChild(&childNode1) {
			t.Error("parent has not the desired childNodes")
		}

		parentNode.RemoveChild(&childNode)

		if parentNode.HasChild(&childNode) {
			t.Error("parent still has childNode, but it should not")
		}

		if !parentNode.HasChild(&childNode1) {
			t.Error("there should be one remainding child node")
		}
	})

	t.Run("remove child from middle", func(t *testing.T) {
		parentNode := GraphNode{
			Value:    99,
			Children: []*GraphNode{},
		}

		parentNode.AddChild(&childNode)
		parentNode.AddChild(&childNode1)
		parentNode.AddChild(&childNode2)
		parentNode.AddChild(&childNode3)

		if !parentNode.HasChild(&childNode) && !parentNode.HasChild(&childNode1) && !parentNode.HasChild(&childNode2) && !parentNode.HasChild(&childNode3) {
			t.Error("parent has not the desired childNodes")
		}

		parentNode.RemoveChild(&childNode2)

		if parentNode.HasChild(&childNode2) {
			t.Error("parent still has childNode, but it should not")
		}

		if !parentNode.HasChild(&childNode) && !parentNode.HasChild(&childNode1) && !parentNode.HasChild(&childNode3) {
			t.Error("there should be one remaining child nodes")
		}

		if parentNode.Children[2] != &childNode3 && parentNode.Children[1] != &childNode1 {
			t.Error("following child nodes at unexpected position")
		}
	})

	t.Run("remove child from end", func(t *testing.T) {
		parentNode := GraphNode{
			Value:    99,
			Children: []*GraphNode{},
		}

		parentNode.AddChild(&childNode)
		parentNode.AddChild(&childNode1)
		parentNode.AddChild(&childNode2)
		parentNode.AddChild(&childNode3)

		if !parentNode.HasChild(&childNode) && !parentNode.HasChild(&childNode1) && !parentNode.HasChild(&childNode2) && !parentNode.HasChild(&childNode3) {
			t.Error("parent has not the desired childNodes")
		}

		parentNode.RemoveChild(&childNode3)

		if parentNode.HasChild(&childNode3) {
			t.Error("parent still has childNode, but it should not")
		}

		if !parentNode.HasChild(&childNode) && !parentNode.HasChild(&childNode1) && !parentNode.HasChild(&childNode2) {
			t.Error("there should be one remaining child nodes")
		}

		if len(parentNode.Children) > 3 {
			t.Error("too many child nodes left")
		}

	})
}

func testGraphHelper(t *testing.T) (graph Graph, nodeList []*GraphNode) {
	childNode := GraphNode{
		Value:    0,
		Children: []*GraphNode{},
	}
	childNode1 := GraphNode{
		Value:    1,
		Children: []*GraphNode{},
	}
	childNode2 := GraphNode{
		Value:    2,
		Children: []*GraphNode{},
	}
	childNode3 := GraphNode{
		Value:    3,
		Children: []*GraphNode{},
	}

	nodeList = []*GraphNode{
		&childNode,
		&childNode1,
		&childNode2,
		&childNode3,
	}
	graph.Nodes = append(graph.Nodes, nodeList...)

	return
}

func TestGraph(t *testing.T) {

	t.Run("TestGraph manual init", func(t *testing.T) {
		graph, nodeList := testGraphHelper(t)

		if len(graph.Nodes) < 4 {
			t.Error("no edges found, but should have 4")
		}

		if graph.Nodes[0] != nodeList[0] {
			t.Error("Lists do not match")
		}
	})

	t.Run("NewGraph", func(t *testing.T) {
		start := 0
		end := 20
		ladders := map[int]int{
			2:  13,
			4:  12,
			10: 19,
		}
		snakes := map[int]int{
			14: 9,
			17: 2,
		}
		graph, err := NewGraph(start, end, ladders, snakes)

		if err != nil {
			t.Error(err)
		}

		if len(graph.Nodes) > end+1 {
			t.Error("too many nodes in graph")
		}

		if !graph.Nodes[1].HasChild(graph.Nodes[13]) {
			t.Error("node is missing a specific child")
		}

		if !graph.Nodes[13].HasChild(graph.Nodes[9]) {
			t.Error("node is missing a specific child")
		}

		for index, node := range graph.Nodes {
			if len(node.Children) > 6 {
				t.Errorf("node %d has more than 6 children", index)
			}
		}

		fmt.Printf("Node %d has children %v \n", graph.Nodes[2].Value, graph.Nodes[2].Children)
	})

	t.Run("AddEdge with existing nodes on graph", func(t *testing.T) {
		graph, nodeList := testGraphHelper(t)

		err := graph.AddEdge(nodeList[0], nodeList[1])

		if err != nil {
			t.Error("Got an error but did not expect one", err.Error())
		}

		if !nodeList[0].HasChild(nodeList[1]) {
			t.Error("edge was not created")
		}

	})

	t.Run("AddEdge with non existing nodes on graph", func(t *testing.T) {
		graph, nodeList := testGraphHelper(t)

		newNode := GraphNode{Value: 100, Children: []*GraphNode{}}

		err := graph.AddEdge(nodeList[0], &newNode)

		if err == nil {
			t.Error("Got no error but expect one", err.Error())
		}

		if nodeList[0].HasChild(&newNode) || newNode.HasChild(nodeList[0]) {
			t.Error("edge was created, but it should not")
		}

	})

	t.Run("RemoveEdge", func(t *testing.T) {
		graph, nodeList := testGraphHelper(t)

		graph.AddEdge(nodeList[0], nodeList[1])
		graph.AddEdge(nodeList[0], nodeList[2])

		graph.RemoveEdge(nodeList[0], nodeList[2])

		if nodeList[0].HasChild(nodeList[2]) || nodeList[2].HasChild(nodeList[0]) {
			t.Error("edge was not removed")
		}
	})

}

func TestGraph_ShortestPathLadderGame(t *testing.T) {
	t.Run("Shortest path simple case", func(t *testing.T) {

		start := 0
		end := 12
		ladders := map[int]int{
			2: 11,
		}
		snakes := map[int]int{
			7: 3,
		}
		graph, err := NewGraph(start, end, ladders, snakes)
		expectedTraversedNodes := []*GraphNode{
			graph.Nodes[0],
			graph.Nodes[11],
			graph.Nodes[12],
		}
		//expectedDice := []int{2, 1}

		if err != nil {
			t.Fatal("got error but did not expect one", err.Error())
		}

		traversedNodes, _ := graph.ShortestPathLadderGame()

		if !reflect.DeepEqual(traversedNodes, expectedTraversedNodes) {
			t.Errorf("Got %v but expexted %v", traversedNodes, expectedTraversedNodes)
		}

	})

	t.Run("Shortest path sophisticated case", func(t *testing.T) {

		start := 0
		end := 100
		ladders := map[int]int{
			3:  51,
			6:  27,
			20: 70,
			36: 55,
			63: 95,
			68: 98,
		}
		snakes := map[int]int{
			25: 5,
			34: 1,
			47: 19,
			65: 52,
			87: 57,
			91: 61,
			99: 69,
		}
		graph, err := NewGraph(start, end, ladders, snakes)
		expectedTraversedNodes := []*GraphNode{
			graph.Nodes[0],
			graph.Nodes[51],
			graph.Nodes[57],
			graph.Nodes[95],
			graph.Nodes[100],
		}
		//expectedDice := []int{2, 1}

		if err != nil {
			t.Fatal("got error but did not expect one", err.Error())
		}

		traversedNodes, _ := graph.ShortestPathLadderGame()

		if !reflect.DeepEqual(traversedNodes, expectedTraversedNodes) {
			t.Errorf("Got %v but expexted %v", traversedNodes, expectedTraversedNodes)

			for _, node := range traversedNodes {
				fmt.Printf("%d, ", node.Value)
			}
			fmt.Println()
		}

	})
	t.Run("Shortest path more sophisticated case", func(t *testing.T) {

		start := 0
		end := 100
		ladders := map[int]int{
			1:  38,
			4:  14,
			9:  31,
			21: 42,
			28: 84,
			36: 44,
			51: 67,
			71: 91,
			80: 100,
		}
		snakes := map[int]int{
			16: 6,
			47: 26,
			49: 11,
			56: 53,
			62: 19,
			64: 60,
			87: 24,
			93: 73,
			95: 75,
			98: 78,
		}
		graph, err := NewGraph(start, end, ladders, snakes)
		expectedTraversedNodes := []*GraphNode{
			graph.Nodes[0],
			graph.Nodes[38],
			graph.Nodes[39],
			graph.Nodes[45],
			graph.Nodes[67],
			graph.Nodes[68],
			graph.Nodes[74],
			graph.Nodes[100],
		}
		//expectedDice := []int{2, 1}

		if err != nil {
			t.Fatal("got error but did not expect one", err.Error())
		}

		traversedNodes, _ := graph.ShortestPathLadderGame()

		if !reflect.DeepEqual(traversedNodes, expectedTraversedNodes) {
			t.Errorf("Got %v but expexted %v", traversedNodes, expectedTraversedNodes)

			for _, node := range traversedNodes {
				fmt.Printf("%d, ", node.Value)
			}
			fmt.Println()
		}

	})
}
