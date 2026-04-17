package lesson_01_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func DemoGoroutineBasics() {
	fmt.Println("\n--- Goroutine 基础演示 ---")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine A:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine B:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	fmt.Println("Main goroutine waiting...")
	wg.Wait()
	fmt.Println("All goroutines done")
}

func DemoGoroutineScheduling() {
	fmt.Println("\n--- Goroutine 调度演示 ---")

	fmt.Printf("当前 GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("当前 Goroutine 数量: %d\n", runtime.NumGoroutine())

	var wg sync.WaitGroup
	const numWorkers = 3

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		workerID := i
		go func() {
			defer wg.Done()
			fmt.Printf("Worker %d 开始运行\n", workerID)
			for j := 0; j < 3; j++ {
				fmt.Printf("Worker %d 执行任务 %d\n", workerID, j)
				runtime.Gosched()
			}
			fmt.Printf("Worker %d 完成\n", workerID)
		}()
	}

	fmt.Printf("启动后 Goroutine 数量: %d\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("所有 Worker 完成")
}

func DemoGoroutineLeak() {
	fmt.Println("\n--- Goroutine 泄露演示 ---")
	fmt.Println("Goroutine leak example (commented out to avoid actual leak):")
	fmt.Println(`
	// 错误示例 - 会导致 Goroutine 泄露
	ch := make(chan int)
	go func() {
		val := <-ch
		fmt.Println("Received:", val)
	}()
	// Goroutine 永远阻塞在这里，造成泄露
	`)
}
