package lesson_01_goroutine_csp

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// DemoGoroutineScheduling 演示 Goroutine 调度与栈扩容
func DemoGoroutineScheduling() {
	fmt.Println("\n--- Section 2: Goroutine 调度与栈扩容 ---")

	fmt.Println("\n1. G-M-P 调度器基本信息：")
	// 查看当前 GOMAXPROCS 设置
	fmt.Printf("  GOMAXPROCS: %d (默认等于 CPU 核心数)\n", runtime.GOMAXPROCS(0))
	fmt.Printf("  CPU 核心数: %d\n", runtime.NumCPU())
	fmt.Printf("  当前 Goroutine 数量: %d\n", runtime.NumGoroutine())

	fmt.Println("\n2. 创建大量 Goroutine（演示轻量级特性）：")
	var wg sync.WaitGroup
	const numGoroutines = 1000

	fmt.Printf("  准备创建 %d 个 Goroutine...\n", numGoroutines)
	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 模拟一些轻量工作
			time.Sleep(1 * time.Millisecond)
		}(i)
	}

	fmt.Printf("  创建后 Goroutine 数量: %d\n", runtime.NumGoroutine())
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("  %d 个 Goroutine 执行完成，耗时: %v\n", numGoroutines, elapsed)
	fmt.Printf("  结束后 Goroutine 数量: %d\n", runtime.NumGoroutine())

	fmt.Println("\n3. 演示工作窃取（Work Stealing）：")
	demoWorkStealing()
}

func demoWorkStealing() {
	var wg sync.WaitGroup
	const numTasks = 20

	fmt.Println("  启动多个任务，观察调度器的负载均衡...")

	// 创建一些不同耗时的任务
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 不同任务有不同的执行时间
			workTime := time.Duration(id%3+1) * 10 * time.Millisecond
			time.Sleep(workTime)
		}(i)
	}

	wg.Wait()
	fmt.Println("  所有任务完成，调度器通过工作窃取实现了负载均衡")
}
