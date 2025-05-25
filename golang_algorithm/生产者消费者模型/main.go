package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Consumed:", num)
	}
}

func main() {
	ch := make(chan int, 5)

	wg1 := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg1.Add(1)
		go consumer(ch, &wg1)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go producer(ch, &wg)
	}
	wg.Wait()
	close(ch)
	fmt.Println("生产完毕！")

	wg1.Wait()

	fmt.Println("over.........")
}
