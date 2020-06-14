package main

import "fmt"

func PipLine2() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}

		close(squares)
	}()

	// Printer (in main goroutine)
	for {
		x, ok := <-squares
		if !ok {

			fmt.Println("done")
			break
		}
		fmt.Println(x)
	}
}

func PipLine1() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {

		for x := range naturals {

			squares <- x * x
		}

		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
	fmt.Println("done")
}
func main() {
	PipLine1()
	PipLine2()
}
