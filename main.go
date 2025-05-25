package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func splitSlice(s []int) [][]int {
	len_s := len(s)
	var ret [][]int
	if len_s%3 == 0 {
		ret = make([][]int, len_s/3)
	} else {
		ret = make([][]int, len_s/3+1)
	}

	count := 0
	for i := 0; i < len_s; i++ {
		if i%3 == 0 {
			count++
		}
		ret[count] = append(ret[count], s[i])
	}
	return ret
}

func MaskRealName(realName string) string {
	slice := []rune(realName)
	len := len(slice)
	if len == 1 {
		return string(slice)
	}
	if len == 2 {
		slice[1] = rune('*')
		return string(slice)
	}
	for i := 0; i < len; i++ {
		if i > 0 && i < len-1 {
			slice[i] = rune('*')
		}
	}
	return string(slice)
}

func main() {
	var sum atomic.Int32
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a := sum.Add(1)
			fmt.Println(a)
		}()
	}
	wg.Wait()
	fn1()
	// 这里期望能打印10000，但有bug,请改造main函数代码（任意方案均可）
	fmt.Println(sum)
}

// todo 请写splitSlice函数，返回二维切片

func fn1() int {
	return 1
}
