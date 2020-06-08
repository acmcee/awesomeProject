package main

import "fmt"

func ChangeSlice1(s []string) {
	fmt.Println("from ChangeSlice1 len:", len(s), ", cap:", cap(s))
	s[1] = "modifyChangeSlice1" //  虽然是引用类型，发生了拷贝，但是指向的底层数组是同一个。
	fmt.Println("ChangeSlice1 point1:", &s)
	s = append(s, "AddChangeSlice1 ") // 对s 的引用重新赋值，因此修改对实参不生效
	fmt.Println("ChangeSlice1 point2:", &s)

}

func ChangeSlice2(s []string) {
	fmt.Println("from ChangeSlice2 len:", len(s), ", cap:", cap(s))
	s[1] = "modifyChangeSlice2" //  虽然是引用类型，发生了拷贝，但是指向的底层数组是同一个。
	fmt.Println("ChangeSlice2 point1:", &s)
	sp := &s
	*sp = append(*sp, "AddChangeSlice2 ") // 对s 的引用重新赋值，因此修改对实参不生效
	fmt.Println("ChangeSlice2 point2:", &s)
}

func ChangeSlice3(s *[]string) {
	fmt.Println("from ChangeSlice3 len:", len(*s), ", cap:", cap(*s))
	(*s)[1] = "modifyChangeSlice3"
	fmt.Println("ChangeSlice3 point1:", s)
	*s = append(*s, "AddChangeSlice3 ") // 对s 指针的实际内容进行修改
	fmt.Println("ChangeSlice3 point2:", s)
}

func ChangeSlice4(s *[]string) {
	fmt.Println("from ChangeSlice3 len:", len(*s), ", cap:", cap(*s))
	s = new([]string)
	fmt.Println("ChangeSlice4 point1:", s)
	*s = append(*s, "AddChangeSlice4 ") // 对s 的引用重新赋值，因此修改对实参不生效
	fmt.Println("ChangeSlice4 point2:", s)
}

func main() {

	s1 := make([]string, 2)
	s2 := make([]string, 2)
	s3 := make([]string, 2)
	s4 := make([]string, 2)
	ChangeSlice1(s1)
	ChangeSlice2(s2)
	ChangeSlice3(&s3)
	ChangeSlice4(&s4)
	fmt.Println("main s1", s1) //main s1 [ modifyChangeSlice1]

	fmt.Println("main s2", s2) // main s2 [ modifyChangeSlice2]

	fmt.Println("main s3", s3) // main s3 [ modifyChangeSlice3 AddChangeSlice3 ]

	fmt.Println("main s4", s4) // main s4 [ ]

}
