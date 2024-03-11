package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	. "ladderGame/pkg/types"
	"testing"
)

type TestCase struct {
	gameField      GameField
	expectedResult []int
}

func TestName(t *testing.T) {
	testCases := []TestCase{
		{
			gameField: GameField{
				Ladders: map[int]int{
					1:  38,
					4:  14,
					9:  31,
					21: 42,
					28: 84,
					36: 44,
					51: 67,
					71: 91,
					80: 100},
				Snakes: map[int]int{
					16: 6,
					47: 26,
					49: 11,
					56: 53,
					62: 19,
					64: 60,
					87: 24,
					93: 73,
					95: 75,
					98: 78},
				Dice:  []int{6, 5, 4, 3, 2, 1},
				Start: 0,
				End:   100,
			},
			expectedResult: []int{1, 6, 6, 1, 4, 6, 3},
		},
		{
			gameField: GameField{
				Ladders: map[int]int{
					3:  51,
					6:  27,
					20: 70,
					36: 55,
					63: 95,
					68: 98,
				},
				Snakes: map[int]int{
					25: 5,
					34: 1,
					47: 19,
					65: 52,
					87: 57,
					91: 61,
					99: 69,
				},
				Dice:  []int{6, 5, 4, 3, 2, 1},
				Start: 0,
				End:   100,
			},
			expectedResult: []int{3, 6, 6, 5},
		},
		{
			gameField: GameField{
				Ladders: map[int]int{
					4:  25,
					13: 46,
					33: 49,
					42: 63,
					50: 69,
					62: 81,
					74: 92,
				},
				Snakes: map[int]int{
					27: 5,
					40: 3,
					43: 18,
					54: 31,
					66: 45,
					76: 58,
					89: 53,
					99: 41,
				},
				Dice:  []int{6, 5, 4, 3, 2, 1},
				Start: 0,
				End:   100,
			},
			expectedResult: []int{4, 6, 2, 1, 5, 6, 2},
		},
		{
			gameField: GameField{
				Ladders: map[int]int{},
				Snakes:  map[int]int{},
				Dice:    []int{6, 5, 4, 3, 2, 1},
				Start:   0,
				End:     100,
			},
			expectedResult: []int{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 4},
		}}

	for i, testcase := range testCases {
		t.Run("test case "+fmt.Sprint(i+1), func(t *testing.T) {
			got := aStarLadderGame(testcase.gameField)
			assert.Equal(t, testcase.expectedResult, got)
			fmt.Printf("Shortest dice compination is : %v \n", got)
		})
	}
}
