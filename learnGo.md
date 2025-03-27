### http
任何实现了 ServeHTTP 方法的对象都可以作为 HTTP 的 Handler。
```Go
type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("hello World!"))
}

func main() {
	var s server
	http.ListenAndServe("localhost:9999", &s)
}
```

**panic和error**
panic适用于不能恢复的错误（recover可以截取），error是可以恢复的错误，一个是完全错误，一个业务错误

**type...和...type**
...type表示可变参数，type...就是对切片展开
```Go
func (m *Map) Add(keys ...string) {
	fmt.Println(keys) // keys 是一个 []string 切片
}
func main() {
	var myMap Map
	myMap.Add("apple", "banana", "cherry")
	// 输出: ["apple" "banana" "cherry"]

    fruits := []string{"apple", "banana", "cherry"}
    myMap.Add(fruits...) // 需要使用 "..." 展开切片 不展开就报错

}
```
**Context**
`ctx, cancel := context.WithCancel(context.Background())`
context.Backgroud() 创建根 Context，通常在 main 函数、初始化和测试代码中创建，作为顶层 Context。
context.WithCancel(parent) 创建可取消的子 Context，同时返回函数 cancel。


**defer与return**
return注册返回值-->defer-->return退出函数
defer是否能修改return的值关键就在于，返回值是否是在函数声明的时候就注册好的或者是不是引用传递（如果函数声明的时候没有说明具体返回谁，那此时的return i可以看作是一个局部变量，后续的defer修改不起作用）

**sync.Map**
与普通的Map相比，不仅实现了并发读写，还是实现了读写分离，性能一般更高

**reflect**
fmt中获取类型还是值其实底层都是反射：
```Golang
switch verb {
case 'T':
	p.fmt.fmtS(reflect.TypeOf(arg).String())
	return
case 'p':
	p.fmtPointer(reflect.ValueOf(arg), 'p')
	return
}
```
