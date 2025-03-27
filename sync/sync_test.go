package test

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

// 测试Lock和RWLock
func TestLock1(t *testing.T) {
	var lock sync.RWMutex
	GOTIME := "2006-01-02 15:04:05"
	//read lock
	fmt.Println("start read lock at", time.Now().Format(GOTIME))
	for i := range 5 {
		go func(i int) {
			defer lock.RUnlock()
			lock.RLock()
			fmt.Println("func", i, "get read lock at", time.Now().Format(GOTIME))
			time.Sleep(time.Second)
			fmt.Println("func", i, "release read lock at", time.Now().Format(GOTIME))
		}(i)
	}

	time.Sleep(time.Second / 10)

	//write lock
	fmt.Println("start write lock at", time.Now().Format(GOTIME))
	for i := range 5 {
		go func(i int) {
			defer lock.Unlock()
			lock.Lock()
			fmt.Println("func", i, "get write lock at", time.Now().Format(GOTIME))
			time.Sleep(time.Second)
			fmt.Println("func", i, "release write lock at", time.Now().Format(GOTIME))
		}(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("main func over...........")

}

// 测试WaitGroup
func TestWaitGroup1(t *testing.T) {
	var wg sync.WaitGroup
	GOTIME := "2006-01-02 15:04:05"
	//read lock
	fmt.Println("start read lock at", time.Now().Format(GOTIME))
	for i := range 5 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("func", i, "get read lock at", time.Now().Format(GOTIME))
			time.Sleep(time.Second)
			fmt.Println("func", i, "release read lock at", time.Now().Format(GOTIME))
		}(i)
	}

	wg.Wait()
	fmt.Println("main func over...........")
}

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
