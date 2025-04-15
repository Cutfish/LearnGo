package test

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

// 测试sync.Pool
type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

func TestSyncPool(t *testing.T) {
	var StudentPool = sync.Pool{
		New: func() interface{} {
			return new(Student)
		},
	}
	stu := StudentPool.Get().(*Student)
	fmt.Println("stu: ", stu)
	json.Unmarshal(buf, stu)
	fmt.Println("buf: ", buf)
	fmt.Println("stu: ", stu)
	StudentPool.Put(stu)
}

// 测试sync.Cond
var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "starts reading")
	time.Sleep(time.Second * 2)
	c.L.Unlock()
}
func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
}
func TestSyncCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 6)
}

// 测试sync.Once
func TestSyncOnce(t *testing.T) {
	once := sync.Once{}
	for i := 0; i < 1000; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("sasasaassas")
			}) // 只会执行一次
		}()
	}
}
