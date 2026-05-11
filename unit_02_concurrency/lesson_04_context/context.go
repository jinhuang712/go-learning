package lesson_04_context

import (
	"context"
	"fmt"
	"time"
)

// Run 运行 Context 传递与级联取消课程代码
func Run() {
	fmt.Println("\n=== Lesson 4: Context 传递与级联取消 ===")

	DemoContextBasics()
	DemoWithCancel()
	DemoWithTimeout()
	DemoWithValue()
}

func DemoContextBasics() {
	fmt.Println("\n--- Context 基础 ---")

	// 创建根 Context
	ctx := context.Background()
	fmt.Println("  context.Background():", ctx)

	// TODO Context（用于不确定用什么时）
	ctx = context.TODO()
	fmt.Println("  context.TODO():", ctx)
}

func DemoWithCancel() {
	fmt.Println("\n--- WithCancel - 手动取消 ---")

	ctx, cancel := context.WithCancel(context.Background())

	// 启动一个监听 Context 的 Goroutine
	go func() {
		<-ctx.Done()
		fmt.Println("  Goroutine: 收到取消信号！")
	}()

	// 等待一下然后取消
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Main: 发送取消信号...")
	cancel()

	time.Sleep(50 * time.Millisecond)
}

func DemoWithTimeout() {
	fmt.Println("\n--- WithTimeout - 超时取消 ---")

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 模拟一个耗时操作
	done := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("  操作完成")
	case <-ctx.Done():
		fmt.Println("  ⚠️  超时:", ctx.Err())
	}
}

func DemoWithValue() {
	fmt.Println("\n--- WithValue - 元数据传递 ---")

	type contextKey string
	const requestIDKey contextKey = "requestID"

	// 创建带值的 Context
	ctx := context.WithValue(context.Background(), requestIDKey, "req-12345")

	// 在函数间传递 Context
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	type contextKey string
	const requestIDKey contextKey = "requestID"

	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Printf("  处理请求, RequestID: %s\n", requestID)
	}
}
