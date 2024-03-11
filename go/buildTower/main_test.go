package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCombination struct {
	Input           int
	ExpectedOutcome []string
}

var testArray = []TestCombination{
	TestCombination{
		Input:           0,
		ExpectedOutcome: []string{},
	},
	TestCombination{
		Input:           1,
		ExpectedOutcome: []string{"*"},
	},
	TestCombination{
		Input:           2,
		ExpectedOutcome: []string{" * ", "***"},
	},
	TestCombination{
		Input:           3,
		ExpectedOutcome: []string{"  *  ", " *** ", "*****"},
	},
	TestCombination{
		Input:           4,
		ExpectedOutcome: []string{"   *   ", "  ***  ", " ***** ", "*******"},
	},
}

func TestTowerBuilder(t *testing.T) {
	for _, testComb := range testArray {
		t.Run("testing input of "+fmt.Sprintf("%d", testComb.Input), func(t *testing.T) {
			assert.Equal(t, testComb.ExpectedOutcome, TowerBuilder(testComb.Input))
		})

	}

}
