package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestFileInput(t *testing.T) {
	var testFile *fileMeta = setupFile()
	t.Run("file Content to be as expected", func(t *testing.T) {
		err := testFile.GetFileContents()
		want := []string{"one", "testdata.txt", "two"}

		if err != nil {
			t.Errorf("Got an error but did not expect one. %s", err.Error())
		}

		if !reflect.DeepEqual(testFile.fileContents, want) {
			t.Errorf("Got %v, wanted %v", testFile.fileContents, want)
		}
	})

	t.Run("file Content not to be as expected", func(t *testing.T) {
		err := testFile.GetFileContents()
		want := []string{"one", "testdata.txt"}

		if err != nil {
			t.Errorf("Got an error but did not expect one. %s", err.Error())
		}

		if reflect.DeepEqual(testFile.fileContents, want) {
			t.Errorf("Expected results to differ, which they do not.")
		}
	})

}

func TestConsoleOutput(t *testing.T) {

	t.Run("Test result output", func(t *testing.T) {
		b := bytes.Buffer{}
		flags, file := setupResultTest()
		printResults(flags, file, &b)
		//time.Sleep(5 * time.Second)
		got := b.String()
		want := append(file.fileContents)

		fmt.Printf("got %s", got)
		for _, substr := range want {
			if !strings.Contains(got, substr) {
				t.Errorf("Did not find one (ore more) required strings ('%s')", substr)
			}
		}
	})
}

func setupResultTest() (*flags, *fileMeta) {
	file := fileMeta{
		fileName:     "./testData/testdata.txt",
		lineCount:    3,
		fileContents: []string{"one", "testdata.txt", "two"},
	}

	flags := flags{
		filename: file.fileName,
		verbose:  true,
	}

	return &flags, &file
}

func setupFile() *fileMeta {
	testFile := fileMeta{
		fileName:     "./testData/testdata.txt",
		lineCount:    0,
		fileContents: nil,
	}

	return &testFile
}
