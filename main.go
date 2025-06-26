package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	ret := ""
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			ret += "number"
		} else {
			ret += string(s[i])
		}
	}
}

func isDigit(b byte) bool {

	if b <= '9' && b >= '0' {
		return true
	}
	return false
}
