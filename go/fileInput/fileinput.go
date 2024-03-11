package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type fileMeta struct {
	fileName     string
	lineCount    int
	fileContents []string
}

type flags struct {
	filename string
	verbose  bool
}

func (f *fileMeta) GetFileContents() error {
	file, err := os.Open(f.fileName)
	defer file.Close()

	if err != nil {
		return err
	}
	fileReader := bufio.NewReader(file)
	f.lineCount = 0

	for {
		line, err := fileReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}

		}
		line = strings.TrimSuffix(line, "\n")
		if line == "" {
			break
		} else {
			f.lineCount++
			f.fileContents = append(f.fileContents, line)
		}
	}

	return nil
}

func (f *flags) getFileNameFromFlags() error {
	fileNameFlag := flag.String("file", "", "Name of the file to process")
	verboseFlag := flag.Bool("verbose", false, "verbose mode: will show extra info (such as filename and line count)")

	flag.Parse()

	if *fileNameFlag == "" {

		return errors.New("please provide a file name")
	}

	f.verbose = *verboseFlag
	f.filename = *fileNameFlag

	return nil

}
func printResults(flags *flags, file *fileMeta, writer io.Writer) {
	if flags.verbose {
		fmt.Fprintf(writer, "Showing contents of %s: \n\n", file.fileName)
	}
	for _, line := range file.fileContents {
		fmt.Fprintln(writer, line)
	}
	if flags.verbose {
		fmt.Fprintf(writer, "\nTotal of %d lines \n", file.lineCount)
	}
}

func main() {
	flags := flags{}
	err := flags.getFileNameFromFlags()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	file := fileMeta{
		fileName: flags.filename,
	}

	errGetContents := file.GetFileContents()

	if errGetContents != nil {
		fmt.Println(errGetContents.Error())
		os.Exit(1)
	} else {
		printResults(&flags, &file, os.Stdout)

	}

}
