package main

import "fmt"

// 在为一个类型定义接收器方法时，返回其自身值

type Person struct {
	Age  int
	Name string
}

func (p *Person) AddAge() *Person {
	p.Age++
	return p
}

func (p *Person) Rename(name string) *Person {
	p.Name = name
	return p
}

func main() {
	p := Person{
		Age:  10,
		Name: "fishqiudao",
	}
	p.AddAge().Rename("cutfish")

	fmt.Println(p)
}
