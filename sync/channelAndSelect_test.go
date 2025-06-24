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

func TestSelectAndTick(t *testing.T) {
	c := make(chan int, 1)
	tick := time.Tick(time.Second)
	for {
		select {
		case <-c:
			fmt.Println("random 01")
		case <-tick:
			fmt.Println("tick")
		case <-time.After(400 * time.Millisecond):
			fmt.Println("timeout")
		}
	}
}

// nil的channel读写都会阻塞
func TestNilChannel(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	go func() {
		for i := 0; i < 2; i++ {
			select {
			case a <- 1:
				a = nil
			case b <- 2:
				b = nil
			}
		}
	}()
	fmt.Println(<-a)
	fmt.Println(<-b)
}

// 对一个关闭的channel进行读操作，永远不会阻塞，如果缓冲区里没有数值了，那就返回对应的零值
func TestClosedChannel(t *testing.T) {
	a := make(chan int, 10)
	b := make(chan int, 10)
	close(a)
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-a:
				fmt.Println("a")
			case b <- 1:
				fmt.Println("b")
			default:
				fmt.Println("default")
			}
		}
	}()
	fmt.Println(len(a))
	fmt.Println(len(b))
}
