package lesson_02_channel_patterns

import "fmt"

// DemoChannelDirection 演示 Channel 方向约束
func DemoChannelDirection() {
	fmt.Println("\n--- Section 5: Channel 方向约束 ---")

	ch := make(chan string)

	// 发送函数（只写 Channel）
	sender := func(ch chan<- string, msg string) {
		ch <- msg
	}

	// 接收函数（只读 Channel）
	receiver := func(ch <-chan string) string {
		return <-ch
	}

	go sender(ch, "方向约束消息")
	fmt.Println("  收到:", receiver(ch))
}
