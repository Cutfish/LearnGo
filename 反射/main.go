package main

import (
	"fmt"
	"reflect"
)

type MyInt int
type User struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func main() {
	u := User{
		Name: "hanxuegang",
		Id:   12345,
	}
	t := reflect.TypeOf(u)
	if temp, ok := t.FieldByName("Id"); ok {
		fmt.Printf("%v %v %v %v %v\n", temp.Type.Name(), temp.Name, temp.Index, temp.Type.Kind(), temp.Tag.Get("json"))
	}
	fmt.Println(t.Name(), t.Kind()) // student struct
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%v %v %v %v %v\n", f.Type, f.Name, f.Index, f.Type.Kind(), f.Tag.Get("json"))
	}
}
