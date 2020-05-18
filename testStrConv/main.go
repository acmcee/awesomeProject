package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.Itoa(111111))
	fmt.Println(strconv.Atoi("1111112222a"))
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseFloat("3.14", 64))

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(64, 2))
	a := []byte("123")
	fmt.Printf("%s \n", strconv.AppendInt(a, 1, 10))

}
