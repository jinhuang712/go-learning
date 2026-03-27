package main

import (
	"fmt"
	"go-learning/unit_01_fundamentals/lesson_01_basics"
	"go-learning/unit_01_fundamentals/lesson_02_collections"
	"go-learning/unit_01_fundamentals/lesson_03_structs_interfaces"
	"go-learning/unit_01_fundamentals/lesson_04_pointers"
	"go-learning/unit_01_fundamentals/lesson_05_error_handling"
)

func main() {
	fmt.Println("Starting Go Learning Course...")
	fmt.Println("=================================")

	// 运行第一课：基础语法
	basics.Run()

	// 运行第二课：集合类型
	lesson_02_collections.Run()
	
	// 运行第三课：结构体与接口
	structs_interfaces.Run()

	// 运行第四课：指针与值语义
	pointers.Run()

	// 运行第五课：错误处理与 Panic/Recover
	lesson_05_error_handling.Run()

	fmt.Println("\n=================================")
	fmt.Println("Course execution finished.")
}
