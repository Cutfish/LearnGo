package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// 测试原子递增
func TestAtomicAdd(t *testing.T) {
	var count int32 = 0
	var wg sync.WaitGroup

	// 并发 1000 个 goroutine 递增 count
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1) // 原子递增
		}()
	}

	wg.Wait()
	fmt.Println("最终 count =", count) // 预期输出 1000
}

func TestAtomicNum(t *testing.T) {
	var a atomic.Int32
	a.Store(1)
	fmt.Println(a.Load())
	b := a.Load()
	fmt.Println(b)
}
