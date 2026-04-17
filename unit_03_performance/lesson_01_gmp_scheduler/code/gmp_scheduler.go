package lesson_01_gmp_scheduler

import (
	"fmt"
	"runtime"
)

func DemoGMPModel() {
	fmt.Println("\n--- G-M-P 模型演示 ---")
	fmt.Println("请查看 Section 1 文档了解详细内容")
	fmt.Println("G-M-P 模型组件：")
	fmt.Println("  M (Machine): 操作系统线程")
	fmt.Println("  P (Processor): 逻辑处理器")
	fmt.Println("  G (Goroutine): Goroutine")
	fmt.Printf("当前 GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("当前 Goroutine 数量: %d\n", runtime.NumGoroutine())
}

func DemoSchedulingTriggers() {
	fmt.Println("\n--- Goroutine 调度时机演示 ---")
	fmt.Println("请查看 Section 2 文档了解详细内容")
	fmt.Println("Goroutine 调度会在以下情况发生：")
	fmt.Println("  1. 函数调用（尤其是栈扩容时）")
	fmt.Println("  2. 系统调用后")
	fmt.Println("  3. 显式调用 runtime.Gosched()")
	fmt.Println("  4. 使用锁、Channel 等同步原语")
}

func DemoBlockingPrinciples() {
	fmt.Println("\n--- 阻塞原理与处理演示 ---")
	fmt.Println("请查看 Section 3 文档了解详细内容")
}
