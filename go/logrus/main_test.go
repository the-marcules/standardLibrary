package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	. "logrus/pkg/multi_file_writer"
	"os"
	"testing"
)

var message = LogMsgStructure{
	File:    "sido.mme",
	Line:    1,
	Message: "Mama sagt nicht leider; Mama sagt 'OK'",
}

const jsonFileName = "logs/log.json"
const textFileName = "logs/log.txt"

var loggerOptions = LoggerOptions{
	LogJson:      true,
	LogText:      true,
	JsonFileName: jsonFileName,
	TextFileName: textFileName,
}

var loggerOptionsTxtOnly = LoggerOptions{
	LogJson:      false,
	LogText:      true,
	JsonFileName: "",
	TextFileName: textFileName,
}

var logMsg = LogMsgStructure{
	File:    "bla.txt",
	Line:    1,
	Message: "das ist ein TEST",
}

type mockMultiWriter struct {
	Options          LoggerOptions
	JsonFileHandler  *os.File
	TextFileHandler  *os.File
	WriteInvocations int
	CloseInvocations int
	LastLogMsg       []byte
}

func (m *mockMultiWriter) Write(b []byte) (n int, err error) {
	m.WriteInvocations++
	m.LastLogMsg = b
	fmt.Printf("called\n")
	return 0, nil
}

func (m *mockMultiWriter) Close() {
	m.CloseInvocations++
}

func TestInitLogger(t *testing.T) {
	writer, _ := NewMultiFileWriter(loggerOptions)
	_, err := InitLogger(writer)

	t.Run("throws no error", func(t *testing.T) {
		if err != nil {
			t.Error("got err\n")
		}
		TestHelperCleanUp(t)
	})

}

func TestLogInfo(t *testing.T) {
	writer, _ := NewMultiFileWriter(loggerOptions)
	logger, err := InitLogger(writer)

	if err != nil {
		t.Error("got an init error but did not expect one: " + err.Error())
	}

	logMsgStr, _ := json.Marshal(logMsg)
	err = logger.LogInfo(logMsg)
	assert.NoError(t, err)
	textFileContent := getLogFileContents(loggerOptions.TextFileName)
	jsonFileContent := getLogFileContents(loggerOptions.JsonFileName)

	assert.Contains(t, jsonFileContent, string(logMsgStr))
	assert.Contains(t, jsonFileContent, "\"level\":\"info\"")

	assert.Contains(t, textFileContent, logMsg.Message)
	assert.Contains(t, textFileContent, fmt.Sprintf("%s:%d", logMsg.File, logMsg.Line))
	assert.Contains(t, textFileContent, "[info]")
	TestHelperCleanUp(t)
}

func TestWriteMocked(t *testing.T) {
	writer := &mockMultiWriter{
		Options: LoggerOptions{
			TextFileName: textFileName,
			LogText:      true,
			JsonFileName: jsonFileName,
			LogJson:      true,
		},
		JsonFileHandler: nil,
		TextFileHandler: nil,
	}
	msgStr, _ := json.Marshal(message)

	writer.Write(msgStr)

	assert.Equal(t, 1, writer.WriteInvocations)
	assert.Contains(t, string(writer.LastLogMsg), string(msgStr))
}

func TestLogInfoMocked(t *testing.T) {
	writer := &mockMultiWriter{
		Options: LoggerOptions{
			TextFileName: textFileName,
			LogText:      true,
			JsonFileName: jsonFileName,
			LogJson:      false,
		},
		JsonFileHandler: nil,
		TextFileHandler: nil,
	}
	msgStr, _ := json.Marshal(message)

	logger, _ := InitLogger(writer)
	logger.LogInfo(message)
	assert.Equal(t, 1, writer.WriteInvocations)
	assert.Contains(t, string(writer.LastLogMsg), string(msgStr))
}

func TestLogger_LoggerClose(t *testing.T) {
	writer, _ := NewMultiFileWriter(loggerOptionsTxtOnly)

	logger, _ := InitLogger(writer)

	logger.LoggerClose()

	logMsgStr, _ := json.Marshal(logMsg)

	_, err := logger.writer.Write([]byte(logMsgStr))

	assert.Contains(t, err.Error(), fmt.Sprintf("%s: file already closed", loggerOptions.TextFileName))
	if err == nil {
		t.Error("logger not closed successfully")
	}

	TestHelperCleanUp(t)
}

func TestLoggerCloseMock(t *testing.T) {
	writer := &mockMultiWriter{
		Options: LoggerOptions{
			TextFileName: textFileName,
			LogText:      true,
			JsonFileName: jsonFileName,
			LogJson:      false,
		},
		JsonFileHandler: nil,
		TextFileHandler: nil,
	}

	logger, _ := InitLogger(writer)
	logger.LoggerClose()
	assert.Equal(t, 1, writer.CloseInvocations)
}

func TestHelperCleanUp(t *testing.T) {
	_ = os.Remove(jsonFileName)
	_ = os.Remove(textFileName)
}

func getLogFileContents(filename string) string {
	content, _ := os.ReadFile(filename)

	return string(content)
}
