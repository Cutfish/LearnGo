package main

import (
	"time"
)

func main() {
	go func(i int) {

	}(2)
	time.Sleep(time.Second * 2)
}
