package main

import "strconv"

func main() {
	count := 0
	for i := 372304; i <= 847060; i++ {
		if checkPassword(strconv.Itoa(i)) {
			println(i)
			count++
		}
	}
	println(count)
	//println(checkPassword("233789"))
}

func checkPassword(password string) bool {
	if len(password) != 6 {
		return false
	}
	adjacentLetters, decreasing := false, false
	letter := password[0]
	for i := 1; i < len(password); i++ {
		if letter == password[i] {
			adjacentLetters = true
		}
		if letter > password[i] {
			decreasing = true
		}
		letter = password[i]
	}
	return adjacentLetters && !decreasing
}
