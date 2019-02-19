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

var gates = make(map[string]uint16)

func valueOfWire(wireLabel string) uint16 {
	val, ok := gates[wireLabel]
	if ok {
		return val
	}
	executeCommand(ops[wireLabel], wireLabel)
	return valueOfWire(wireLabel)
}

func executeCommand(command string, label string) {
	switch {
	case strings.Contains(command, "NOT"):
		doNot(command, label)
	case strings.Contains(command, "AND"):
		doAnd(command, label)
	case strings.Contains(command, "OR"):
		doOr(command, label)
	case strings.Contains(command, "LSHIFT"):
		doLShift(command, label)
	case strings.Contains(command, "RSHIFT"):
		doRShift(command, label)
	default:
		doDirectAttribution(command, label)
	}
}

func doNot(command string, label string) {
	command = strings.Replace(command, "NOT ", "", -1)
	gates[label] = ^valueOfWire(command)
}
func doAnd(command string, label string) {
	command = strings.Replace(command, "AND ", "", -1)
	ws := strings.Split(command, " ")
	if ws[0] == "1" {
		gates[label] = 1 & valueOfWire(ws[1])
	} else {
		gates[label] = valueOfWire(ws[0]) & valueOfWire(ws[1])
	}
}
func doOr(command string, label string) {
	command = strings.Replace(command, "OR ", "", -1)
	ws := strings.Split(command, " ")
	gates[label] = valueOfWire(ws[0]) | valueOfWire(ws[1])
}
func doLShift(command string, label string) {
	command = strings.Replace(command, "LSHIFT ", "", -1)
	ws := strings.Split(command, " ")
	shiftBy, _ := strconv.Atoi(ws[1])
	gates[label] = valueOfWire(ws[0]) << uint16(shiftBy)
}
func doRShift(command string, label string) {
	command = strings.Replace(command, "RSHIFT ", "", -1)
	ws := strings.Split(command, " ")
	shiftBy, _ := strconv.Atoi(ws[1])
	gates[label] = valueOfWire(ws[0]) >> uint16(shiftBy)
}
func doDirectAttribution(command string, label string) {
	if command == "lx" {
		gates[label] = valueOfWire("lx")
	} else {
		val, _ := strconv.ParseUint(command, 10, 0)
		gates[label] = uint16(val)
	}
}

var lines []string
var ops = make(map[string]string)

func main() {
	lines, _ = readLines("/home/alex/projects/advent/2015/seven/seven.in")
	for _, l := range lines {
		tokens := strings.Split(l, " -> ")
		ops[tokens[1]] = tokens[0]
	}
	a := valueOfWire("a")
	fmt.Println(a)
	for _, l := range lines {
		tokens := strings.Split(l, " -> ")
		ops[tokens[1]] = tokens[0]
	}
	ops["b"] = strconv.Itoa(int(a))
	gates = make(map[string]uint16)
	a = valueOfWire("a")
	fmt.Println(a)
}
