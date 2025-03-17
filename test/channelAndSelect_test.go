package test

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelANDSelect(t *testing.T) {
	GOTIME := "2006-01-02 15:04:05"
	ch := make(chan string, 5)
	go test3(ch, GOTIME)
outer:
	for {
		select {
		case value, ok := <-ch:
			if ok {
				fmt.Println("value is ", value, "ok is", ok)
			} else {
				fmt.Println("value is ", value, "ok is", ok)
				break outer
			}
		case <-time.After(2 * time.Second):
			fmt.Println("time out......")
			break outer
		}
	}
	fmt.Println(time.Now().Format(GOTIME), "over/.......")
}

func test3(ch chan string, gotime string) {

	defer func() {
		fmt.Println("i am test3 new close chan")
		close(ch)
	}()

	for i := 0; i < 5; i++ {
		v := time.Now().Format(gotime)
		fmt.Println("i is", i, "v is ", v)
		ch <- v
		time.Sleep(3 * time.Second)
	}
}
