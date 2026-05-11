package lesson_01_goroutine_csp

import "fmt"

// Run 运行 Goroutine & CSP 模型基础课程代码
func Run() {
	fmt.Println("\n=== Lesson 1: Goroutine & CSP 模型基础 ===")

	DemoGoroutineBasics()
	DemoGoroutineScheduling()
	DemoCSPModel()
	DemoWaitGroupBasics()
}
