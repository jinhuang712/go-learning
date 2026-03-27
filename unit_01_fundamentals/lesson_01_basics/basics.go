package basics

import (
	"fmt"
	"go-learning/pkg/calc"
)

// Run 运行基础语法课程代码
func Run() {
	fmt.Println("=== Lesson 1: Basics (Variables, Functions, Control Flow) ===")

	// 1. 显式变量声明
	var greeting string = "Hello, Java Developer!"
	var year int = 2024

	// 2. 类型推断
	var language = "Go"

	// 3. 短变量声明
	message := "Welcome to the world of Go."

	// 打印输出
	fmt.Println(greeting)
	fmt.Printf("It is %d. We are learning %s.\n", year, language)
	fmt.Println(message)

	// 4. 控制流 (Loop, Switch, Defer)
	fmt.Println("\n--- Control Flow ---")
	
	// Defer 演示
	defer fmt.Println("This is deferred and will print at the very end of Run().")

	// for 循环
	for i := 0; i < 3; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// if 带初始化的特殊写法
	// mockError() 是一个假装返回错误的函数，我们把它声明在初始化语句中
	if err := mockError(); err != nil {
		fmt.Printf("Error handled inside if block: %v\n", err)
	}
	// 注意：在这里你是无法访问 err 变量的，因为它的作用域只在上面的 if 块内

	// switch 演示 (不需要 break)
	day := "Monday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a workday.")
	}

	// 5. 函数调用
	sum := add(10, 20)
	fmt.Printf("\nSum of 10 and 20 is: %d\n", sum)

	// 6. 多返回值演示
	q, r := divMod(10, 3)
	fmt.Printf("10 / 3 = %d remainder %d\n", q, r)

	// 7. 可见性演示 (Package Visibility)
	fmt.Println("\n--- Visibility Demo ---")
	res := calc.PublicFunc(10, 20)
	fmt.Printf("calc.PublicFunc(10, 20) = %d\n", res)
	fmt.Printf("calc.PublicConst = %d\n", calc.PublicConst)
}

// add 仅在 basics 包内可见
func add(a int, b int) int {
	return a + b
}

// mockError 模拟一个返回错误的函数
func mockError() error {
	return fmt.Errorf("this is a simulated error")
}

// divMod 仅在 basics 包内可见
func divMod(a, b int) (int, int) {
	return a / b, a % b
}
