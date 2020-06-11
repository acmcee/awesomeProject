package main

import (
	"fmt"
)

type A struct {
	Name string
	BS   []*B
}

type B struct {
	Name string
}

func modifyB(b *B) {
	if b != nil {
		b.Name = "bbbbbbb"
	}
}

func main() {

	NewBS := make([]*B, 2)
	NewBS = append(NewBS, &B{Name: "aaa"})
	for _, x := range NewBS {
		modifyB(x)
	}
	fmt.Println(*NewBS[2])
}
