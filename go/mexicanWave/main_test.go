package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Input           string
	ExpectedOutcome []string
}

var testCases = []TestCase{
	TestCase{
		Input:           " x yz",
		ExpectedOutcome: []string{" X yz", " x Yz", " x yZ"},
	},
	TestCase{
		Input:           "abc",
		ExpectedOutcome: []string{"Abc", "aBc", "abC"},
	},
	TestCase{
		Input:           "abc",
		ExpectedOutcome: []string{"Abc", "aBc", "abC"},
	},
	TestCase{
		Input:           " ab  c",
		ExpectedOutcome: []string{" Ab  c", " aB  c", " ab  C"},
	},
	TestCase{
		Input:           "",
		ExpectedOutcome: []string{},
	},
	TestCase{
		Input:           "z",
		ExpectedOutcome: []string{"Z"},
	},
	TestCase{
		Input:           "a a a a a",
		ExpectedOutcome: []string{"A a a a a", "a A a a a", "a a A a a", "a a a A a", "a a a a A"},
	},
	TestCase{
		Input:           "aaaaa",
		ExpectedOutcome: []string{"Aaaaa", "aAaaa", "aaAaa", "aaaAa", "aaaaA"},
	},
	TestCase{
		Input:           "                                                           ",
		ExpectedOutcome: []string{},
	},
}

func TestWave(t *testing.T) {

	for _, tc := range testCases {
		t.Run("testing string "+tc.Input, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedOutcome, wave(tc.Input))
		})
	}

}
