package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
)

func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

var doubleQuoteMatcher = regexp.MustCompile(`(^")|("$)`)
var quoteMatcher = regexp.MustCompile(`\\"`)
var slashMatcher = regexp.MustCompile(`\\\\`)
var hexMatcher = regexp.MustCompile(`\\x[0-9a-f][0-9a-f]`)

func stripCruft(input string) string {
	result := doubleQuoteMatcher.ReplaceAllString(input, "")
	result = quoteMatcher.ReplaceAllString(result, "!")
	result = hexMatcher.ReplaceAllString(result, "?")
	result = slashMatcher.ReplaceAllString(result, "+")
	return result
}

func main() {
	lines, _ := readLines("D:\\code\\advent\\2015\\eight\\eight.in")
	sum := 0
	for _, line := range lines {
		fmt.Println(line)
		fmt.Println(stripCruft(line))
		sum += len(line) - len(stripCruft(line))
	}
	fmt.Println(sum)
}
