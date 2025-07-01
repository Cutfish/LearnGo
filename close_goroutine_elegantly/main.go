package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(ctx context.Context, workerID int) {
	fmt.Printf("协程 %d: 启动并开始工作...\n", workerID)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("协程 %d: 收到停止信号，正在优雅退出...\n", workerID)
			// 在这里可以执行一些清理工作，比如关闭文件句柄、保存数据等
			fmt.Printf("协程 %d: 清理完成，退出。\n", workerID)
			return
		default:
			// 模拟协程正在执行的耗时操作
			fmt.Printf("协程 %d: 正在工作...\n", workerID)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	fmt.Println("主程序: 正在启动...")

	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT)

	// 可以启动多个协程
	go worker(ctx, 1)
	go worker(ctx, 2)

	sig := <-sigChan
	fmt.Printf("\n主程序: 收到信号 [%s]，正在通知协程退出...\n", sig.String())

	cancel()

	// 给协程一些时间来完成清理工作并退出
	// 实际应用中，你可能需要一个更健壮的机制来确保所有协程都已退出
	fmt.Println("主程序: 等待协程退出...")
	time.Sleep(3 * time.Second) // 给予3秒等待协程退出

	fmt.Println("主程序: 所有协程已退出，程序结束。")
}
