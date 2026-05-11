package lesson_05_sync_atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Run 运行 Sync 包与 Atomic 原子操作课程代码
func Run() {
	fmt.Println("\n=== Lesson 5: Sync 包与 Atomic 原子操作 ===")

	DemoMutex()
	DemoRWMutex()
	DemoOnce()
	DemoAtomic()
}

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func DemoMutex() {
	fmt.Println("\n--- Mutex - 互斥锁 ---")

	counter := &SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("  计数结果: %d\n", counter.Get())
}

type SafeReadWrite struct {
	mu    sync.RWMutex
	value string
}

func (s *SafeReadWrite) Read() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.value
}

func (s *SafeReadWrite) Write(v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.value = v
}

func DemoRWMutex() {
	fmt.Println("\n--- RWMutex - 读写锁 ---")

	rw := &SafeReadWrite{}
	rw.Write("初始值")

	var wg sync.WaitGroup

	// 启动多个读操作
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("  Reader %d: 读取到 %s\n", id, rw.Read())
		}(i)
	}

	wg.Wait()
}

var singletonOnce sync.Once
var singletonInstance *Singleton

type Singleton struct {
	Value string
}

func GetSingleton() *Singleton {
	singletonOnce.Do(func() {
		fmt.Println("  单例实例化（仅执行一次）")
		singletonInstance = &Singleton{Value: "唯一实例"}
	})
	return singletonInstance
}

func DemoOnce() {
	fmt.Println("\n--- Once - 单例模式 ---")

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			instance := GetSingleton()
			fmt.Printf("  Goroutine %d: 获取到 %s\n", id, instance.Value)
		}(i)
	}

	wg.Wait()
}

type AtomicCounter struct {
	count int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func (c *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&c.count)
}

func DemoAtomic() {
	fmt.Println("\n--- Atomic 原子操作 ---")

	counter := &AtomicCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("  原子计数结果: %d\n", counter.Get())
}
