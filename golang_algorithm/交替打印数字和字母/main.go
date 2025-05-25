package main

import (
	"fmt"
	"sync"
)

func main() {
	numCh, letterCh := make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 打印数字
	go func() {
		defer wg.Done()
		i := 1
		for i <= 28 {
			<-numCh
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			if i <= 28 {
				letterCh <- struct{}{}
			}
		}
	}()
	go func() {
		defer wg.Done()
		c := 'A'
		for c <= 'Z' {
			<-letterCh
			fmt.Print(string(c))
			c++
			fmt.Print(string(c))
			c++
			numCh <- struct{}{}
		}

	}()
	numCh <- struct{}{}
	wg.Wait()
	fmt.Println("over............")
}
