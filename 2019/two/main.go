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
	lines, _ := readLines("a.in")
	//lines := [1]string{"1,9,10,3,2,3,11,0,99,30,40,50"}

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			memory := strings.Split(lines[0], ",")
			memory[1] = strconv.Itoa(noun)
			memory[2] = strconv.Itoa(verb)
			value := compute(memory)
			if value == "19690720" {
				println(100*noun + verb)
				println("found")
				break
			}
		}
	}

}

func compute(prog []string) string {
	position := 0
	endEncountered := false
	for !endEncountered {
		if prog[position] == "99" {
			break
		}
		arg1, _ := strconv.Atoi(prog[position+1])
		arg2, _ := strconv.Atoi(prog[position+2])
		arg3, _ := strconv.Atoi(prog[position+3])
		val1, _ := strconv.Atoi(prog[arg1])
		val2, _ := strconv.Atoi(prog[arg2])
		if arg3 > len(prog) {
			return "-1"
		}

		switch opCode := prog[position]; opCode {
		case "1":
			prog[arg3] = strconv.Itoa(val1 + val2)
		case "2":
			prog[arg3] = strconv.Itoa(val1 * val2)
		}
		position += 4
	}
	return prog[0]
}
