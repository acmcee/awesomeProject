package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Harvey struct {
	name string
}

func (h *Harvey) Write(b []byte) (int, error) {
	return 0, nil
}

func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

func main() {
	var w io.Writer
	fmt.Printf("%T\n", w) // "<nil>"
	// 调用了一个具体类型到接口类型的隐式转换
	w = os.Stdout         // 如果没有这一行，直接调用下面的write， 将panic，因为 w 的接口对象是nil
	fmt.Printf("%T\n", w) // "*os.File"
	w.Write([]byte("a"))
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) // "*bytes.Buffer"

	// 注意接口有接口类型和接口的值，  对于一个接口的零值就是它的类型和值的部分都是nil

	var buf1 io.Writer       // 此时 声明了一个buf1 为io.Writer接口，但是接口类型和接口的值都是空的
	f(buf1)                  //  f 判断 为nil,不输出
	buf1 = new(bytes.Buffer) // enable collection of output
	f(buf1)                  // OK

	var buf2 *bytes.Buffer   // 此时声明了一个buf2，对应的类型是 bytes.Buffer, 但是还没有赋值，所以直接调用它的write 将失败
	f(buf2)                  //  ERROR!  f 判断 不为nil, 进行输出的时候， ERROR ，接口的值为nill
	buf2 = new(bytes.Buffer) // enable collection of output
	f(buf2)                  //  OK
	fmt.Println(buf2)

}
