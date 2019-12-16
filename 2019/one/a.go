package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
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

func main() {
	lines, _ := readLines("a.in")
	sum, sumb := 0, 0
	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		sum += fuel(i)
		sumb += totalFuel(i)
	}
	fmt.Println(sum, sumb)
}

func fuel(mass int) int {
	return int(math.Floor(float64(mass/3.0))) - 2
}

func fuelB(mass int) int {
	result := fuel(mass)
	if result > 0 {
		return result
	}
	return 0
}

func totalFuel(mass int) int {
	if mass <= 0 {
		return 0
	}
	needs := fuelB(mass)
	fmt.Println(needs)
	return needs + totalFuel(needs)
}
