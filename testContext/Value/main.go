package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.WithValue(context.Background(), "aaa", "aaa")

	fmt.Println(ctx.Value("aaa"))

}
