package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func getUnicodeText() {
	var i int
	var e int
	myF, err := os.Create("/Users/wanghangwei/Desktop/a.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer myF.Close()
	flag.IntVar(&i, "b", 0, "unicode 起始码点")
	flag.IntVar(&e, "e", 130000, "unicode 结束")
	flag.Parse()
	myWrite := bufio.NewWriter(myF)
	_, _ = myWrite.Write([]byte(fmt.Sprintf("# -> %d ", i)))

	for {
		_, err := myWrite.Write([]byte(fmt.Sprintf("%c ", i)))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		i++
		if i%50 == 0 {
			_, err := myWrite.Write([]byte("\n"))
			_, _ = myWrite.Write([]byte(fmt.Sprintf("# -> %d ", i)))

			myWrite.Flush()

			fmt.Printf("%d个了、、、\n", i)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(-1)
			}

		}
		if i > e {
			fmt.Println("结束啦")
			myWrite.Flush()
			os.Exit(0)
		}

	}
}

func BitCal() {

	/*
	   测试byte 使用， #  表示使用大写OX 输出， [1] 表示使用第一个操作数进行输出。
	*/
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o \n", o) // "438 666 0666"
	x := int64(0xdeaDDdbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[2]X\n", x, x)

}

func main() {
	BitCal()
}
