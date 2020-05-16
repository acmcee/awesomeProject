package main

import (
	"fmt"
	"reflect"
)

func main() {
	deleteFromMap()
}

func deleteFromMap()  {

	mMap:= make(map[int] string, 10)
	mMap[0] = "张三"
	mMap[100] = "李四"
	mMap[200] = "王五"
	fmt.Println("before delete", mMap)
	delete(mMap,100)
	fmt.Println("after delete", mMap)
}


func testCopy()  {
	srcSlice:=make([] string, 5)
	srcSlice[0] = "src-111"
	srcSlice[1] = "src-222"
	srcSlice[2] = "src-333"


	dstSlice:=make([] string, 2)
	dstSlice[0] = "dst-111"
	dstSlice[1] = "dst-222"

	copy(dstSlice, srcSlice)

	fmt.Println("dstSlice", dstSlice)
	fmt.Println("srcSlice", srcSlice)
}


func makeSlice() []string {
	mSlice:= make([]string, 2)
	fmt.Println(reflect.TypeOf(mSlice))
	mSlice[0] = "dog"
	mSlice[1] = "这种"
	fmt.Println(mSlice)
	return mSlice

}

func newMap()  {
	mMap:= new(map[int]string)
	mMakeMap:= make(map[int]string)
	fmt.Println(reflect.TypeOf(mMap))
	fmt.Println(reflect.TypeOf(mMakeMap))
}

func makeMap()  {
	mMap:= make(map[int]string)
	mMap[100] = "zhangsan"
	mMap[200] = "死锁"
	fmt.Println(mMap)
}

func makeChan()  {
	mChan := make(chan int, 3)
	close(mChan)

}