package main

import "fmt"

/*
chan 是不同线程之间通信的比较占资源，用完需要立马关掉
len: string, array, slice, map, chan
cap: array, slice,chan
close: chan
 */

func main() {


	// Test get len
	// getLen()


	// Test close chan
	closeChan()
	
}

func getLen()  {
	mSlice:=make([] string, 2, 3)
	mSlice[0] = "id 1"
	mSlice[1] = "id 2"
	fmt.Println("before append")
	fmt.Println("len:", len(mSlice), "cap: ", cap(mSlice))

	mSlice= append(mSlice, "id 3")
	mSlice= append(mSlice, "id 4")

	fmt.Println("after append")
	fmt.Println("len:", len(mSlice), "cap: ", cap(mSlice))  // 可以发现cap 动态扩容了,底层数组发生了拷贝哦


}

func closeChan()  {
	mChan:=make(chan int, 3)
	defer close(mChan) // 不调用close 程序也不会报错， 保证只写需要的数据，最后要关闭chan
	mChan <- 1
	mChan <- 2
}