package multi_file_writer

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type MultiFileWriterAPI interface {
	io.Writer
	Close()
}

type MultiFileWriter struct {
	Options         LoggerOptions
	JsonFileHandler *os.File
	TextFileHandler *os.File
}

type LoggerOptions struct {
	LogJson      bool
	JsonFileName string
	LogText      bool
	TextFileName string
}
type LogMsgStructure struct {
	File    string `json:"File"`
	Line    int    `json:"Line"`
	Message string `json:"Message"`
}

type JsonLogOutput struct {
	Data  LogMsgStructure
	Level string `json:"level"'`
	Msg   string `json:"msg"`
	Time  string `json:"time"`
}

func NewMultiFileWriter(options LoggerOptions) (*MultiFileWriter, error) {

	writer := MultiFileWriter{
		Options:         options,
		JsonFileHandler: nil,
		TextFileHandler: nil,
	}

	var err error
	if writer.Options.LogJson == true {
		writer.JsonFileHandler, err = os.OpenFile(writer.Options.JsonFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return nil, err
		}
	} else {
		writer.Options.JsonFileName = ""
	}

	if writer.Options.LogText == true {
		writer.TextFileHandler, err = os.OpenFile(writer.Options.TextFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return nil, err
		}
	} else {
		writer.Options.TextFileName = ""
	}

	return &writer, nil
}

func (fw *MultiFileWriter) Write(b []byte) (n int, err error) {
	var jsonMsg JsonLogOutput
	err = json.Unmarshal(b, &jsonMsg)
	if err != nil {
		return 0, err
	}

	if fw.Options.LogJson == true {
		_, err = fw.JsonFileHandler.Write(b)
	}

	if fw.Options.LogText == true {
		_, err = fw.TextFileHandler.Write([]byte(fmt.Sprintf("[%s] %s File: %s:%d - %s\n", jsonMsg.Level, jsonMsg.Time, jsonMsg.Data.File, jsonMsg.Data.Line, jsonMsg.Data.Message)))
	}

	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (fw *MultiFileWriter) Close() {
	if fw.Options.LogJson == true {
		_ = fw.JsonFileHandler.Close()
	}

	if fw.Options.LogText == true {
		_ = fw.TextFileHandler.Close()
	}
}
