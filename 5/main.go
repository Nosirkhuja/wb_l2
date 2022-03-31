package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type args struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

// Grep options command line
type Grep struct {
	args
	pattern string
	files   []string
	result  []string
}

// New constructor return Grep instance
func New(after int, before int, context int, count bool, ignoreCase bool, invert bool, fixed bool, lineNum bool, pattern string, files []string) *Grep {
	return &Grep{
		args: args{
			after:      after,
			before:     before,
			context:    context,
			count:      count,
			ignoreCase: ignoreCase,
			invert:     invert,
			fixed:      fixed,
			lineNum:    lineNum,
		},

		pattern: pattern,
		files:   files,
	}
}

//Execute writes the result in struct and returns an error
func (g *Grep) Execute() error {
	matches := make([]string, 0)

	if g.ignoreCase {
		g.pattern = "(?i)" + g.pattern
	}

	for _, file := range g.files {
		fileMatches, err := searchFile(g.pattern, g.args, file)
		if err != nil {
			return err
		}

		if len(g.files) > 1 {
			for _, m := range fileMatches {
				matches = append(matches, file+":"+m)
			}
		} else {
			matches = append(matches, fileMatches...)
		}
	}

	g.result = matches

	return nil
}

// Output writes the result in Stdout and returns an error
func (g *Grep) Output() error {
	_, err := fmt.Fprintln(os.Stdout, strings.Join(g.result, "\n"))
	return err
}

func searchFile(pattern string, options args, path string) ([]string, error) {
	matches := make([]string, 0)
	reader, err := open(path)
	if err != nil {
		return nil, err
	}

	inverted := options.invert
	var i int
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		i++
		line = line[:len(line)-1]
		if match, _ := regexp.Match(pattern, line); match != inverted {
			if options.lineNum {
				line = append([]byte(fmt.Sprintf("%d:", i)), line...)
			}
			if options.fixed {
				regex := regexp.MustCompile(pattern)
				out := regex.ReplaceAll(line, []byte(fmt.Sprintf("\033[1;34m%s\033[0m", pattern)))
				line = out
			}

			matches = append(matches, string(line))
		}
	}

	if options.count {
		strLenOfMatch := fmt.Sprint(len(matches))
		sliceLenOfMatch := strings.Split(strLenOfMatch, "")
		matches = sliceLenOfMatch
	}

	return matches, nil
}

func open(path string) (*bufio.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("no such file or directory")
	}

	return bufio.NewReader(file), nil
}

func usage() {
	log.Printf("Usage: ./grep [OPTION]... PATTERN [FILE]... \n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var after = flag.Int("A", 0, "печатать +N строк после совпадения")
	var before = flag.Int("B", 0, "печатать +N строк до совпадения")
	var context = flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	var count = flag.Bool("c", false, "количество строк")                         // +
	var ignoreCase = flag.Bool("i", false, "игнорировать регистр")                // +
	var invert = flag.Bool("v", false, "вместо совпадения, исключать")            // +
	var fixed = flag.Bool("F", false, "точное совпадение со строкой, не паттерн") // +
	var lineNum = flag.Bool("n", false, "печатать номер строки")                  // +

	var showHelp = flag.Bool("h", false, "Show help message") // +

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	args := flag.Args()

	if len(args) < 2 {
		showUsageAndExit(1)
	}

	var pattern = flag.Args()[0]
	var files = flag.Args()[1:]

	Grep := New(*after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum, pattern, files)

	err := Grep.Execute()
	if err != nil {
		log.Fatalf("grep: %s", err)
	}

	err = Grep.Output()
	if err != nil {
		log.Fatalf("grep: %s", err)
	}
}
