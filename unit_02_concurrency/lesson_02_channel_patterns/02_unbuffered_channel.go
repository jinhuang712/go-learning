package lesson_02_channel_patterns

import (
	"fmt"
	"time"
)

// DemoUnbufferedChannel 演示无缓冲 Channel
func DemoUnbufferedChannel() {
	fmt.Println("\n--- Section 2: 无缓冲 Channel（同步通信） ---")

	ch := make(chan string)

	// 无缓冲 Channel：发送者和接收者必须同时准备好
	go func() {
		fmt.Println("  Goroutine: 准备发送数据...")
		time.Sleep(500 * time.Millisecond)
		ch <- "同步消息"
		fmt.Println("  Goroutine: 数据已发送")
	}()

	fmt.Println("  Main: 等待接收...")
	msg := <-ch
	fmt.Println("  Main: 收到:", msg)
}
