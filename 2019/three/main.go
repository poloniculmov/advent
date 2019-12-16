package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
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

func main() {
	lines, _ := readLines("input")
	paths := [40000][40000]int{}
	firstWire := strings.Split(lines[0], ",")
	secondWire := strings.Split(lines[1], ",")
	x, y := 10000, 10000
	for step := range firstWire {
		dir := string(firstWire[step][0])
		size, _ := strconv.Atoi(firstWire[step][1:])
		switch dir {
		case "R":
			for arg := y; arg <= y+size; arg++ {
				paths[x][arg] += 1
			}
			y += size
		case "L":
			for arg := y; arg >= y-size; arg-- {
				paths[x][arg] += 1
			}
			y -= size
		case "D":
			for arg := x; arg <= x+size; arg++ {
				paths[arg][y] += 1
			}
			x += size
		case "U":
			for arg := x; arg >= x-size; arg-- {
				paths[arg][y] += 1
			}
			x -= size
		}
	}
	x, y = 10000, 10000

	for step := range secondWire {
		dir := string(secondWire[step][0])
		size, _ := strconv.Atoi(secondWire[step][1:])
		switch dir {
		case "R":
			for arg := y; arg <= y+size; arg++ {
				paths[x][arg] += 100
			}
			y += size
		case "L":
			for arg := y; arg >= y-size; arg-- {
				paths[x][arg] += 100
			}
			y -= size
		case "D":
			for arg := x; arg <= x+size; arg++ {
				paths[arg][y] += 100
			}
			x += size
		case "U":
			for arg := x; arg >= x-size; arg-- {
				paths[arg][y] += 100
			}
			x -= size
		}
	}
	var vals []int
	for i := range paths {
		for j := range paths[i] {
			if paths[i][j] > 100 && paths[i][j] < 200 {
				vals = append(vals, int(math.Abs(float64(i-10000))+math.Abs(float64(j-10000))))
			}
		}
	}
	sort.Ints(vals)
	fmt.Printf("%v", vals)
}
