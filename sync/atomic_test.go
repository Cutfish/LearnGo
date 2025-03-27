package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// 测试WaitGroup
func TestWaitG1roup(t *testing.T) {
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
