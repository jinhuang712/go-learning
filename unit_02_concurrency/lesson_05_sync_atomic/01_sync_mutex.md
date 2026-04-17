# Section 1: Mutex &amp; RWMutex

## 1. 知识点核心说明

### sync.Mutex
Mutex 是互斥锁，用于保护临界区，同一时间只有一个 Goroutine 能持有锁。

```go
var mu sync.Mutex

mu.Lock()
defer mu.Unlock()
// 临界区
```

### sync.RWMutex
RWMutex 是读写锁，允许多个 Goroutine 同时读，但写的时候独占。

```go
var rwMu sync.RWMutex

// 读锁
rwMu.RLock()
defer rwMu.RUnlock()

// 写锁
rwMu.Lock()
defer rwMu.Unlock()
```

---

## 2. Java 与 Go 的对比说明

| 特性 | Java | Go |
|------|------|----|
| **互斥锁** | `ReentrantLock` | `sync.Mutex` |
| **读写锁** | `ReentrantReadWriteLock` | `sync.RWMutex` |
| **可重入** | 支持 | 不支持 |
| **公平锁** | 可选 | 不支持 |
| **Condition** | 支持 | 通过 `sync.Cond` 单独提供 |

### 代码对照

**Java (ReentrantLock)**:
```java
import java.util.concurrent.locks.ReentrantLock;

public class Counter {
    private final ReentrantLock lock = new ReentrantLock();
    private int count = 0;
    
    public void increment() {
        lock.lock();
        try {
            count++;
        } finally {
            lock.unlock();
        }
    }
    
    public int getCount() {
        lock.lock();
        try {
            return count;
        } finally {
            lock.unlock();
        }
    }
}
```

**Go (sync.Mutex)**:
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

**Java (ReentrantReadWriteLock)**:
```java
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class Cache {
    private final ReentrantReadWriteLock rwLock = new ReentrantReadWriteLock();
    private Map&lt;String, String&gt; data = new HashMap&lt;&gt;();
    
    public String get(String key) {
        rwLock.readLock().lock();
        try {
            return data.get(key);
        } finally {
            rwLock.readLock().unlock();
        }
    }
    
    public void set(String key, String value) {
        rwLock.writeLock().lock();
        try {
            data.put(key, value);
        } finally {
            rwLock.writeLock().unlock();
        }
    }
}
```

**Go (sync.RWMutex)**:
```go
package main

import (
	"sync"
)

type Cache struct {
	rwMu sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &amp;Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) string {
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()
	return c.data[key]
}

func (c *Cache) Set(key, value string) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()
	c.data[key] = value
}
```

**心智迁移要点**:
- Java: `ReentrantLock` 可重入、支持公平锁、功能强大
- Go: `sync.Mutex` 简单、轻量、不可重入
- Java: 锁和 Condition 绑定在一起
- Go: 锁和 Condition 分开（`sync.Cond`）
- Java: `finally` 块确保解锁
- Go: `defer` 确保解锁（更简洁）
- Java: 读写锁通过 `readLock()` / `writeLock()` 获取
- Go: 读写锁通过 `RLock()` / `Lock()` 获取

---

## 3. 可运行代码示例

### 错误写法：忘记解锁

```go
// 错误写法：忘记解锁，导致死锁
var mu sync.Mutex

func bad() {
	mu.Lock()
	// 忘记 mu.Unlock()！
	// 如果这个函数返回或 panic，锁永远不会被释放
}
```

### 正确写法：使用 defer 解锁

```go
// 正确写法：defer 确保解锁
var mu sync.Mutex

func good() {
	mu.Lock()
	defer mu.Unlock() // 无论函数怎么返回，都会执行
	// 临界区
}
```

### RWMutex 适用场景

```go
package main

import (
	"sync"
	"time"
)

type Config struct {
	rwMu sync.RWMutex
	data map[string]string
}

func (c *Config) Get(key string) string {
	c.rwMu.RLock() // 读锁，多个 Goroutine 可以同时读
	defer c.rwMu.RUnlock()
	return c.data[key]
}

func (c *Config) Set(key, value string) {
	c.rwMu.Lock() // 写锁，独占
	defer c.rwMu.Unlock()
	c.data[key] = value
}

func main() {
	cfg := &amp;Config{data: make(map[string]string)}
	cfg.Set("key", "value")
	
	var wg sync.WaitGroup
	
	// 启动多个读 Goroutine
	for i := 0; i &lt; 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j &lt; 100; j++ {
				_ = cfg.Get("key")
			}
		}()
	}
	
	wg.Wait()
	println("All reads done")
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: Mutex 不可重入**
```go
var mu sync.Mutex

func foo() {
	mu.Lock()
	defer mu.Unlock()
	bar() // 会死锁！foo 已经持有锁了
}

func bar() {
	mu.Lock() // 死锁！
	defer mu.Unlock()
}
```

⚠️ **坑点 2: 复制 Mutex**
```go
var mu sync.Mutex

// 错误写法：复制 Mutex
func bad(mu sync.Mutex) { // 值传递，复制了 Mutex！
	mu.Lock()
	// ...
}

// 正确写法：传递指针
func good(mu *sync.Mutex) {
	mu.Lock()
	// ...
}
```

⚠️ **坑点 3: 过度使用 RWMutex**
```go
// 不推荐：读操作很少时，RWMutex 比 Mutex 还慢
type BadCache struct {
	rwMu sync.RWMutex // 读很少，写很多
	data map[string]string
}

// 推荐：直接用 Mutex
type GoodCache struct {
	mu   sync.Mutex
	data map[string]string
}
```

---

## 5. 练习题

### 练习 1.1: 实现线程安全的计数器
**目标**: 写一个线程安全的计数器，使用 Mutex 保护，支持 Increment() 和 Get() 两个方法。

**验收标准**:
- 正确使用 Mutex
- 使用 defer 确保解锁
- 启动 1000 个 Goroutine 并发调用 Increment()
- 最终结果应该是 1000

### 练习 1.2: 理解 Mutex（概念题）
**问题**:
1. Go 的 Mutex 为什么不支持可重入？这是设计缺陷还是有意为之？
2. RWMutex 在什么场景下比 Mutex 更有优势？什么场景下反而更差？

**验收标准**: 能够清晰解释这两个问题。

---

查看 [code/sync_atomic.go](./code/sync_atomic.go) 中的 `DemoSyncMutex()` 函数来运行示例代码。
