package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello # world"
	subStr := "he"
	fmt.Println("strings.Contains", str, "包含", subStr, "么？", strings.Contains(str, subStr))
	fmt.Println("strings.Index", strings.Index(str, subStr))
	fmt.Println("strings.Split", strings.Split(str, "#"))
	fmt.Println("strings.Join", strings.Join(strings.Split(str, "#"), ","))
	fmt.Println("strings.HasPrefix", strings.HasPrefix(str, "hello"))
	fmt.Println("strings.HasSuffix", strings.HasSuffix(str, "world"))
}
