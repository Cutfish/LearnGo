package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect1(t *testing.T) {

	str := "hello world!"
	reflectType := reflect.TypeOf(str)
	fmt.Printf("the type of str is %T\n", str)
	fmt.Println("the reflecttype of str is", reflectType)

	// 声明一个any类型的变量
	var eface any
	// 赋值
	eface = 100
	// 通过Kind方法，来获取其类型
	fmt.Printf("the type of eface is %T\n", eface)
	fmt.Println(reflect.TypeOf(eface).Kind())
	fmt.Println(reflect.TypeOf(eface))
}

func TestReflect2(t *testing.T) {
	var eface any
	eface = map[string]int{}
	rType := reflect.TypeOf(eface)
	// key()会返回map的键反射类型
	fmt.Println(rType.Key())
	fmt.Println(rType.Elem().Kind())

	// 获取结构体反射类型
	rType := reflect.TypeOf([]int{}).Elem()
	// 输出方法个数
	fmt.Println(rType.NumMethod())
	// 遍历输出方法信息
	for i := 0; i < rType.NumMethod(); i++ {
		method := rType.Method(i)
		fmt.Println(method.Index, method.Name, method.Type, method.IsExported())
	}
}
