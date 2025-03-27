package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "hello world!"
	reflectType := reflect.TypeOf(str)
	fmt.Println(reflectType)
}
