package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	for i := 1000; i < 1000000000; i++ {
		test := "yzbqklnj" + strconv.Itoa(i)
		data := []byte(test)
		hash := md5.Sum(data)
		s := hex.EncodeToString(hash[:])
		if s[:6] == "000000" {
			fmt.Println(i)
			fmt.Printf(s)
			break
		}
	}
}
