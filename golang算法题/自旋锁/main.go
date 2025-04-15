package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type SpinLock struct {
	locked int32
}

func (s *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32(&s.locked, 0, 1) {
		// 随机等待一点时间避免过多CPU占用
		time.Sleep(time.Nanosecond)
	}
}

func (s *SpinLock) UnLock() {
	atomic.StoreInt32(&s.locked, 0)
}

func main() {
	var lock SpinLock
	go func() {
		lock.Lock()
		lock.Lock()

		lock.UnLock()
		lock.UnLock()
	}()
	// 等待足够的时间以确保goroutines完成
	time.Sleep(time.Second)
	fmt.Println("over........")
}
