package linked_list

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func setup() (*Node[string], *Node[string], *Node[string]) {
	var node = &Node[string]{content: "hello"}
	var node2 = &Node[string]{content: "heini"}
	var node3 = &Node[string]{content: "!"}

	return node, node2, node3
}

func TestNewLinkedList(t *testing.T) {
	t.Run("NewLinked list with a origin node", func(t *testing.T) {
		node, _, _ := setup()
		lList := NewLinkedList(node)

		require.NotEmpty(t, lList)
		require.Equal(t, node, lList.Origin)
	})

	t.Run("NewLinked list with no origin node", func(t *testing.T) {

		lList := NewLinkedList[string]()

		require.Empty(t, lList.Origin)
	})
}

func TestLinkedList_Append(t *testing.T) {
	t.Run("append a node to linked list", func(t *testing.T) {
		node, node2, _ := setup()
		lList := NewLinkedList(node)
		lList.Append(node2)

		require.Equal(t, node2, lList.Origin.Next)
	})

	t.Run("append to an empty list", func(t *testing.T) {
		node, _, _ := setup()
		list := NewLinkedList[string]()
		list.Append(node)

		require.Equal(t, node, list.Origin)
	})
}

func TestLinkedList_LastNode(t *testing.T) {

	t.Run("returns nil if no node exists", func(t *testing.T) {
		list := NewLinkedList[string]()

		require.Nil(t, list.LastNode())
	})

	t.Run("returns last node", func(t *testing.T) {
		node, node2, _ := setup()
		lList := NewLinkedList(node)
		lList.Append(node2)

		require.Equal(t, node2, lList.LastNode())
	})

	t.Run("last node should be origin if no other node exists", func(t *testing.T) {
		node, _, _ := setup()
		lList := NewLinkedList(node)

		require.Equal(t, node, lList.LastNode())
	})
}

func TestLinkedList_Remove(t *testing.T) {
	t.Run("remove last node from linked list", func(t *testing.T) {
		node, node2, _ := setup()
		lList := NewLinkedList(node)
		lList.Append(node2)
		lList.Remove(node2)

		require.Equal(t, node, lList.LastNode())
	})

	t.Run("remove not present node", func(t *testing.T) {
		node, node2, _ := setup()
		lList := NewLinkedList(node)
		lList.Remove(node2)

		require.Equal(t, node, lList.Origin)
		require.Equal(t, node, lList.LastNode())

	})

	t.Run("remove a in between node", func(t *testing.T) {
		node, node2, node3 := setup()
		lList := NewLinkedList(node, node2, node3)
		lList.Remove(node2)

		require.Equal(t, node3, lList.Origin.Next)
	})

	t.Run("remove origin node", func(t *testing.T) {
		node, node2, _ := setup()
		lList := NewLinkedList(node, node2)
		lList.Remove(node)

		require.Equal(t, node2, lList.Origin)
		require.Equal(t, node2, lList.LastNode())
	})
}

func TestLinkedList_Len(t *testing.T) {

	t.Run("len of empty should be 0", func(t *testing.T) {
		lList := NewLinkedList[string]()

		require.Equal(t, 0, lList.Len())
	})

	t.Run("len should be 1", func(t *testing.T) {
		node, _, _ := setup()
		lList := NewLinkedList[string](node)

		require.Equal(t, 1, lList.Len())
	})

	t.Run("len should be 2", func(t *testing.T) {
		node, node2, _ := setup()
		lList := NewLinkedList[string](node, node2)

		require.Equal(t, 2, lList.Len())
	})

	t.Run("len should still be 2", func(t *testing.T) {
		node, node2, node3 := setup()
		lList := NewLinkedList[string](node, node2, node3)
		lList.Remove(node3)

		require.Equal(t, 2, lList.Len())
	})

}

func TestLinkedList_String(t *testing.T) {
	t.Run("should print content of nodes", func(t *testing.T) {
		node, node2, node3 := setup()
		lList := NewLinkedList[string](node, node2, node3)

		expected := "hello\nheini\n!\n"
		require.Equal(t, expected, lList.String())
	})

	t.Run("should use custom stringer", func(t *testing.T) {
		node, node2, node3 := setup()
		list := NewLinkedList[string](node, node2, node3)
		list.Stringer = func(val string) string {
			return val + " "
		}

		expected := "hello heini ! "
		require.Equal(t, expected, list.String())
	})

}

func TestLinkedList_CustomType(t *testing.T) {

	type WestCostCustoms struct {
		West    string
		Coast   string
		Customs int
	}

	t.Run("should work with custom Type", func(t *testing.T) {
		node := &Node[WestCostCustoms]{
			content: WestCostCustoms{
				West:    "High",
				Coast:   "way",
				Customs: 1,
			},
		}
		list := NewLinkedList(node)

		require.NotEmpty(t, list)
		require.NotEmpty(t, list.Origin)
		require.Equal(t, node, list.Origin)
		require.Equal(t, node, list.LastNode())
	})

	t.Run("should string with custom stringer for custom type", func(t *testing.T) {
		node := &Node[WestCostCustoms]{
			content: WestCostCustoms{
				West:    "High",
				Coast:   "way",
				Customs: 1,
			},
		}
		node2 := &Node[WestCostCustoms]{
			content: WestCostCustoms{
				West:    "Rou",
				Coast:   "te",
				Customs: 66,
			},
		}
		list := NewLinkedList(node, node2)
		list.Stringer = func(val WestCostCustoms) string {
			return fmt.Sprintf("%s%s %d\n", val.West, val.Coast, val.Customs)
		}

		expectedString := "Highway 1\nRoute 66\n"
		require.Equal(t, expectedString, list.String())
	})
}
