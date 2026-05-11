package lesson_02_channel_patterns

import "fmt"

// Run 运行 Channel 模式与底层原理课程代码
func Run() {
	fmt.Println("\n=== Lesson 2: Channel 模式与底层原理 ===")

	DemoChannelBasics()
	DemoUnbufferedChannel()
	DemoBufferedChannel()
	DemoChannelClose()
	DemoChannelDirection()
	DemoProducerConsumer()
}
