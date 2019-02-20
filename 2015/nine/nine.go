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

var distances = make(map[string]map[string]int)
var visited = make(map[string]bool)

func processDistances(lines []string) {
	for _, line := range lines {
		tokens := strings.Split(line, " = ")
		cities := strings.Split(tokens[0], " to ")
		if distances[cities[0]] == nil {
			distances[cities[0]] = make(map[string]int)
		}
		if distances[cities[1]] == nil {
			distances[cities[1]] = make(map[string]int)
		}
		distances[cities[0]][cities[1]], _ = strconv.Atoi(tokens[1])
		distances[cities[1]][cities[0]], _ = strconv.Atoi(tokens[1])
	}
}

func generateRoute(startingCity string, lengthSoFar int) int {
	if len(visited) == len(distances) {
		return lengthSoFar
	}
	dists := distances[startingCity]
	nextCity, min := "", 99999999
	for city, length := range dists {
		if !visited[city] && length < min {
			nextCity = city
			min = length
		}
	}
	visited[nextCity] = true
	fmt.Println("Selecting" + nextCity)
	return lengthSoFar + generateRoute(nextCity, min)
}

func main() {
	lines, _ := readLines("/home/alex/projects/advent/2015/nine/nine.in")
	fmt.Println(lines)
	processDistances(lines)
	minDistance := 99999999
	for city := range distances {
		visited = make(map[string]bool)
		visited[city] = true
		newRouteCost := generateRoute(city, 0)
		if minDistance > newRouteCost {
			minDistance = newRouteCost
		}
	}
	fmt.Println(minDistance)
}
