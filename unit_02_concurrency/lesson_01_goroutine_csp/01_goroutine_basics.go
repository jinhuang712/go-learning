package lesson_01_goroutine_csp

import (
	"fmt"
	"time"
)

// DemoGoroutineBasics 演示 Goroutine 基础用法
func DemoGoroutineBasics() {
	fmt.Println("\n--- Section 1: Goroutine 基础入门 ---")

	fmt.Println("\n1. 简单的 Goroutine 示例：")
	// 启动一个 Goroutine
	go sayHello()

	// 给 Goroutine 一点时间执行（演示用，生产环境请用 WaitGroup）
	time.Sleep(50 * time.Millisecond)

	fmt.Println("\n2. 多个 Goroutine 并发执行：")
	// 启动多个 Goroutine
	go printNumbers("Goroutine A")
	go printNumbers("Goroutine B")
	go printNumbers("Goroutine C")

	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n3. 匿名函数 Goroutine：")
	// 使用匿名函数创建 Goroutine
	message := "Hello from anonymous Goroutine"
	go func(msg string) {
		fmt.Println(msg)
	}(message) // 注意：参数需要显式传递，避免闭包捕获变量的问题

	time.Sleep(50 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello from a Goroutine!")
}

func printNumbers(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		// 短暂休眠，让调度器有机会切换到其他 Goroutine
		time.Sleep(10 * time.Millisecond)
	}
}
