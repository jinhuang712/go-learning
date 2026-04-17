package lesson_03_select

import (
	"fmt"
	"time"
)

func DemoSelectBasics() {
	fmt.Println("\n--- Select 基础演示 ---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "来自 Channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "来自 Channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("收到:", msg1)
		case msg2 := <-ch2:
			fmt.Println("收到:", msg2)
		}
	}
}

func DemoSelectPatterns() {
	fmt.Println("\n--- Select 模式演示 ---")
	fmt.Println("请查看 Section 2 文档了解详细模式")
	fmt.Println("常用模式包括：")
	fmt.Println("  1. 超时控制 (time.After)")
	fmt.Println("  2. 非阻塞读写 (default)")
	fmt.Println("  3. 多路等待")
}

func DemoSelectPitfalls() {
	fmt.Println("\n--- Select 坑点演示 ---")
	fmt.Println("请查看 Section 3 文档了解常见坑点")
}
