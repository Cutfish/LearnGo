package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Produced:", i)
		time.Sleep((100 * time.Millisecond))
	}
}

func consumer(ch chan int, done chan bool) {
	for num := range ch {
		fmt.Println("Consumed:", num)
		time.Sleep((200 * time.Millisecond))
	}
	done <- true
}

func main() {
	ch := make(chan int, 5)
	done := make(chan bool)

	for i := 0; i < 3; i++ {
		go consumer(ch, done)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go producer(ch, &wg)
	}
	wg.Wait()
	close(ch)

	for i := 0; i < 3; i++ {
		<-done
	}
	fmt.Println("over.........")
}
