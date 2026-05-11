package lesson_02_channel_patterns

import "fmt"

// DemoChannelClose 演示 Channel 关闭与广播
func DemoChannelClose() {
	fmt.Println("\n--- Section 4: Channel 关闭与广播 ---")

	// 场景 1: 关闭后继续接收
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	fmt.Println("  从已关闭的 Channel 接收:")
	fmt.Println("  ", <-ch)
	fmt.Println("  ", <-ch)
	fmt.Println("  ", <-ch) // 缓冲区空后返回零值

	// 场景 2: 使用 comma-ok 判断
	data, ok := <-ch
	fmt.Printf("  接收数据: %d, Channel 关闭? %v\n", data, !ok)

	// 场景 3: 广播模式
	done := make(chan struct{})
	for i := 1; i <= 3; i++ {
		go func(id int) {
			<-done
			fmt.Printf("  Goroutine %d: 收到广播信号\n", id)
		}(i)
	}
	fmt.Println("  发送广播信号...")
	close(done)
}
