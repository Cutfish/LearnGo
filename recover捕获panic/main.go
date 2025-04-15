package main

import "fmt"

func main() {
	fmt.Println("程序开始运行")
	safeRun() // 调用包含 panic 的函数
	fmt.Println("程序安全退出")
}

func safeRun() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到 panic：", r)
		}
	}()

	fmt.Println("开始执行可能出错的逻辑")
	mayPanic() // 这里会触发 panic
	fmt.Println("这一行不会被执行，因为上面 panic 了")
}

func mayPanic() {
	panic("哎呀，出错了！") // 人为制造一个 panic
}
