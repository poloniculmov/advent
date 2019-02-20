package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func elvenSay(input string) string {
	oldLetter, letterCount := string(input[0]), 1
	var result bytes.Buffer
	for i := 1; i < len(input); i++ {
		if oldLetter == string(input[i]) {
			letterCount++
		} else {
			result.WriteString(strconv.Itoa(letterCount))
			result.WriteString(oldLetter)
			letterCount = 1
			oldLetter = string(input[i])
		}
	}
	result.WriteString(strconv.Itoa(letterCount))
	result.WriteString(oldLetter)
	return result.String()
}

func main() {
	input := "1321131112"
	for i := 0; i < 50; i++ {
		input = elvenSay(input)
	}
	fmt.Println(len(input))
}
