package main

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
)

/*
go get github.com/hashicorp/golang-lru
go get github.com/hashicorp/golang-lru@0.5.1
*/

func main() {

	l, _ := lru.New(255)

	for i := 0; i < 300; i++ {
		l.Add(i, nil)
	}
	fmt.Println("----> ", l.Len())
	fmt.Println(l.GetOldest())
}
