package lesson_02_channel_patterns

import "fmt"

// DemoBufferedChannel 演示有缓冲 Channel
func DemoBufferedChannel() {
	fmt.Println("\n--- Section 3: 有缓冲 Channel（异步通信） ---")

	// 创建容量为 3 的有缓冲 Channel
	ch := make(chan int, 3)

	fmt.Println("  发送 3 个数据到缓冲 Channel（不会阻塞）...")
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("  发送完成，缓冲区已满")

	fmt.Println("  接收数据:")
	fmt.Println("  ", <-ch)
	fmt.Println("  ", <-ch)
	fmt.Println("  ", <-ch)
}
