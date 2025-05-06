package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mu1 := sync.Mutex{}
	mu2 := sync.Mutex{}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu1.Lock()
		defer mu1.Unlock()
		fmt.Println("func1 got mu1!")

		time.Sleep(time.Second * 2)
		mu2.Lock()
		defer mu2.Unlock()
		fmt.Println("func1 got mu2!")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		mu2.Lock()
		defer mu2.Unlock()
		fmt.Println("func2 got mu2!")

		time.Sleep(time.Second * 2)
		mu1.Lock()
		defer mu1.Unlock()
		fmt.Println("func2 got mu1!")
	}()

	wg.Wait()
	fmt.Println("over..............")
}
