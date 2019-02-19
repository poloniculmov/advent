package main

import (
	"bufio"
	"bytes"
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

var lights [1000][1000]bool
var newLights [1000][1000]int

func turnOn(startX int, startY int, stopX int, stopY int) {
	for i := startX; i <= stopX; i++ {
		for j := startY; j <= stopY; j++ {
			lights[i][j] = true
		}
	}
}

func turnOff(startX int, startY int, stopX int, stopY int) {
	for i := startX; i <= stopX; i++ {
		for j := startY; j <= stopY; j++ {
			lights[i][j] = false
		}
	}
}

func toggle(startX int, startY int, stopX int, stopY int) {
	for i := startX; i <= stopX; i++ {
		for j := startY; j <= stopY; j++ {
			lights[i][j] = !lights[i][j]
		}
	}
}

func howManyOn() (size int) {
	size = 0
	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			if lights[i][j] == true {
				size++
			}
		}
	}
	return
}

var offMatch = regexp.MustCompile(`turn off (\d*),(\d*) through (\d*),(\d*)`)
var onMatch = regexp.MustCompile(`turn on (\d*),(\d*) through (\d*),(\d*)`)
var toggleMatch = regexp.MustCompile(`toggle (\d*),(\d*) through (\d*),(\d*)`)

func doCommand(command string) {
	var match = offMatch.FindStringSubmatch(command)
	if len(match) > 1 {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		c, _ := strconv.Atoi(match[3])
		d, _ := strconv.Atoi(match[4])
		turnOff(a, b, c, d)
	}
	match = onMatch.FindStringSubmatch(command)
	if len(match) > 1 {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		c, _ := strconv.Atoi(match[3])
		d, _ := strconv.Atoi(match[4])
		turnOn(a, b, c, d)
	}
	match = toggleMatch.FindStringSubmatch(command)
	if len(match) > 1 {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		c, _ := strconv.Atoi(match[3])
		d, _ := strconv.Atoi(match[4])
		toggle(a, b, c, d)
	}
}

func newTurnOn(startX int, startY int, stopX int, stopY int) {
	for i := startX; i <= stopX; i++ {
		for j := startY; j <= stopY; j++ {
			newLights[i][j]++
		}
	}
}

func newTurnOff(startX int, startY int, stopX int, stopY int) {
	for i := startX; i <= stopX; i++ {
		for j := startY; j <= stopY; j++ {
			newLights[i][j]--
			if newLights[i][j] < 0 {
				newLights[i][j] = 0
			}
		}
	}
}

func newToggle(startX int, startY int, stopX int, stopY int) {
	for i := startX; i <= stopX; i++ {
		for j := startY; j <= stopY; j++ {
			newLights[i][j] += 2
		}
	}
}

func newDoCommand(command string) {
	var match = offMatch.FindStringSubmatch(command)
	if len(match) > 1 {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		c, _ := strconv.Atoi(match[3])
		d, _ := strconv.Atoi(match[4])
		newTurnOff(a, b, c, d)
	}
	match = onMatch.FindStringSubmatch(command)
	if len(match) > 1 {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		c, _ := strconv.Atoi(match[3])
		d, _ := strconv.Atoi(match[4])
		newTurnOn(a, b, c, d)
	}
	match = toggleMatch.FindStringSubmatch(command)
	if len(match) > 1 {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		c, _ := strconv.Atoi(match[3])
		d, _ := strconv.Atoi(match[4])
		newToggle(a, b, c, d)
	}
}

func totalBrightnes() (brightness int) {
	brightness = 0
	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			brightness += newLights[i][j]
		}
	}
	return
}

func main() {
	lines, _ := readLines("/home/alex/projects/advent/2015/six/six.in")
	for _, line := range lines {
		doCommand(line)
		newDoCommand(line)
	}
	fmt.Println(howManyOn())
	fmt.Println(totalBrightnes())
}
