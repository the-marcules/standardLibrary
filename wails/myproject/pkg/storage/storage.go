package storage

import (
	"fmt"
	"os"
)

type Response string

func (s Response) string() string {
	return string(s)
}

const FileWrittenResponse = Response("File written successfully")
const FileOperationFailed = Response("File operation failed")

type Storage struct {
	FileName string
}

func NewStorage(fileName string) *Storage {
	return &Storage{
		FileName: fileName,
	}
}

func (s *Storage) writeToFile(input string) string {
	if input != "" {
		input = fmt.Sprintf("%s\n", input)
	}
	byteInput := []byte(input)

	err := os.WriteFile(s.FileName, byteInput, 0644)

	if err != nil {
		f, err := os.Create(s.FileName)
		if err != nil {
			return fmt.Sprintf("%s: %s", FileOperationFailed.string(), err.Error())
		}

		defer f.Close()

		err = os.WriteFile(s.FileName, byteInput, 0644)
		if err != nil {
			return fmt.Sprintf("%s: %s", FileOperationFailed.string(), err.Error())
		}
	}

	return FileWrittenResponse.string()
}

func (s *Storage) Write(input string) string {
	_, err := os.Stat(s.FileName)

	if err != nil {
		return s.writeToFile(input)
	}

	fileContent := s.Read()

	return s.writeToFile(fmt.Sprintf("%s%s", fileContent, input))
}

func (s *Storage) Read() string {
	content, err := os.ReadFile(s.FileName)
	if err != nil {
		return "-1"
	}
	return string(content)
}

func (s *Storage) Truncate() string {
	return s.writeToFile("")
}
