package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	str   string
	count int
}

var testCases = []TestCase{
	{
		"abracadabra",
		5,
	},
	{
		"abracadabr",
		4,
	},
	{
		"aeiou",
		5,
	},
	{
		"aEioU",
		5,
	},
}

func TestGetCount(t *testing.T) {
	for _, tc := range testCases {
		t.Run("testing "+tc.str, func(t *testing.T) {
			assert.Equal(t, tc.count, GetCount(tc.str))
		})
	}
}
