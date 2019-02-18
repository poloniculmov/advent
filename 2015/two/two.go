package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func min(sizes [3]int) (m int) {
	for i, e := range sizes {
		if i == 0 || e < m {
			m = e
		}
	}
	return
}

func totalArea(width int, height int, length int) int {
	var face1 = width * height
	var face2 = width * length
	var face3 = height * length
	return 2*face1 + 2*face2 + 2*face3 + min([3]int{face1, face2, face3})
}

func ribbonSize(width int, height int, length int) int {
	faces := []int{2 * (width + height), 2 * (width + length), 2 * (height + length)}
	sort.Ints(faces)
	return faces[0] + width*height*length
}

func parseSize(size string) (x int, y int, z int) {
	var sizes = strings.Split(size, "x")
	x, _ = strconv.Atoi(sizes[0])
	y, _ = strconv.Atoi(sizes[1])
	z, _ = strconv.Atoi(sizes[2])
	return
}

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
	total, totalRibbon := 0, 0
	lines, err := readLines("D:\\code\\advent-master\\2015\\two\\two.in")
	for _, line := range lines {
		total += totalArea(parseSize(line))
		totalRibbon += ribbonSize(parseSize(line))
	}
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println(total)
	fmt.Println(totalRibbon)
}
