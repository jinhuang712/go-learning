package lesson_05_sync_atomic

import (
	"fmt"
)

func DemoSyncMutex() {
	fmt.Println("\n--- Mutex &amp; RWMutex 演示 ---")
	fmt.Println("请查看 Section 1 文档了解详细内容")
	fmt.Println("包含：")
	fmt.Println("  1. sync.Mutex - 互斥锁")
	fmt.Println("  2. sync.RWMutex - 读写锁")
}

func DemoSyncOther() {
	fmt.Println("\n--- WaitGroup, Once, Cond, Map 演示 ---")
	fmt.Println("请查看 Section 2 文档了解详细内容")
	fmt.Println("包含：")
	fmt.Println("  1. sync.WaitGroup - 等待多个 Goroutine")
	fmt.Println("  2. sync.Once - 只执行一次")
	fmt.Println("  3. sync.Cond - 条件变量")
	fmt.Println("  4. sync.Map - 并发安全 Map")
}

func DemoAtomic() {
	fmt.Println("\n--- Atomic 原子操作演示 ---")
	fmt.Println("请查看 Section 3 文档了解详细内容")
	fmt.Println("包含：")
	fmt.Println("  1. 基本原子操作 (Add, Load, Store, CAS)")
	fmt.Println("  2. 原子类型 (Go 1.19+)")
}
