package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 测试Lock和RWLock
func TestLock1(t *testing.T) {
	var lock sync.RWMutex
	GOTIME := "2006-01-02 15:04:05"
	//read lock
	fmt.Println("start read lock at", time.Now().Format(GOTIME))
	for i := range 5 {
		go func(i int) {
			defer lock.RUnlock()
			lock.RLock()
			fmt.Println("func", i, "get read lock at", time.Now().Format(GOTIME))
			time.Sleep(time.Second)
			fmt.Println("func", i, "release read lock at", time.Now().Format(GOTIME))
		}(i)
	}

	time.Sleep(time.Second / 10)

	//write lock
	fmt.Println("start write lock at", time.Now().Format(GOTIME))
	for i := range 5 {
		go func(i int) {
			defer lock.Unlock()
			lock.Lock()
			fmt.Println("func", i, "get write lock at", time.Now().Format(GOTIME))
			time.Sleep(time.Second)
			fmt.Println("func", i, "release write lock at", time.Now().Format(GOTIME))
		}(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("main func over...........")

}

// 测试WaitGroup
func TestWaitGroup1(t *testing.T) {
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
