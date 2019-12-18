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
	parent   *planet
}

var topNode planet
var lines []string
var you planet
var target planet

func main() {
	lines, _ = readLines("input")
	fmt.Printf("%v", childrenOf("COM"))
	topNode = planet{name: "COM", depth: 0}
	setChildrenFor("COM", &topNode)
	println(aTotal)
	println(you.depth)
	println(target.depth)
	//myParents := parentOf(&you)
	for _, x := range parentOf(&target) {
		fmt.Printf("%+v \n", x)
	}

}

var aTotal = 0

func setChildrenFor(planetName string, planetNode *planet) {
	aTotal += planetNode.depth
	children := childrenOf(planetName)
	for i := range children {
		newPlanet := planet{name: children[i], parent: planetNode, depth: planetNode.depth + 1}
		if newPlanet.name == "YOU" {
			you = newPlanet
		}
		if newPlanet.name == "SAN" {
			target = newPlanet
		}

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

func parentOf(node *planet) []*planet {
	if node.name == "COM" {
		return []*planet{}
	}
	x := []*planet{node}
	return append(x, parentOf(node.parent)...)
}
