package lesson_01_goroutine_csp

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

// TestGoroutineBasics 测试 Goroutine 基本功能
func TestGoroutineBasics(t *testing.T) {
	done := make(chan bool)
	go func() {
		done <- true
	}()

	select {
	case <-done:
		// 测试通过
	case <-time.After(1 * time.Second):
		t.Fatal("Goroutine 没有在预期时间内完成")
	}
}

// TestWaitGroup 测试 WaitGroup 功能
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	counter := 0
	var mu sync.Mutex

	const iterations = 100

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	if counter != iterations {
		t.Errorf("期望计数器为 %d，实际为 %d", iterations, counter)
	}
}

// TestGoroutineLeakCheck 测试是否有 Goroutine 泄露（简单检测）
func TestGoroutineLeakCheck(t *testing.T) {
	// 获取测试前的 Goroutine 数量
	before := runtime.NumGoroutine()

	// 运行一些测试代码
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
		}()
	}
	wg.Wait()

	// 给一点时间让 Goroutine 清理
	time.Sleep(50 * time.Millisecond)

	// 获取测试后的 Goroutine 数量
	after := runtime.NumGoroutine()

	// 允许少量差异（可能有后台 Goroutine）
	if after-before > 5 {
		t.Logf("警告：Goroutine 数量从 %d 增加到 %d，可能存在泄露", before, after)
	}
}

// BenchmarkGoroutineCreation 测试 Goroutine 创建性能
func BenchmarkGoroutineCreation(b *testing.B) {
	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 空操作
		}()
	}
	wg.Wait()
}
