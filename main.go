package main

import "fmt"

func main() {
	const a = 1
	fmt.Println(add1(a))
}

func isDigit(b byte) bool {

	if b <= '9' && b >= '0' {
		return true
	}
	return false
}

func add1(a int) int {
	return a + 1
}
