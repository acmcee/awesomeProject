package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type OnlyWrite struct {
}

func (w OnlyWrite) Write(b []byte) (int, error) {
	return 0, nil
}

func main() {

	var w1 io.Writer
	w1 = os.Stdout
	f, ok := w1.(*os.File) // 类型断言为具体的类型，可以直接调用该类型的方法
	fmt.Printf("f, %v, %v \n", f, ok)
	g, ok := w1.(*bytes.Buffer)
	fmt.Printf("g, %v, %v \n", g, ok) // false

	var w2 io.Writer
	w2 = os.Stdout
	rw, ok := w2.(io.ReadWriter) // 断言接口类型，进行转换，获取更多的方法
	fmt.Printf("rw, %v, %v \n", rw, ok)

	w2 = new(OnlyWrite)
	ow, ok := w2.(io.ReadWriter) // panic: *ByteCounter has no Read method
	fmt.Printf("ow, %v, %v \n", ow, ok)

}
