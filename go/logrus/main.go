package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	. "logrus/pkg/multi_file_writer"
)

type Logger struct {
	Logger *logrus.Logger
	writer MultiFileWriterAPI
}

func InitLogger(writer MultiFileWriterAPI) (*Logger, error) {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}

	logger.SetOutput(writer)

	return &Logger{
		Logger: logger,
		writer: writer,
	}, nil
}

func (l *Logger) LogInfo(logMsg LogMsgStructure) error {
	l.Logger.WithFields(logrus.Fields{
		"Data": logMsg,
	}).Info(logMsg.Message)

	return nil
}

func (l *Logger) LoggerClose() {
	l.writer.Close()
}

func main() {
	writer, writerErr := NewMultiFileWriter(LoggerOptions{
		LogJson:      true,
		LogText:      true,
		JsonFileName: "log.json",
		TextFileName: "log.txt",
	})
	if writerErr != nil {
		fmt.Println("Error on initializing writer: ", writerErr.Error())
		return
	}

	logger, err := InitLogger(writer)
	if err != nil {
		fmt.Println("Error on Initializing logger: ", err.Error())
		return
	}

	logger.LogInfo(LogMsgStructure{File: "some File", Message: "Hello Program Start.", Line: 1})
}
