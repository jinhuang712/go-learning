# Section 3: 修复数据竞争

## 1. 知识点核心说明

### 常见修复方案

| 方案 | 适用场景 | 性能 |
|------|---------|------|
| sync.Mutex | 通用场景 | 中 |
| sync.RWMutex | 读多写少 | 中高 |
| Channel | CSP 风格 | 中 |
| sync/atomic | 简单计数 | 高 |

---

## 2. Java 与 Go 的对比说明

| 方案 | Java | Go |
|------|------|----|
| **互斥锁** | `synchronized` / `ReentrantLock` | `sync.Mutex` |
| **读写锁** | `ReentrantReadWriteLock` | `sync.RWMutex` |
| **原子操作** | `AtomicInteger` 等 | `sync/atomic` 包 |
| **并发集合** | `ConcurrentHashMap` | `sync.Map` |

### 代码对照

**Java (修复数据竞争)**:
```java
import java.util.concurrent.atomic.AtomicInteger;

public class FixedCounter {
    private final AtomicInteger count = new AtomicInteger(0);
    
    public void increment() {
        count.incrementAndGet();
    }
    
    public int getCount() {
        return count.get();
    }
}
```

**Go (修复数据竞争 - Mutex)**:
```go
package main

import (
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
```

**Go (修复数据竞争 - Atomic)**:
```go
package main

import (
	"sync/atomic"
)

type Counter struct {
	count int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&amp;c.count, 1)
}

func (c *Counter) GetCount() int64 {
	return atomic.LoadInt64(&amp;c.count)
}
```

**Go (修复数据竞争 - Channel)**:
```go
package main

type Counter struct {
	ch chan int
}

func NewCounter() *Counter {
	c := &amp;Counter{ch: make(chan int, 1)}
	c.ch &lt;- 0
	return c
}

func (c *Counter) Increment() {
	count := &lt;-c.ch
	count++
	c.ch &lt;- count
}

func (c *Counter) GetCount() int {
	count := &lt;-c.ch
	c.ch &lt;- count
	return count
}
```

**心智迁移要点**:
- Java: 有多种修复方案，选择取决于场景
- Go: 同样有多种修复方案，Mutex 最通用，Atomic 性能最高
- Java: `synchronized` 是语言特性，简单但不够灵活
- Go: `sync.Mutex` 是库，更灵活但需要手动管理
- Java: `ConcurrentHashMap` 等并发集合是标准库一部分
- Go: `sync.Map` 适用于特定场景，普通场景推荐 `RWMutex+map`

---

## 3. 可运行代码示例

### 示例 1: 三种修复方案对比

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 方案 1: Mutex
type MutexCounter struct {
	mu    sync.Mutex
	count int
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *MutexCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// 方案 2: Atomic
type AtomicCounter struct {
	count int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&amp;c.count, 1)
}

func (c *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&amp;c.count)
}

// 方案 3: Channel
type ChannelCounter struct {
	ch chan int
}

func NewChannelCounter() *ChannelCounter {
	c := &amp;ChannelCounter{ch: make(chan int, 1)}
	c.ch &lt;- 0
	return c
}

func (c *ChannelCounter) Inc() {
	count := &lt;-c.ch
	count++
	c.ch &lt;- count
}

func (c *ChannelCounter) Get() int {
	count := &lt;-c.ch
	c.ch &lt;- count
	return count
}

func main() {
	const num = 1000000

	// Mutex
	mc := &amp;MutexCounter{}
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i &lt; num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mc.Inc()
		}()
	}
	wg.Wait()
	fmt.Printf("Mutex: %d, time: %v\n", mc.Get(), time.Since(start))

	// Atomic
	ac := &amp;AtomicCounter{}
	start = time.Now()
	for i := 0; i &lt; num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ac.Inc()
		}()
	}
	wg.Wait()
	fmt.Printf("Atomic: %d, time: %v\n", ac.Get(), time.Since(start))

	// Channel (比较慢，仅用于演示)
	cc := NewChannelCounter()
	start = time.Now()
	for i := 0; i &lt; 1000; i++ { // 少跑一些
		wg.Add(1)
		go func() {
			defer wg.Done()
			cc.Inc()
		}()
	}
	wg.Wait()
	fmt.Printf("Channel: %d, time: %v\n", cc.Get(), time.Since(start))
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 过度使用 Channel**
```go
// 不推荐：简单的计数器用 Channel，性能差
type BadCounter struct {
	ch chan int
}

// 推荐：简单的计数器用 Atomic 或 Mutex
```

⚠️ **坑点 2: 用 sync.Map 替代所有 map**
```go
// 不推荐：写多读少时 sync.Map 性能比 RWMutex+map 差
var m sync.Map

// 推荐：根据场景选择
var (
	mu sync.RWMutex
	m  = make(map[string]string)
)
```

⚠️ **坑点 3: 不验证修复结果**
```go
// 错误做法：修复后不验证
// 正确做法：修复后用 race detector 再次验证
go test -race ./...
```

---

## 5. 练习题

### 练习 3.1: 修复数据竞争
**目标**: 修复下面代码中的数据竞争，使用三种不同的方案（Mutex、Atomic、Channel）。

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i &lt; 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // 这里有数据竞争！
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
```

**验收标准**:
- 使用 Mutex 修复
- 使用 Atomic 修复
- 使用 Channel 修复
- 每种方案都能用 race detector 验证没有数据竞争
- 简单对比一下三种方案的性能（可选）

### 练习 3.2: 理解修复方案（概念题）
**问题**:
1. 什么时候用 Mutex，什么时候用 Atomic，什么时候用 Channel？
2. 这三种方案在性能上有什么差异？

**验收标准**: 能够清晰解释这两个问题。

---

查看 [code/race_detector.go](./code/race_detector.go) 中的 `DemoFixingRaces()` 函数来了解更多示例。
