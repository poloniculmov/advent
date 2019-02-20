package main

import (
	"fmt"
	"regexp"
	"strings"
)

var illegalMatcher = regexp.MustCompile(`i|o|l`)
var straight = regexp.MustCompile(`(abc)|(bcd)|(cde)|(def)|(efg)|(fgh)|(ghi)|(hij)|(ijk)|(jkl)|(klm)|(lmn)|(mno)|(nop)|(opq)|(pqr)|(qrs)|(rst)|(stu)|(tuv)|(uvw)|(vwx)|(wxy)|(xyz)`)
var doubleMatcher = regexp.MustCompile(`(aa)|(bb)|(cc)|(dd)|(ee)|(ff)|(gg)|(hh)|(ii)|(jj)|(kk)|(ll)|(mm)|(nn)|(oo)|(pp)|(qq)|(rr)|(ss)|(tt)|(uu)|(vv)|(ww)|(xx)|(yy)|(zz)`)

func checkPassword(password string) bool {
	if illegalMatcher.MatchString(password) {
		return false
	}
	if !straight.MatchString(password) {
		return false
	}
	result := make(map[string]bool)
	matches := doubleMatcher.FindAllString(password, -1)
	for _, str := range matches {
		result[str] = true
	}
	if len(result) < 2 {
		return false
	}
	return true
}

func incrementPassword(password string) string {
	for i := len(password) - 1; i >= 0; i-- {
		letter := password[i]
		newLetter := letter + 1
		if string(letter) == "z" {
			newLetter = ([]byte("a"))[0]
			password = strings.Join([]string{password[:i], password[i+1:]}, string(newLetter))
		} else {
			password = strings.Join([]string{password[:i], password[i+1:]}, string(newLetter))
			return password
		}
	}
	return password
}

func main() {
	pass := "cqjxxyzz"
	for {
		pass = incrementPassword(pass)
		if checkPassword(pass) {
			fmt.Println(pass)
			break
		}
	}
}
