package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(read io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := read.Read(p)
	if err != nil {
		fmt.Println("read error: ", err.Error())
		return p, err
	}
	return p[:n], nil
}

func ReadString() {
	str := "readfrom string"
	data, _ := ReadFrom(strings.NewReader(str), 15)
	fmt.Println(string(data))
}

func ReadStdin() {
	fmt.Println("please input string")
	data, err := ReadFrom(os.Stdin, 15)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))
}

func ReadFile() {
	fmt.Println("please read file name:")
	fileName, _ := ReadFrom(os.Stdin, 100)
	fd, err := os.Open(strings.Trim(string(fileName), "\n"))
	defer fd.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := ReadFrom(fd, 1000)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))
}

func BufferIOTest() {

	bufReader := bufio.NewReader(strings.NewReader("hello world"))
	data, _ := bufReader.ReadString(' ')

	fmt.Println(string(data))

	bufWriter := bufio.NewWriter(os.Stdout)
	bufWriter.Write([]byte("hello , 我是通过Buffer 打印的\n"))
	fmt.Fprint(bufWriter, "哈哈哈\n")
	bufWriter.Flush()
}

func CountFileLine() {
	if len(os.Args) < 2 {
		fmt.Println("参数个数小于2， 请输入文件名")
		return
	}
	fileName := os.Args[1]
	fd, _ := os.Open(fileName)
	bufReader := bufio.NewReader(fd)
	lines := 0
	for ; ; {
		data, isPrefix, err := bufReader.ReadLine()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(string(data))
		if !isPrefix {
			lines++
		}
	}
	fmt.Println(lines)

}

func main() {
	// 从string 创建阅读器
	//ReadString()

	// 从标准输入创建阅读器
	//ReadStdin()

	// 从文件读取
	//ReadFile()

	// buffer io test
	//BufferIOTest()

	//CountFileLine
	CountFileLine()
}
