package main

import "fmt"

func TestLoop1() {
	var TestDo []func()
	strs := []string{"a", "b", "c"}
	for _, s := range strs {
		var news string // 重新声明
		news = s        // 进行赋值
		TestDo = append(TestDo, func() {
			fmt.Print(" ", news) // 是针对news 的引用
		})
	}
	fmt.Print("loop1:")

	for _, d := range TestDo {
		d() // loop1: a b c

	}
	fmt.Println()
}

func TestLoop2() {
	var TestDo []func()
	strs := []string{"a", "b", "c"}
	for _, s := range strs {
		news := s // 重新声明并且赋值
		TestDo = append(TestDo, func() {
			fmt.Print(" ", news)
		})
	}
	fmt.Print("loop2:")

	for _, d := range TestDo {
		d() //loop2: a b c

	}
	fmt.Println()
}

func TestLoop3() {
	var TestDo []func()
	strs := []string{"a", "b", "c"}
	var news string // 外面声明
	for _, s := range strs {
		news = s // 内部赋值，只是赋值
		TestDo = append(TestDo, func() {
			fmt.Print(" ", news) // 保存的是引用，外面的每次更改，都会影响内部
		})
	}
	fmt.Print("loop3:")
	for _, d := range TestDo {
		d() //loop3: c c c

	}
	fmt.Println()
}

func TestLoop4() {
	var TestDo []func()
	strs := []string{"a", "b", "c"}
	for _, s := range strs {
		TestDo = append(TestDo, func() {
			fmt.Print(" ", s) // 保存的是引用，外面的每次更改，都会影响内部
		})
	}
	fmt.Print("loop4:")
	for _, d := range TestDo {
		d() //loop4: c c c

	}
	fmt.Println()
}

func TestAdd() func() int {

	var i int
	return func() int {
		i++
		return i * i
	}

}

func main() {

	TestLoop1()
	TestLoop2()
	TestLoop3()
	TestLoop4()

	f := TestAdd()
	fmt.Println("TestAdd1:", f()) // TestAdd1: 1
	fmt.Println("TestAdd2:", f()) // TestAdd2: 4   闭包里面的修改会影响外面的值

}
