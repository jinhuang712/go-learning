# Section 3: Atomic 原子操作

## 1. 知识点核心说明

### 基本原子操作
atomic 包提供了基本的原子操作，无需加锁。

```go
import "sync/atomic"

var counter int64

atomic.AddInt64(&amp;counter, 1)        // 加法
atomic.LoadInt64(&amp;counter)            // 读取
atomic.StoreInt64(&amp;counter, 100)       // 存储
atomic.CompareAndSwapInt64(&amp;counter, 0, 1) // CAS
```

### 原子类型
Go 1.19+ 提供了更易用的原子类型：

```go
import "sync/atomic"

var counter atomic.Int64

counter.Add(1)
counter.Load()
counter.Store(100)
counter.CompareAndSwap(0, 1)
```

---

## 2. Java 与 Go 的对比说明

| Java | Go | 说明 |
|------|----|------|
| `AtomicInteger` | `atomic.Int32` / `atomic.AddInt32()` | 原子整数 |
| `AtomicLong` | `atomic.Int64` / `atomic.AddInt64()` | 原子长整数 |
| `AtomicReference` | `atomic.Value` | 原子引用 |
| `AtomicInteger.compareAndSet()` | `atomic.CompareAndSwapInt32()` | CAS 操作 |

### 代码对照

**Java (AtomicInteger)**:
```java
import java.util.concurrent.atomic.AtomicInteger;

public class Counter {
    private final AtomicInteger count = new AtomicInteger(0);
    
    public void increment() {
        count.incrementAndGet();
    }
    
    public int getCount() {
        return count.get();
    }
}
```

**Go (atomic 包)**:
```go
package main

import (
	"sync"
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

func main() {
	c := &amp;Counter{}
	var wg sync.WaitGroup

	for i := 0; i &lt; 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()
	println("Count:", c.GetCount())
}
```

**Go 1.19+ 原子类型**:
```go
package main

import (
	"sync"
	"sync/atomic"
)

type Counter struct {
	count atomic.Int64
}

func (c *Counter) Increment() {
	c.count.Add(1)
}

func (c *Counter) GetCount() int64 {
	return c.count.Load()
}
```

**心智迁移要点**:
- Java: `Atomic*` 类封装了原子操作
- Go: 旧版本用函数式 API (`atomic.AddInt64()`)，新版本用原子类型 (`atomic.Int64`)
- Java: `AtomicReference` 处理引用类型
- Go: `atomic.Value` 处理任意类型
- Java: CAS 操作通过 `compareAndSet()`
- Go: CAS 操作通过 `CompareAndSwap*()`
- 两者性能相似，都是基于 CPU 的原子指令

---

## 3. 可运行代码示例

### 示例 1: 原子计数器 vs 锁计数器

```go
package main

import (
	"sync"
	"sync/atomic"
	"time"
)

type AtomicCounter struct {
	count int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&amp;c.count, 1)
}

type MutexCounter struct {
	mu    sync.Mutex
	count int
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func main() {
	const num = 1000000

	// 原子计数器
	atomicCounter := &amp;AtomicCounter{}
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i &lt; num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomicCounter.Inc()
		}()
	}
	wg.Wait()
	println("Atomic counter:", atomicCounter.count, "time:", time.Since(start))

	// 锁计数器
	mutexCounter := &amp;MutexCounter{}
	start = time.Now()
	for i := 0; i &lt; num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutexCounter.Inc()
		}()
	}
	wg.Wait()
	println("Mutex counter:", mutexCounter.count, "time:", time.Since(start))
}
```

### 示例 2: atomic.Value

```go
package main

import (
	"sync"
	"sync/atomic"
)

type Config struct {
	Value string
}

func main() {
	var config atomic.Value
	config.Store(&amp;Config{Value: "v1"})

	var wg sync.WaitGroup

	// 读 Config
	for i := 0; i &lt; 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cfg := config.Load().(*Config)
			println("Read:", cfg.Value)
		}()
	}

	// 更新 Config
	wg.Add(1)
	go func() {
		defer wg.Done()
		config.Store(&amp;Config{Value: "v2"})
		println("Updated to v2")
	}()

	wg.Wait()
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 过度使用 atomic**
```go
// 不推荐：复杂逻辑用 atomic，难以维护
var count int64
atomic.AddInt64(&amp;count, 1)
if atomic.LoadInt64(&amp;count) &gt; 100 {
    // ...
}

// 推荐：复杂逻辑用 Mutex，更易读
var mu sync.Mutex
var count int
mu.Lock()
count++
if count &gt; 100 {
    // ...
}
mu.Unlock()
```

⚠️ **坑点 2: atomic.Value 存储不同类型**
```go
// 错误做法：atomic.Value 存储不同类型会 panic
var v atomic.Value
v.Store(1)
v.Store("string") // panic!

// 正确做法：atomic.Value 只能存储相同类型
var v atomic.Value
v.Store(&amp;Config{Value: "v1"})
v.Store(&amp;Config{Value: "v2"}) // OK
```

⚠️ **坑点 3: 非对齐的 64 位原子操作**
```go
// 错误做法：32 位系统上非对齐的 64 位原子操作会 panic
type Bad struct {
	b byte
	i int64 // 可能未对齐！
}

var bad Bad
atomic.AddInt64(&amp;bad.i, 1) // 32 位系统可能 panic

// 正确做法：确保 64 位对齐
type Good struct {
	i int64 // 放在第一位，确保对齐
	b byte
}

var good Good
atomic.AddInt64(&amp;good.i, 1) // 安全
```

---

## 5. 练习题

### 练习 3.1: 实现无锁计数器
**目标**: 使用 atomic 包实现一个无锁的计数器，并与 Mutex 版本做性能对比。

**验收标准**:
- 正确使用 `atomic.AddInt64()` 和 `atomic.LoadInt64()`
- 启动 1000000 个 Goroutine 并发递增
- 最终结果应该是 1000000
- 简单对比一下 atomic 和 Mutex 的性能（可选）

### 练习 3.2: 理解 atomic（概念题）
**问题**:
1. 什么时候用 atomic，什么时候用 Mutex？
2. CAS（Compare And Swap）操作是什么？为什么它是原子操作的基础？

**验收标准**: 能够清晰解释这两个问题。

---

查看 [code/sync_atomic.go](./code/sync_atomic.go) 中的 `DemoAtomic()` 函数来运行示例代码。
