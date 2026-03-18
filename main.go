package main

import (
	"fmt"
	"go-learning/lessons/basics"
	"go-learning/lessons/collections"
	"go-learning/lessons/structs_interfaces"
)

func main() {
	fmt.Println("Starting Go Learning Course...")
	fmt.Println("=================================")

	// 运行第一课：基础语法
	basics.Run()

	// 运行第二课：集合类型
	collections.Run()
	
	// 运行第三课：结构体与接口
	structs_interfaces.Run()

	fmt.Println("\n=================================")
	fmt.Println("Course execution finished.")
}
