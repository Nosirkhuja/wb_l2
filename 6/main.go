package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Cut options command line
type Cut struct {
	fields    string
	delimiter string
	separated bool
	result    []string
}

// New constructor return Cut instance
func New(fields string, delimiter string, separated bool) *Cut {
	return &Cut{
		fields:    fields,
		delimiter: delimiter,
		separated: separated,
	}
}

//Execute writes the result in struct and returns an error
func (c *Cut) Execute() error {
	matches := make([]string, 0)

	fields := strings.Split(c.fields, ",")

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadBytes('\n')

		if err == io.EOF {
			break
		}

		line = line[:len(line)-1]

		lineSliceByDelimiter := bytes.Split(line, []byte(c.delimiter))

		newLine := make([][]byte, 0)

		match, _ := regexp.Match(c.delimiter, line)
		if match {
			for _, v := range fields {
				field, err := strconv.Atoi(v)
				if err != nil {
					return errors.New("invalid field value")
				}
				if field <= 0 {
					return errors.New("fields are numbered from 1")
				}
				if field <= len(lineSliceByDelimiter) {
					newLine = append(newLine, lineSliceByDelimiter[field-1])
				}
			}
			normalizeLine := bytes.Join(newLine, []byte(c.delimiter))
			matches = append(matches, string(normalizeLine))
		} else {
			if !c.separated {
				matches = append(matches, string(line))
			}
		}
	}

	c.result = matches

	return nil
}

// Output writes the result in Stdout and returns an error
func (c *Cut) Output() error {
	_, err := fmt.Fprintln(os.Stdout, strings.Join(c.result, "\n"))
	return err
}

func usage() {
	log.Printf("Usage: ./cut [OPTION]... STDIN... \n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var fields = flag.String("f", "", "выбрать поля (колонки)")
	var delimiter = flag.String("d", "", "использовать другой разделитель")
	var separated = flag.Bool("s", false, "только строки с разделителем")

	var showHelp = flag.Bool("h", false, "Show help message")

	if *delimiter == "" {
		*delimiter = "\t"
	}

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	if len(*delimiter) > 1 {
		log.Println("cut: the delimiter must be a single character")
		showUsageAndExit(1)
	}

	if len(*fields) < 1 {
		log.Println("cut: you must specify a list of fields")
		showUsageAndExit(1)
	}

	cut := New(*fields, *delimiter, *separated)

	err := cut.Execute()
	if err != nil {
		log.Fatalf("cut: %s", err)
	}

	err = cut.Output()
	if err != nil {
		log.Fatalf("cut: %s", err)
	}
}
