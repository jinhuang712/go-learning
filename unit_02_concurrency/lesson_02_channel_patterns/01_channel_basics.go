package lesson_02_channel_patterns

import "fmt"

// DemoChannelBasics 演示 Channel 基础语法
func DemoChannelBasics() {
	fmt.Println("\n--- Section 1: Channel 基础语法 ---")

	// 创建 Channel
	ch := make(chan string)

	// 启动 Goroutine 发送数据
	go func() {
		ch <- "Hello from Channel!"
	}()

	// 接收数据
	msg := <-ch
	fmt.Println("收到:", msg)

	// 演示类型安全
	intCh := make(chan int)
	go func() {
		intCh <- 42
	}()
	fmt.Println("收到整数:", <-intCh)
}
