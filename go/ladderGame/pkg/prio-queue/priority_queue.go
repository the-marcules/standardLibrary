package prio_queue

import (
	"ladderGame/pkg/types"
	"math"
)

type PrioQ struct {
	queue     []types.Field
	nextIndex int
}

func (q *PrioQ) Insert(field types.Field) {
	q.queue = append(q.queue, field)
	q.HeapifyUp(q.nextIndex)
	q.nextIndex++
}

func (q *PrioQ) Pop() (min types.Field) {
	min = q.queue[0]
	q.queue = q.queue[1:]
	q.nextIndex--
	q.HeapifyDown(0)
	return
}

func (q *PrioQ) HasItems() bool {
	return q.nextIndex > 0
}

func LeftChildIndex(index int) int {
	return 2*index + 1
}
func RightChildIndex(index int) int {
	return 2*index + 2
}

func ParentIndex(index int) int {
	return int(math.Floor(float64(index) / 2.0))
}

func (q *PrioQ) IsValidItem(index int) bool {
	if index >= q.nextIndex {
		return false
	}
	return true
}

func (q *PrioQ) Swap(a, b int) {
	q.queue[a], q.queue[b] = q.queue[b], q.queue[a]
}

func (q *PrioQ) HeapifyUp(index int) {
	parent := ParentIndex(index)
	if q.queue[parent].RemainingDistance > q.queue[index].RemainingDistance {
		q.Swap(index, parent)
		q.HeapifyUp(parent)
	}
}

func (q *PrioQ) HeapifyDown(index int) {
	minIndex := index
	leftChildIndex := LeftChildIndex(index)
	rightChildIndex := RightChildIndex(index)

	if q.IsValidItem(leftChildIndex) && q.queue[leftChildIndex].RemainingDistance < q.queue[minIndex].RemainingDistance {
		minIndex = leftChildIndex
	}

	if q.IsValidItem(rightChildIndex) && q.queue[rightChildIndex].RemainingDistance < q.queue[minIndex].RemainingDistance {
		minIndex = rightChildIndex
	}
	if minIndex != index {
		q.Swap(minIndex, index)
		q.HeapifyDown(minIndex)
	}

}
