package show

import (
	"fmt"
	"awesomeProject/testImport/learn"
)

func init()  {
	fmt.Println("我是show 里面的init")
}

func Show()  {
	fmt.Println("我想要展示")
	learn.Learn()
}