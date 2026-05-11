package lesson_01_goroutine_csp

import (
	"fmt"
	"sync"
)

type cspTask struct {
	ID  int
	Job string
}

// DemoCSPModel 演示 CSP 模型核心思想
func DemoCSPModel() {
	fmt.Println("\n--- Section 3: CSP 模型核心思想 ---")

	fmt.Println("\n1. 对比：共享内存 + 锁 vs Channel 通信")
	fmt.Println("   Java 风格（共享内存 + 锁）:")
	demoJavaStyle()

	fmt.Println("\n   Go 风格（Channel 通信）:")
	demoGoStyle()

	fmt.Println("\n2. CSP 模式的实际应用：任务队列")
	demoTaskQueue()
}

// demoJavaStyle 演示 Java 风格的共享内存 + 锁方式
func demoJavaStyle() {
	var mu sync.Mutex
	count := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("   计数结果（锁保护）: %d\n", count)
}

// demoGoStyle 演示 Go 风格的 Channel 通信方式
func demoGoStyle() {
	increment := make(chan struct{})
	result := make(chan int)

	// 启动计数 Goroutine，状态完全隔离
	go func() {
		count := 0 // 只有这个 Goroutine 能访问 count
		for range increment {
			count++
		}
		result <- count
	}()

	// 发送 1000 次增量信号
	for i := 0; i < 1000; i++ {
		increment <- struct{}{}
	}
	close(increment)

	fmt.Printf("   计数结果（Channel 通信）: %d\n", <-result)
}

// demoTaskQueue 演示 CSP 模式的任务队列
func demoTaskQueue() {
	tasks := make(chan cspTask, 10)
	results := make(chan string, 10)
	var wg sync.WaitGroup

	// 启动 3 个 Worker Goroutine
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go cspWorker(i, tasks, results, &wg)
	}

	// 发送任务
	go func() {
		for i := 1; i <= 5; i++ {
			tasks <- cspTask{ID: i, Job: fmt.Sprintf("job-%d", i)}
		}
		close(tasks)
	}()

	// 收集结果
	go func() {
		wg.Wait()
		close(results)
	}()

	// 输出结果
	for result := range results {
		fmt.Printf("   结果: %s\n", result)
	}
}

func cspWorker(id int, tasks <-chan cspTask, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		results <- fmt.Sprintf("Worker %d 完成了 %s", id, task.Job)
	}
}
