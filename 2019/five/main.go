package main

import (
	"bufio"
	"bytes"
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

func main() {
	lines, _ := readLines("input")
	memory := strings.Split(lines[0], ",")
	_ = compute(memory)
}

func compute(prog []string) string {
	input := 1
	position := 0
	endEncountered := false
	for !endEncountered {
		if prog[position] == "99" {
			break
		}
		opCode, arg1, arg2, arg3 := parseInstruction(prog[position], prog, position)
		val1, _ := strconv.Atoi(prog[arg1])

		switch opCode {
		case "01":
			val2, _ := strconv.Atoi(prog[arg2])
			prog[arg3] = strconv.Itoa(val1 + val2)
			position += 4
		case "02":
			val2, _ := strconv.Atoi(prog[arg2])
			prog[arg3] = strconv.Itoa(val1 * val2)
			position += 4
		case "03":
			prog[arg1] = strconv.Itoa(input)
			position += 2
		case "04":
			println(prog[arg1])
			position += 2
		}
	}
	return prog[0]
}

func padOperation(instruction string) string {
	for len(instruction) < 5 {
		instruction = "0" + instruction
	}
	return instruction
}

func parseInstruction(instruction string, memory []string, position int) (string, int, int, int) {
	paddedOp := padOperation(instruction)
	var opCode string
	var arg1, arg2, arg3 int
	opCode = string(paddedOp[3:5])
	if string(paddedOp[2]) == "1" {
		arg1 = position + 1
	} else {
		arg1, _ = strconv.Atoi(memory[position+1])
	}
	if string(paddedOp[1]) == "1" {
		arg2 = position + 2
	} else {
		arg2, _ = strconv.Atoi(memory[position+2])
	}
	if string(paddedOp[0]) == "1" {
		arg3 = position + 3
	} else {
		arg3, _ = strconv.Atoi(memory[position+3])
	}
	return opCode, arg1, arg2, arg3
}
