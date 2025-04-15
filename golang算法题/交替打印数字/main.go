package main

import (
	"fmt"
	"sync"
)

func main() {

}

// 方法1(使用channel和wg)
func func1() {
	ch0, ch1 := make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i += 2 {
			<-ch0
			fmt.Printf("第一个：%d ", i)
			ch1 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 11; i += 2 {
			<-ch1
			fmt.Printf("第二个：%d ", i)
			if i < 11 {
				ch0 <- struct{}{}
			}
		}
	}()
	ch0 <- struct{}{}
	wg.Wait()
	fmt.Println("over...........")
}

//方法二：使用select、Mutex
