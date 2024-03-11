package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	height          float64
	bounce          float64
	windowHeight    float64
	expectedOutcome int
}

var testCases = []TestCase{
	TestCase{3, 0.66, 1.5, 3},
	TestCase{40, 0.4, 10, 3},
	TestCase{10, 0.6, 10, -1},
	TestCase{40, 1, 10, -1},
	TestCase{5, -1, 1.5, -1},
}

func TestBouncingBall(t *testing.T) {
	for no, tc := range testCases {
		t.Run("testing case "+fmt.Sprintf("%d", no), func(t *testing.T) {
			assert.Equal(t, tc.expectedOutcome, BouncingBall(tc.height, tc.bounce, tc.windowHeight))
		})
	}
}
