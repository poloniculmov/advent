package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
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

func isNice(line string) (nice bool) {
	nice = false
	vowelCount := 0
	prevLetter := ""
	doubleLetter := false
	if badString.MatchString(line) {
		return false
	}
	for i := 0; i < len(line); i++ {
		currentLetter := string(line[i])
		if currentLetter == prevLetter {
			doubleLetter = true
		}
		prevLetter = currentLetter
		if vowelMatcher.MatchString(currentLetter) {
			vowelCount++
		}
	}
	if doubleLetter == true && vowelCount >= 3 {
		nice = true
	}
	return
}

var vowelMatcher = regexp.MustCompile(`[aeiou]`)
var badString = regexp.MustCompile(`ab|cd|pq|xy`)

func isNicer(line string) bool {
	foundPair, foundDouble := false, false
	for i := 0; i < len(line)-2; i++ {
		currentGroup := string(line[i]) + string(line[i+1])
		if strings.Contains(line[i+2:], currentGroup) {
			foundPair = true
		}
		if line[i] == line[i+2] {
			foundDouble = true
		}
	}
	return foundDouble && foundPair
}

func main() {
	lines, _ := readLines("/home/alex/projects/advent/2015/five/five.in")
	count, nicerCount := 0, 0
	for _, line := range lines {
		if isNice(line) == true {
			count++
		}
		if isNicer(line) == true {
			nicerCount++
		}
	}
	fmt.Println(count)
	fmt.Println(nicerCount)
}
