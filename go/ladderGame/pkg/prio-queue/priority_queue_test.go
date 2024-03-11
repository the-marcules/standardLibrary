package prio_queue

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ladderGame/pkg/types"
	"sort"
	"testing"
)

func TestPrioQ(t *testing.T) {
	testCases := [][]int{
		{4, 3, 5, 6, 1, 2},
		{1, 1, 5, 6, 1, 2, 6, 6, 10},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	for i, testcase := range testCases {
		t.Run("test case "+fmt.Sprint(i), func(t *testing.T) {
			prio := PrioQ{}
			for _, num := range testcase {
				prio.Insert(types.Field{num, []int{}, num})
			}
			sort.Ints(testcase)

			var got []int

			for prio.HasItems() == true {
				tmpField := prio.Pop()
				got = append(got, tmpField.FieldNum)
			}

			assert.Equal(t, testcase, got)
		})
	}

	t.Run("multiple insert and pop operations", func(t *testing.T) {
		prio := PrioQ{}
		for _, num := range []int{1, 5, 10} {
			prio.Insert(types.Field{num, []int{}, num})
		}
		assert.Equal(t, 1, prio.Pop().FieldNum)
		prio.Insert(types.Field{2, []int{}, 2})
		prio.Insert(types.Field{16, []int{}, 16})
		assert.Equal(t, 2, prio.Pop().FieldNum)
		assert.Equal(t, 5, prio.Pop().FieldNum)
		assert.Equal(t, 10, prio.Pop().FieldNum)
		prio.Insert(types.Field{13, []int{}, 13})
		assert.Equal(t, 13, prio.Pop().FieldNum)
		assert.Equal(t, 16, prio.Pop().FieldNum)
		assert.Equal(t, false, prio.HasItems())
	})

}
