package lesson_05_error_handling

import (
	"fmt"
)

// Run 统一执行 Lesson 5 的所有示例
func Run() {
	fmt.Println("\n=== Lesson 5: Error Handling & Panic/Recover ===")
	DemoErrorInterface()
	DemoPanicRecover()
}
