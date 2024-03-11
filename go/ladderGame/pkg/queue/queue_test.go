package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Items come back in expexted order", func(t *testing.T) {
		input := []int{1, 2, 3}
		var qOutput []int
		var q Queue

		for _, value := range input {
			q.Enq(value)
		}

		for q.IsEmpty() == false {
			val, ok := q.Deq()
			if !ok {
				t.Fatal("Queue is unexpectedly empty")
			}
			qOutput = append(qOutput, val.(int))
		}
		assert.Equal(t, input, qOutput)

	})
}
