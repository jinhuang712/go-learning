package lesson_01_goroutine_csp

import (
	"fmt"
	"sync"
	"time"
)

// DemoWaitGroupBasics 演示 WaitGroup 基础用法
func DemoWaitGroupBasics() {
	fmt.Println("\n--- Section 4: WaitGroup 基础用法 ---")

	fmt.Println("\n1. 正确用法：标准 WaitGroup 模式")
	demoCorrectUsage()

	fmt.Println("\n2. 错误示例：在 Goroutine 内调用 Add")
	demoBadAddPosition()

	fmt.Println("\n3. 实际应用：并行处理任务")
	demoParallelProcessing()
}

func demoCorrectUsage() {
	var wg sync.WaitGroup
	tasks := []string{"task-A", "task-B", "task-C"}

	for _, task := range tasks {
		wg.Add(1) // 在启动 Goroutine 前调用 Add
		go func(t string) {
			defer wg.Done() // 使用 defer 确保 Done 被调用
			fmt.Printf("   开始处理: %s\n", t)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("   完成处理: %s\n", t)
		}(task)
	}

	fmt.Println("   等待所有任务完成...")
	wg.Wait()
	fmt.Println("   所有任务完成！")
}

func demoBadAddPosition() {
	fmt.Println("   ⚠️  错误示例：在 Goroutine 内部调用 Add 可能导致问题")
	fmt.Println("   （此演示不实际运行，避免非确定性行为）")
	fmt.Println("   正确做法：始终在启动 Goroutine 之前调用 wg.Add(1)")
}

func demoParallelProcessing() {
	var wg sync.WaitGroup
	const numWorkers = 3

	fmt.Printf("   启动 %d 个 Worker 处理数据...\n", numWorkers)
	start := time.Now()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go waitGroupWorker(i, &wg)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("   所有 Worker 完成，总耗时: %v\n", elapsed)
}

func waitGroupWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("   Worker %d: 开始工作\n", id)
	// 模拟不同的工作时间
	workTime := time.Duration(id) * 150 * time.Millisecond
	time.Sleep(workTime)
	fmt.Printf("   Worker %d: 完成工作 (耗时: %v)\n", id, workTime)
}
