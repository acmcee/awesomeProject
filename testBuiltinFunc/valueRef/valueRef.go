package main

import "fmt"

func ChangeSlice(s []string) {
	fmt.Println("from change slice len:", len(s), ", cap:", cap(s))
	s[1] = "new slice 1"         //  虽然是引用类型，发生了拷贝，但是指向的底层数组是同一个。
	s = append(s, "new slice 2") // 对s 的引用重新赋值，因此修改对实参不生效
	fmt.Println("from change slice len:", len(s), ", cap:", cap(s))

}

func main() {

	s := make([]string, 2)
	s[0] = "aaaa"
	fmt.Println("from main slice len:", len(s), ", cap:", cap(s))
	ChangeSlice(s)

	fmt.Println(s)
}
