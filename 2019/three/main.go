package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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
	s := 0
	for step := range firstWire {
		dir := string(firstWire[step][0])
		size, _ := strconv.Atoi(firstWire[step][1:])
		switch dir {
		case "R":
			for arg := y; arg <= y+size; arg++ {
				if paths[x][arg] == 0 {
					paths[x][arg] = s
				}
				s++
			}
			y += size
		case "L":
			for arg := y; arg >= y-size; arg-- {
				if paths[x][arg] == 0 {
					paths[x][arg] = s
				}
				s++
			}
			y -= size
		case "U":
			for arg := x; arg <= x+size; arg++ {
				if paths[arg][y] == 0 {
					paths[arg][y] = s
				}
				s++
			}
			x += size
		case "D":
			for arg := x; arg >= x-size; arg-- {
				if paths[arg][y] == 0 {
					paths[arg][y] = s
				}
				s++
			}
			x -= size
		}
		s--
	}
	x, y = 10000, 10000
	s = 0
	var vals []int
	for step := range secondWire {
		dir := string(secondWire[step][0])
		size, _ := strconv.Atoi(secondWire[step][1:])
		switch dir {
		case "R":
			for arg := y; arg <= y+size; arg++ {
				if paths[x][arg] > 0 {
					vals = append(vals, s+paths[x][arg])
				}
				s++
			}
			y += size
		case "L":
			for arg := y; arg >= y-size; arg-- {
				if paths[x][arg] > 0 {
					vals = append(vals, s+paths[x][arg])
				}
				s++
			}
			y -= size
		case "U":
			for arg := x; arg <= x+size; arg++ {
				if paths[arg][y] > 0 {
					vals = append(vals, s+paths[arg][y])
				}
				s++
			}
			x += size
		case "D":
			for arg := x; arg >= x-size; arg-- {
				if paths[arg][y] > 0 {
					vals = append(vals, s+paths[arg][y])
				}
				s++
			}
			x -= size
		}
		s--
	}
	sort.Ints(vals)
	fmt.Printf("%v", vals)
}
