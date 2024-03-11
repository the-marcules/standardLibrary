package linked_list

import "fmt"

type Node[T comparable] struct {
	Next    *Node[T]
	content T
}

type LinkedList[T comparable] struct {
	Origin   *Node[T]
	Stringer func(val T) string
}

func NewLinkedList[T comparable](nodes ...*Node[T]) *LinkedList[T] {
	list := LinkedList[T]{
		Stringer: func(val T) string {
			return fmt.Sprintf("%v\n", val)
		},
	}

	for _, node := range nodes {
		list.Append(node)
	}

	return &list
}

func (l *LinkedList[T]) Append(node *Node[T]) {
	lastNode := l.LastNode()
	if lastNode == nil {
		l.Origin = node
		return
	}
	lastNode.Next = node
}

func (l *LinkedList[T]) LastNode() *Node[T] {
	if l.Origin == nil {
		return nil
	}
	currentNode := l.Origin
	nextNode := currentNode.Next
	for nextNode != nil {
		currentNode = nextNode
		nextNode = currentNode.Next
	}

	return currentNode
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if l.Origin == node {
		l.Origin = l.Origin.Next
		return
	}

	currentNode := l.Origin

	for currentNode.Next != node && currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	if currentNode == nil || currentNode.Next == nil {
		return
	}

	currentNode.Next = currentNode.Next.Next
}

func (l *LinkedList[T]) Len() int {
	count := 0
	if l.Origin == nil {
		return count
	}
	currentNode := l.Origin
	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}
	return count
}

func (l *LinkedList[T]) String() string {
	currentNode := l.Origin
	output := ""
	for currentNode != nil {
		output += l.Stringer(currentNode.content)
		currentNode = currentNode.Next
	}
	return output
}
