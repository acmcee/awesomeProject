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

	fmt.Println(strings.Compare("a", "ba"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))

	fmt.Println(strings.ContainsRune("aardvark", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))

	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
	fmt.Println(strings.EqualFold("Go", "go"))

	fmt.Println(new(interface{}))
}
