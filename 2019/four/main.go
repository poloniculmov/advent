package main

import (
	"regexp"
	"strconv"
)

func main() {
	count := 0
	for i := 372304; i <= 847060; i++ {
		if checkPassword(strconv.Itoa(i)) {
			//println(i)
			count++
		}
	}
	println(count)
	//println(checkPassword("233789"))
}

//
func checkPassword(password string) bool {
	if len(password) != 6 {
		return false
	}

	adjacentLetters, decreasing := false, false
	re := regexp.MustCompile(`(([^1]|^)1{2}([^1]|$))|(([^2]|^)2{2}([^2]|$))|(([^3]|^)3{2}([^3]|$))|(([^4]|^)4{2}([^4]|$))|(([^5]|^)5{2}([^5]|$))|(([^6]|^)6{2}([^6]|$))|(([^7]|^)7{2}([^7]|$))|(([^8]|^)8{2}([^8]|$))|(([^9]|^)9{2}([^9]|$))|(([^0]|^)0{2}([^0]|$))`)
	adjacentLetters = re.MatchString(password)
	letter := password[0]
	for i := 1; i < len(password); i++ {
		if letter > password[i] {
			decreasing = true
		}
		letter = password[i]
	}
	return adjacentLetters && !decreasing
}
