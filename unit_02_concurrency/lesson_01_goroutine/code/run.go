package lesson_01_goroutine

import "fmt"

func Run() {
	fmt.Println("\n=== Lesson 1: Goroutine &amp; CSP 模型基础 ===")

	DemoGoroutineBasics()
	DemoGoroutineScheduling()
	DemoGoroutineLeak()
}
