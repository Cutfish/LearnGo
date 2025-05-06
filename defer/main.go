package main

import (
	"fmt"
	"time"
)

func example() {
	defer func(start time.Time) {
		fmt.Println("耗时:", time.Since(start))
	}(time.Now())
	// 函数逻辑...
	for i := 0; i < 100000; i++ {
		for i := 0; i < 100000; i++ {

		}
	}
}

func main() {
	example()
}
