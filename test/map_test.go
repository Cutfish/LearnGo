package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 测试sync.Map
func TestMap(t *testing.T) {
	var hashmap sync.Map
	var wg sync.WaitGroup

	for i := range 100 {
		fmt.Println("goroutinue id is", i)

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			hashmap.Store(i, fmt.Sprintf("%daaaa", i))
		}(i)
	}

	wg.Wait()
	value, ok := hashmap.Load(1)
	fmt.Println("value is", value, "ok is", ok)

	ac, o12k := hashmap.LoadOrStore(2, "1212121")
	ac1, ok21 := hashmap.LoadOrStore("sasas", 2)
	ac2, o21k := hashmap.LoadOrStore(true, "1212121")
	a3, o1k := hashmap.LoadOrStore(time.Second, "1212121")
	fmt.Println(ac, ac1, ac2, a3, o12k, o21k, o1k, ok21)
	fmt.Println("main func over...........")
}
