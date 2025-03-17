package test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoruntine(t *testing.T) {
	ch := make(chan bool)
	go test1(ch)
	go test2(ch)

	time.Sleep(5 * time.Second)
	fmt.Println("over.......")
}

func test1(ch chan bool) {
	defer close(ch)

	time.Sleep(2 * time.Second)
	fmt.Println("test1 is over now send true to ch")
	ch <- true
}

func test2(ch chan bool) {
loop:
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("v is", v, "now break")
				break loop
			}
		case <-time.After(1 * time.Second):
			fmt.Println("time out ......")

		}
	}
}
