package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
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

type planet struct {
	name     string
	children []*planet
	depth    int
}

var topNode planet
var lines []string

func main() {
	lines, _ = readLines("input")
	fmt.Printf("%v", childrenOf("COM"))
	topNode = planet{name: "COM", depth: 0}
	setChildrenFor("COM", &topNode)
	setDepths(&topNode, 0)
	println(countOrbits(&topNode))
}

func setChildrenFor(planetName string, planetNode *planet) {
	children := childrenOf(planetName)
	for i := range children {
		newPlanet := planet{name: children[i]}
		planetNode.children = append(planetNode.children, &newPlanet)
		setChildrenFor(children[i], &newPlanet)
	}
}

func childrenOf(planetName string) []string {
	result := []string{}
	for i := range lines {
		temp := strings.Split(lines[i], ")")
		if planetName == temp[0] {
			result = append(result, temp[1])
		}
	}
	return result
}

func setDepths(node *planet, level int) {
	node.depth = level
	for _, child := range node.children {
		setDepths(child, level+1)
	}
}

func countOrbits(topNode *planet) int {
	queue := make([]*planet, 0)
	queue = append(queue, topNode)
	//level := 0
	size := 0
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:] //remove nextUp from queue
		size += nextUp.depth
		if len(nextUp.children) > 0 {
			for _, child := range nextUp.children {
				queue = append(queue, child)
			}
		}
	}
	return size
}
