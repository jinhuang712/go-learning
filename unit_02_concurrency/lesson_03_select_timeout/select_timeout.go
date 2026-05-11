package lesson_03_select_timeout

import (
	"fmt"
	"time"
)

// Run 运行 Select 多路复用与超时控制课程代码
func Run() {
	fmt.Println("\n=== Lesson 3: Select 多路复用与超时控制 ===")

	DemoSelectBasics()
	DemoTimeoutPattern()
	DemoNonBlocking()
	DemoFanIn()
	DemoFanOut()
}

func DemoSelectBasics() {
	fmt.Println("\n--- Section 1: Select 基础语法 ---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自 Channel 1"
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		ch2 <- "来自 Channel 2"
	}()

	// select 会选择先就绪的 Channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("  收到:", msg1)
		case msg2 := <-ch2:
			fmt.Println("  收到:", msg2)
		}
	}
}

func DemoTimeoutPattern() {
	fmt.Println("\n--- Section 2: 超时控制模式 ---")

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "耗时操作完成"
	}()

	select {
	case result := <-ch:
		fmt.Println("  收到:", result)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("  ⚠️  操作超时！")
	}
}

func DemoNonBlocking() {
	fmt.Println("\n--- Section 3: 非阻塞 Channel 操作 ---")

	ch := make(chan int, 1)

	// 非阻塞发送
	select {
	case ch <- 1:
		fmt.Println("  发送成功")
	default:
		fmt.Println("  发送失败（Channel 满）")
	}

	// 再次发送会失败
	select {
	case ch <- 2:
		fmt.Println("  发送成功")
	default:
		fmt.Println("  ⚠️  发送失败（Channel 满）")
	}

	// 非阻塞接收
	select {
	case val := <-ch:
		fmt.Printf("  接收成功: %d\n", val)
	default:
		fmt.Println("  接收失败（Channel 空）")
	}
}

func DemoFanIn() {
	fmt.Println("\n--- Section 4: 扇入（Fan-in）模式 ---")

	ch1 := make(chan string)
	ch2 := make(chan string)
	merged := make(chan string)

	// 扇入 Goroutine
	go func() {
		for {
			select {
			case msg := <-ch1:
				merged <- msg
			case msg := <-ch2:
				merged <- msg
			}
		}
	}()

	// 发送数据
	go func() {
		ch1 <- "消息 A-1"
		ch1 <- "消息 A-2"
	}()

	go func() {
		ch2 <- "消息 B-1"
		ch2 <- "消息 B-2"
	}()

	// 接收合并后的数据
	for i := 0; i < 4; i++ {
		fmt.Println("  合并收到:", <-merged)
	}
}

func DemoFanOut() {
	fmt.Println("\n--- Section 5: 扇出（Fan-out）模式 ---")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// 启动 3 个 Worker
	for i := 1; i <= 3; i++ {
		go worker(i, tasks, results)
	}

	// 发送任务
	for j := 1; j <= 5; j++ {
		tasks <- j
	}
	close(tasks)

	// 收集结果
	for k := 0; k < 5; k++ {
		fmt.Println("  结果:", <-results)
	}
}

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("  Worker %d 处理任务 %d\n", id, task)
		time.Sleep(50 * time.Millisecond)
		results <- task * task
	}
}
