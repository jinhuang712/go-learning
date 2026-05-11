package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("主 Goroutine 开始执行")
	fmt.Printf("当前 Goroutine 数量: %d\n\n", runtime.NumGoroutine())

	go func() {
		fmt.Println("匿名函数 Goroutine 开始执行")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("匿名函数 Goroutine 执行完毕")
	}()

	go sayHello()

	fmt.Printf("启动 2 个 Goroutine 后，当前 Goroutine 数量: %d\n\n", runtime.NumGoroutine())

	time.Sleep(1 * time.Second)
	fmt.Println("主 Goroutine 执行完毕")
}

func sayHello() {
	fmt.Println("sayHello Goroutine 开始执行")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("sayHello Goroutine 执行完毕")
}
