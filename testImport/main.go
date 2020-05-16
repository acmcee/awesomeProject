package main

/*
在import 里面，import 是调用调用的常量，然后变量，最后是init，如果有依赖其他的包，则先导入其他的包


 */

import (
	"fmt"
	"awesomeProject/testImport/show"
	_ "awesomeProject/testImport/no_learn" // 这个导入只是调用对应的init 函数
)

func init()  {
	fmt.Println("我是main里面的init")
}

func main() {
	show.Show()
}
