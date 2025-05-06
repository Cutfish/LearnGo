package main

import (
	"fmt"
	"sync"
)

// 构造协程池类
type ITask interface {
	Run()
}

type taskStruct struct {
	id int
}

func (t taskStruct) Run() {
	fmt.Println(t.id)
}

type IGoPool interface {
	Start()
	Schedule(task ITask)
	WaitAndStop()
}

type gPool struct {
	workers int
	tasks   chan ITask
	wg      sync.WaitGroup
}

func NewPool(workers int) IGoPool {
	return &gPool{
		workers: workers,
		tasks:   make(chan ITask, workers),
		wg:      sync.WaitGroup{},
	}
}

func (p *gPool) Start() {
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

func (p *gPool) Schedule(task ITask) {
	p.tasks <- task
	fmt.Println("发送任务")
}

func (p *gPool) WaitAndStop() {
	close(p.tasks)
	p.wg.Wait()
}

func main() {
	p := NewPool(5)

	p.Start()
	for i := 0; i < 100; i++ {
		t := taskStruct{
			id: i,
		}
		p.Schedule(t)
	}
	p.WaitAndStop()
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
