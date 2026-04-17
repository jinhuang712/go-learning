package main

import (
	"fmt"
	"go-learning/unit_01_fundamentals/lesson_00_packages/code"
	"go-learning/unit_01_fundamentals/lesson_01_basics/code"
	"go-learning/unit_01_fundamentals/lesson_02_collections/code"
	"go-learning/unit_01_fundamentals/lesson_03_structs_interfaces/code"
	"go-learning/unit_01_fundamentals/lesson_04_pointers/code"
	"go-learning/unit_01_fundamentals/lesson_05_error_handling/code"
	"go-learning/unit_01_fundamentals/lesson_06_java_go_pitfalls/code"
	"go-learning/unit_02_concurrency/lesson_01_goroutine/code"
)

func main() {
	fmt.Println("Starting Go Learning Course...")
	fmt.Println("=================================")

	// 运行第零课：Go Package 系统详解
	packages.Run()

	// 运行第一课：基础语法
	basics.Run()

	// 运行第二课：集合类型
	lesson_02_collections.Run()

	// 运行第三课：结构体与接口
	lesson_03_structs_interfaces.Run()

	// 运行第四课：指针与值语义
	pointers.Run()

	// 运行第五课：错误处理与 Panic/Recover
	lesson_05_error_handling.Run()

	// 运行第六课：Java 转 Go 常见坑点专项
	lesson_06_java_go_pitfalls.Run()

	// 运行 Unit 2 Lesson 1: Goroutine &amp; CSP 模型基础
	lesson_01_goroutine.Run()

	fmt.Println("\n=================================")
	fmt.Println("Course execution finished.")
}
