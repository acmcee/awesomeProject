package main

import (
	"fmt"
	"sort"
)

func main() {

	a := sort.StringSlice([]string{"d", "a", "b", "c"})
	sort.Strings(a)
	fmt.Println(a)
	sort.Sort(sort.Reverse(a))
	fmt.Println(a)

}
