package helpers

import (
	"fmt"
	"testing"
)

type TestCase struct {
	stack  []string
	needle string
	want   int
}

func TestIndexOf(t *testing.T) {

	testCases := []TestCase{
		{
			stack:  []string{"a"},
			needle: "b",
			want:   -1,
		},
		{
			stack:  []string{"a", "b", "c", "b"},
			needle: "b",
			want:   1,
		}, {
			stack:  []string{"a", "b", "c", "b"},
			needle: "c",
			want:   2,
		},
		{
			stack:  []string{},
			needle: "b",
			want:   -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Find first occurrance of %s", testCase.needle), func(t *testing.T) {
			got := IndexOf(testCase.stack, testCase.needle)

			if got != testCase.want {
				t.Errorf("wanted %d, got %d", testCase.want, got)
			}
		})
	}
}
