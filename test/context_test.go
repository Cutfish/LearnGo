package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserInfo struct {
	Name string
}

func GetUser(ctx context.Context) {
	fmt.Println(ctx.Value("name").(UserInfo).Name)
}

// 测试context
func TestContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", UserInfo{Name: "hahsha"})
	GetUser(ctx)
}

func GetIp(ctx context.Context) (ip string, err error) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("协程取消", ctx.Err())
			err = ctx.Err()
			wait.Done()
			return
		}
	}()
	time.Sleep(4 * time.Second)
	ip = "192.168.0.1"
	wait.Done()
	return
}

// 测试取消协程
var wait = sync.WaitGroup{}

func TestCancelGoroutine(t *testing.T) {
	t1 := time.Now()

	ctx, cancel := context.WithCancel(context.Background())

	wait.Add(1)
	go func() {
		ip, err := GetIp(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ip, err)
		wait.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()
	wait.Wait()
	fmt.Println("执行完成", time.Since(t1))
}

func GetIp1(ctx context.Context, wg *sync.WaitGroup) (ip string, err error) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("协程取消", ctx.Err())
			err = ctx.Err()
			wg.Done()
			return
		}
	}()
	time.Sleep(4 * time.Second)
	ip = "192.168.0.1"
	wg.Done()
	return
}

// 测试截止时间
func TestStopTime(t *testing.T) {
	var wg = sync.WaitGroup{}

	t1 := time.Now()

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))

	wg.Add(1)
	go func() {
		ip, err := GetIp1(ctx, &wg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ip, err)
		wg.Done()
	}()
	// cancel()
	wg.Wait()
	fmt.Println("执行完成", time.Since(t1))
}

func GetIp2(ctx context.Context, wg *sync.WaitGroup) (ip string, err error) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("协程取消", ctx.Err())
			err = ctx.Err()
			wg.Done()
			return
		}
	}()
	time.Sleep(4 * time.Second)
	ip = "192.168.0.1"
	wg.Done()
	return
}

// 测试超时时间
func TestTimeOut(t *testing.T) {
	var wg = sync.WaitGroup{}

	t1 := time.Now()

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	wg.Add(1)
	go func() {
		ip, err := GetIp2(ctx, &wg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ip, err)
		wg.Done()
	}()
	// cancel()
	wg.Wait()
	fmt.Println("执行完成", time.Since(t1))
}
