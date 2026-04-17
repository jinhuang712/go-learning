package lesson_04_context

import (
	"context"
	"fmt"
)

func DemoContextBasics() {
	fmt.Println("\n--- Context 基础演示 ---")

	ctx := context.Background()
	fmt.Println("context.Background():", ctx)

	todoCtx := context.TODO()
	fmt.Println("context.TODO():", todoCtx)
}

func DemoContextDerived() {
	fmt.Println("\n--- 派生 Context 演示 ---")
	fmt.Println("请查看 Section 2 文档了解详细内容")
	fmt.Println("派生方式包括：")
	fmt.Println("  1. WithCancel - 可取消")
	fmt.Println("  2. WithTimeout / WithDeadline - 超时")
	fmt.Println("  3. WithValue - 传值")
}

func DemoContextPropagation() {
	fmt.Println("\n--- Context 传递最佳实践 ---")
	fmt.Println("请查看 Section 3 文档了解详细内容")
	fmt.Println("最佳实践包括：")
	fmt.Println("  1. Context 总是第一个参数")
	fmt.Println("  2. 不要存在结构体里")
	fmt.Println("  3. WithValue key 使用自定义类型")
}
