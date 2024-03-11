package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	isbn  string
	valid bool
}

var testCases = []TestCase{
	TestCase{"1112223339", true},
	TestCase{"048665088X", true},
	TestCase{"1293000000", true},
	TestCase{"1234554321", true},
	TestCase{"1234512345", false},
	TestCase{"1293", false},
	TestCase{"ABCDEFGHIJ", false},
	TestCase{"X123456788", false},
	TestCase{"XXXXXXXXXX", false},
	TestCase{"048665088x", true},
}

func TestValidISBN10(t *testing.T) {
	for _, tc := range testCases {
		t.Run("should be valid "+tc.isbn, func(t *testing.T) {
			assert.Equal(t, tc.valid, ValidISBN10(tc.isbn))
		})
	}

}
