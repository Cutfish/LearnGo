package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 必须设置随机种子（否则每次运行结果相同）
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// 生成随机整数 [0, 100)
	n := rand.Intn(100)
	fmt.Println("随机整数:", n)

	// 生成随机浮点数 [0.0, 1.0)
	f := rand.Float64()
	fmt.Println("随机浮点数:", f)

	fmt.Println(time.Now())
	t := time.Now()

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Since(t))

}
