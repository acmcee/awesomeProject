package main

import (
	"fmt"
	"strconv"
)

type stringMap = map[string]string

func main() {
	s := make(stringMap, 5)
	for i := 0; i < 10; i++ {
		s[strconv.Itoa(i)] = strconv.Itoa(i + 100)
	}
	fmt.Println(s)
}
