package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 测试WaitGroup
func TestWaitG1roup(t *testing.T) {
	var wg sync.WaitGroup
	GOTIME := "2006-01-02 15:04:05"
	//read lock
	fmt.Println("start read lock at", time.Now().Format(GOTIME))
	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("func", i, "get read lock at", time.Now().Format(GOTIME))
			time.Sleep(time.Second)
			fmt.Println("func", i, "release read lock at", time.Now().Format(GOTIME))
		}(i)
	}

	wg.Wait()
	fmt.Println("main func over...........")
}
