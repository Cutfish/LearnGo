package main

import (
	"fmt"
	"sync"
)

func main() {
	a := 1
	b := 1
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 3; i <= 10; i += 3 {
			<-ch1
			fmt.Println(a + b)
			temp := b
			b = a + b
			a = temp
			ch2 <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 4; i <= 10; i += 3 {
			<-ch2
			fmt.Println(a + b)
			temp := b
			b = a + b
			a = temp
			if i < 10 {
				ch3 <- struct{}{}
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 5; i <= 10; i += 3 {
			<-ch3
			fmt.Println(a + b)
			temp := b
			b = a + b
			a = temp
			ch1 <- struct{}{}
		}
	}()
	ch1 <- struct{}{}
	wg.Wait()
	fmt.Println("over...........")
}
