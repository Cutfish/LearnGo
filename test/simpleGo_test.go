package test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Stu struct {
	Name string
}

// 定义 String() 方法
//
//	func (s Stu) String() string {
//		return fmt.Sprintf("Student Name: %s", s.Name)
//	}

// 测试结构体输出的情况
func TestStuPrint(t *testing.T) {
	fmt.Printf("%v\n", Stu{"Tom"})  // Student Name: Tom
	fmt.Printf("%+v\n", Stu{"Tom"}) // Student Name: Tom
}

// 判断两个字符串是否相等
func IntSliceEqualBCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// 测试字符串是否相等
func TestEqualSlice(t *testing.T) {
	var list1 []int
	list2 := []int{}
	if list1 == nil {
		fmt.Println("这是一个nil")
	}

	if IntSliceEqualBCE(list1, list2) {
		fmt.Println("一样的")
	} else {
		fmt.Println("不一样的")
	}
	if reflect.DeepEqual(list1, list2) {
		fmt.Println("一样的")
	} else {
		fmt.Println("不一样的")
	}
}

// 测试空字符串
func TestEnums(t *testing.T) {
	fmt.Println(unsafe.Sizeof(struct{}{})) // 0
}

type Stu1 struct {
	Name string
}

func TestInterface(t *testing.T) {
	var stu1, stu2 interface{} = &Stu1{"Tom"}, &Stu1{"Tom"}
	var stu3, stu4 interface{} = Stu1{"Tom"}, Stu1{"Tom"}
	fmt.Println(stu1 == stu2) // false
	fmt.Println(stu3 == stu4) // true

	// 两个接口值比较时，会先比较 T，再比较 V。
	// 接口值与非接口值比较时，会先将非接口值尝试转换为接口值，再比较。
	// 一个接口等于 nil，当且仅当 T 和 V 处于 unset(nil) 状态
	var p *int = nil
	var i interface{} = p
	fmt.Println(i == p)   // true
	fmt.Println(p == nil) // true
	fmt.Println(i == nil) // false
}
