package lesson_02_channel

import (
	"fmt"
)

func DemoChannelBasics() {
	fmt.Println("\n--- Channel 基础演示 ---")

	ch := make(chan int)

	go func() {
		fmt.Println("Goroutine: 发送数据 42")
		ch <- 42
		fmt.Println("Goroutine: 发送完成")
	}()

	fmt.Println("Main: 等待接收...")
	val := <-ch
	fmt.Printf("Main: 接收到 %d\n", val)

	fmt.Println("\n--- 有缓冲 Channel ---")
	bufferedCh := make(chan string, 2)

	bufferedCh <- "hello"
	bufferedCh <- "world"
	fmt.Println("缓冲区已满，再发送会阻塞...")

	fmt.Println("接收:", <-bufferedCh)
	fmt.Println("接收:", <-bufferedCh)

	fmt.Println("\n--- 关闭 Channel ---")
	dataCh := make(chan int, 3)
	dataCh <- 1
	dataCh <- 2
	dataCh <- 3
	close(dataCh)

	for v := range dataCh {
		fmt.Println("遍历收到:", v)
	}
}

func DemoChannelPatterns() {
	fmt.Println("\n--- Channel 常用模式演示 ---")
	fmt.Println("请查看 Section 2 文档了解详细模式")
	fmt.Println("常用模式包括：")
	fmt.Println("  1. 生产者-消费者模式")
	fmt.Println("  2. 扇入 (Fan-in)")
	fmt.Println("  3. 扇出 (Fan-out)")
	fmt.Println("  4. 信号量模式（限制并发）")
}

func DemoChannelUnderhood() {
	fmt.Println("\n--- Channel 底层原理 ---")
	fmt.Println("请查看 Section 3 文档了解 hchan 结构")
	fmt.Println("hchan 包含：")
	fmt.Println("  - buf: 环形缓冲区")
	fmt.Println("  - sendx/recvx: 发送/接收索引")
	fmt.Println("  - recvq/sendq: 等待队列")
	fmt.Println("  - lock: 互斥锁")
}
