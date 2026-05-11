package lesson_06_race_detector

import (
	"fmt"
	"sync"
)

// Run 运行 Race Detector 数据竞争检测实战课程代码
func Run() {
	fmt.Println("\n=== Lesson 6: Race Detector 数据竞争检测实战 ===")

	DemoDataRace()
	DemoFixedRace()
}

// DemoDataRace 演示数据竞争（实际项目中不要这样写！）
func DemoDataRace() {
	fmt.Println("\n--- 演示：数据竞争 ---")

	var count int
	var wg sync.WaitGroup

	// 这段代码有数据竞争！
	// 实际使用 go test -race 运行时会检测到
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++ // 数据竞争！多个 Goroutine 同时读写 count
		}()
	}

	wg.Wait()
	fmt.Printf("  （有数据竞争）计数结果: %d (每次运行可能不同)\n", count)
}

// DemoFixedRace 演示修复后的数据竞争
func DemoFixedRace() {
	fmt.Println("\n--- 演示：修复数据竞争 ---")

	var count int
	var mu sync.Mutex
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
	fmt.Printf("  （已修复）计数结果: %d (每次运行都应该是 1000)\n", count)
}
