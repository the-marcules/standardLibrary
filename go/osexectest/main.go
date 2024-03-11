package main

import (
	"errors"
	"log"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func filenameWithNeitherPathOrExtension(filename string) string {
	fileExt := filepath.Ext(filename)
	filePath := filepath.Dir(filename)
	filename = strings.TrimPrefix(filename, filePath)
	filename = strings.TrimSuffix(filename, fileExt)
	return filename
}

func isGoodFilename(filename string) bool {

	filename = filenameWithNeitherPathOrExtension(filename)

	regex, err := regexp.Compile(`^[a-zA-Z0-9]+$`)
	if err != nil {
		return false
	}
	return regex.Match([]byte(filename))

}

func parseCharsetFromCmdOutput(output string) (string, error) {
	if output == "" {
		return "", errors.New("output empty")
	}

	if !strings.Contains(output, "charset=") {
		return "", errors.New("no charset found")
	}
	re := regexp.MustCompile(`(?m)[a-zA-Z]+(-[a-zA-Z0-9]+)+$`)
	charset := re.FindString(output)

	if charset == "" {
		return charset, errors.New("could not determine charset")
	}
	return charset, nil
}

func getFileEncoding(fileName string) string {
	if !isGoodFilename(fileName) {
		return "unsafe"
	}
	fileName = filenameWithNeitherPathOrExtension(fileName)
	filenameWithExtension := fileName + ".txt"
	out, err := exec.Command("file", "-I", filenameWithExtension).Output()
	if err != nil {
		log.Fatal(err)
	}
	charset, err := parseCharsetFromCmdOutput(string(out))
	if err != nil {
		log.Fatal(err)
	}
	return charset
}
