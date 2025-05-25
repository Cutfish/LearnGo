package main

import (
	"fmt"
	"sync"
)

// 任务结构体
type Task struct {
	ID int
}

// Run 方法执行任务
func (t Task) Run() {
	fmt.Println("Task ID:", t.ID)
}

// 协程池结构体
type GoPool struct {
	workers int
	tasks   chan Task
	wg      sync.WaitGroup
}

// NewPool 创建一个新的协程池
func NewPool(workers int) *GoPool {
	return &GoPool{
		workers: workers,
		tasks:   make(chan Task, workers),
		wg:      sync.WaitGroup{},
	}
}

// Start 启动协程池中的 worker
func (p *GoPool) Start() {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for task := range p.tasks {
				task.Run()
			}
		}()
	}
}

// Schedule 将任务提交到协程池
func (p *GoPool) Schedule(task Task) {
	p.tasks <- task
	fmt.Println("发送任务:", task.ID)
}

// WaitAndStop 等待所有任务完成并停止协程池
func (p *GoPool) WaitAndStop() {
	close(p.tasks)
	p.wg.Wait()
}

func main() {
	pool := NewPool(5)
	pool.Start()

	for i := 0; i < 100; i++ {
		task := Task{ID: i}
		pool.Schedule(task)
	}

	pool.WaitAndStop()
	fmt.Println("所有任务执行完毕，协程池已停止。")
}

// 简单协程池
func SimpleCoPool() {
	numTasks := 100
	concurrent := 5

	wg := sync.WaitGroup{}
	taskCh := make(chan int, 5)

	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range taskCh {
				fmt.Println("process", i)
			}
		}()
	}

	for i := 0; i < numTasks; i++ {
		taskCh <- i
	}
	close(taskCh)

	wg.Wait()
	fmt.Println("over............")
}
