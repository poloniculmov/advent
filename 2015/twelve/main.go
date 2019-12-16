package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
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

var intMatcher = regexp.MustCompile(`(-?\d*)`)

func sumNumbers(jason string) int {
	matches := intMatcher.FindAllString(jason, -1)
	sum := 0
	for _, x := range matches {
		val, _ := strconv.Atoi(x)
		sum += val
	}
	return sum
}

func main() {
	lines, _ := readLines("/home/alex/projects/advent/2015/twelve/twelve.in")
	fmt.Println(sumNumbers(lines[0]))
	b := []byte(lines[0])
	var f interface{}
	_ = json.Unmarshal(b, &f)
	fmt.Println(f)
}
