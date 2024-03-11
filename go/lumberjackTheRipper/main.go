package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	logFilePath := "main.log"

	logFilePath = "logs/" + logFilePath

	lumberjackLogger := &lumberjack.Logger{
		Filename: filepath.ToSlash(logFilePath),
		MaxAge:   3,
	}
	lumberjackLogger.Rotate()
	time.Sleep(time.Second * 5)

	mw := io.MultiWriter(os.Stdout, lumberjackLogger)
	log.SetOutput(mw)

}
