package multi_file_writer

import (
	"fmt"
	"os"
	"testing"
)

const jsonFileName = "logs/log1.json"
const textFileName = "logs/log1.txt"

var optionsWithBoth = LoggerOptions{
	LogJson:      true,
	LogText:      true,
	JsonFileName: jsonFileName,
	TextFileName: textFileName,
}
var optionsWithJsonOnly = LoggerOptions{
	LogJson:      true,
	LogText:      false,
	JsonFileName: jsonFileName,
	TextFileName: "",
}

var optionsWithTextOnly = LoggerOptions{
	LogJson:      false,
	LogText:      true,
	JsonFileName: "",
	TextFileName: textFileName,
}

func TestNewMultiFileWriter(t *testing.T) {
	t.Run("init with 2 files", func(t *testing.T) {
		writer := SetupTests(t, optionsWithBoth)

		t.Run("filnames match ", func(t *testing.T) {
			fmt.Printf("writer options %v \n", writer.Options)
			if writer.Options.TextFileName != textFileName || writer.Options.JsonFileName != jsonFileName {
				t.Error("File names do not match.")
			}
		})

		t.Run("files exist on FS", func(t *testing.T) {
			_, err := writer.JsonFileHandler.Stat()
			if err != nil {
				t.Error("Got an error but did not expect one: " + err.Error())
			}
			_, err = writer.TextFileHandler.Stat()
			if err != nil {
				t.Error("Got an error but did not expect one: " + err.Error())
			}
		})

		TestHelperCleanUp(t)
	})

	t.Run("init with only json file", func(t *testing.T) {
		writer := SetupTests(t, optionsWithJsonOnly)

		t.Run("filnames match ", func(t *testing.T) {
			fmt.Printf("writer options %v \n", writer.Options)
			if writer.Options.JsonFileName != jsonFileName {
				t.Error("File name does not match.")
			}

			if writer.Options.TextFileName != "" {
				t.Error("TextFileName exists but should not")
			}
		})

		t.Run("files exist on FS", func(t *testing.T) {
			_, err := writer.JsonFileHandler.Stat()
			if err != nil {
				t.Error("Got an error but did not expect one: " + err.Error())
			}
			_, err = writer.TextFileHandler.Stat()
			if err == nil {
				t.Error("File exists, but it should not ")
			}
		})

		TestHelperCleanUp(t)
	})

}
func SetupTests(t *testing.T, options LoggerOptions) *MultiFileWriter {
	writer, err := NewMultiFileWriter(options)
	if err != nil {
		t.Error("Failed to setup tests: " + err.Error())
	}
	return writer
}

func TestHelperCleanUp(t *testing.T) {

	_ = os.Remove(textFileName)
	_ = os.Remove(jsonFileName)

}
