package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	filename string
	expected string
}

func TestInjection(t *testing.T) {

	testcases := []TestCase{
		{
			filename: "latin1.txt",
			expected: "iso-8859-1",
		},
		{
			filename: "||\techo\tich\tbin\tdrin\t||\t.attribute",
			expected: "unsafe",
		},
		{
			filename: "||echo+hallo",
			expected: "unsafe",
		},
		{
			filename: "||echo%20hallo",
			expected: "unsafe",
		},
		{
			filename: "||[ ]echo[ ]ich",
			expected: "unsafe",
		},
		{
			filename: "||ls",
			expected: "unsafe",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("%d testcase", i), func(t *testing.T) {
			output := getFileEncoding(tc.filename)
			assert.Equal(t, tc.expected, output)
		})

	}

}

func Test_parseCharsetFromCmdOutput(t *testing.T) {
	testcases := []struct {
		output          string
		expectedCharset string
	}{
		{
			output:          "latin-1.txt text/plain; charset=iso-8852-1",
			expectedCharset: "iso-8852-1",
		},
		{
			output:          "latin-1.txt text/plain; charset=ISO-8852-1",
			expectedCharset: "ISO-8852-1",
		},
		{
			output:          "latin-1.txt text/plain; charset=utf-8",
			expectedCharset: "utf-8",
		},
		{
			output:          "latin-1.txt text/plain; charset=UTF-8",
			expectedCharset: "UTF-8",
		},
		{
			output:          "latin-1.txt text/plain; charset=UTF-16",
			expectedCharset: "UTF-16",
		},
	}
	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("%d) case %s", i, testcase.expectedCharset), func(t *testing.T) {
			got, err := parseCharsetFromCmdOutput(testcase.output)

			assert.NoError(t, err)
			assert.Equal(t, testcase.expectedCharset, got)
		})
	}

}
