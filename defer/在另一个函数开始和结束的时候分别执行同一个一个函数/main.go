package main

import "fmt"

// 在另一个函数开始和结束的时候分别执行同一个一个函数
func test() {
	defer multiStageDefer()()
	fmt.Println("test func called")
}

func multiStageDefer() func() {
	fmt.Println("Run initialization")

	return func() {
		fmt.Println("Run clearnUp")
	}
}

func main() {
	test()
}
