package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
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

func step(currentPos string, direction string) (result string) {
	temp := strings.Split(currentPos, ",")
	x, _ := strconv.Atoi(temp[0])
	y, _ := strconv.Atoi(temp[1])
	switch direction {
	case "^":
		x--
	case "v":
		x++
	case ">":
		y++
	case "<":
		y--
	}
	result = strconv.Itoa(x) + "," + strconv.Itoa(y)
	return
}

func one() {
	lines, _ := readLines("D:\\code\\advent-master\\2015\\three\\three.in")
	instructions := lines[0]
	positions := make(map[string]bool)
	positions["0,0"] = true
	last := "0,0"
	for _, com := range strings.Split(instructions, "") {
		last = step(last, com)
		positions[last] = true
	}
	fmt.Println(len(positions))
}

func main() {
	lines, _ := readLines("D:\\code\\advent-master\\2015\\three\\three.in")
	instructions := lines[0]
	positions := make(map[string]bool)
	positions["0,0"] = true
	santa, roboSanta := "0,0", "0,0"
	for index, com := range strings.Split(instructions, "") {
		if index%2 == 0 {
			santa = step(santa, com)
			positions[santa] = true
		} else {
			roboSanta = step(roboSanta, com)
			positions[roboSanta] = true
		}
	}
	fmt.Println(len(positions))

}
