# Section 2: WaitGroup, Once, Cond, Map

## 1. 知识点核心说明

### sync.WaitGroup
WaitGroup 用于等待一组 Goroutine 完成。

```go
var wg sync.WaitGroup

wg.Add(3) // 增加计数器
go func() {
    defer wg.Done() // 减少计数器
    // 工作
}()

wg.Wait() // 等待计数器归零
```

### sync.Once
Once 确保一个函数只执行一次。

```go
var once sync.Once

once.Do(func() {
    // 只会执行一次
})
```

### sync.Cond
Cond 是条件变量，用于 Goroutine 之间等待或通知某个条件。

```go
var mu sync.Mutex
var cond = sync.NewCond(&amp;mu)

mu.Lock()
for !condition {
    cond.Wait() // 等待条件
}
mu.Unlock()

cond.Signal() // 唤醒一个等待者
cond.Broadcast() // 唤醒所有等待者
```

### sync.Map
sync.Map 是并发安全的 Map，适用于读多写少的场景。

```go
var m sync.Map

m.Store("key", "value")
val, ok := m.Load("key")
m.Delete("key")
```

---

## 2. Java 与 Go 的对比说明

| Go | Java | 说明 |
|----|------|------|
| `sync.WaitGroup` | `CountDownLatch` / `CyclicBarrier` | 等待多个任务完成 |
| `sync.Once` | 双重检查锁 + volatile | 只执行一次 |
| `sync.Cond` | `Condition` | 条件变量 |
| `sync.Map` | `ConcurrentHashMap` | 并发安全的 Map |

### 代码对照

**Java (CountDownLatch)**:
```java
import java.util.concurrent.CountDownLatch;

public class WaitGroupExample {
    public static void main(String[] args) throws InterruptedException {
        CountDownLatch latch = new CountDownLatch(3);
        
        for (int i = 0; i &lt; 3; i++) {
            final int id = i;
            new Thread(() -&gt; {
                try {
                    System.out.println("Worker " + id + " 工作中");
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                } finally {
                    latch.countDown();
                }
            }).start();
        }
        
        System.out.println("等待所有 Worker...");
        latch.await();
        System.out.println("所有 Worker 完成");
    }
}
```

**Go (sync.WaitGroup)**:
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i &lt; 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d 工作中\n", id)
			time.Sleep(1 * time.Second)
		}(i)
	}

	fmt.Println("等待所有 Worker...")
	wg.Wait()
	fmt.Println("所有 Worker 完成")
}
```

**心智迁移要点**:
- Java: `CountDownLatch` 只能用一次
- Go: `WaitGroup` 可以复用（在 Wait() 之后重新 Add()）
- Java: `CyclicBarrier` 可以循环使用
- Go: `WaitGroup` 更简单直接
- Java: Singleton 需要双重检查锁
- Go: `sync.Once` 更简洁安全
- Java: `Condition` 与 `Lock` 绑定
- Go: `sync.Cond` 需要显式传入 Mutex
- Java: `ConcurrentHashMap` 功能强大
- Go: `sync.Map` 适用特定场景（读多写少）

---

## 3. 可运行代码示例

### 示例 1: sync.Once 实现单例

```go
package main

import (
	"fmt"
	"sync"
)

type Config struct {
	Value string
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		fmt.Println("初始化 Config（只会执行一次）")
		instance = &amp;Config{Value: "default"}
	})
	return instance
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i &lt; 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cfg := GetConfig()
			fmt.Println("Got config:", cfg.Value)
		}()
	}

	wg.Wait()
}
```

### 示例 2: sync.Map

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// 存储
	m.Store("key1", "value1")
	m.Store("key2", "value2")

	// 加载
	val, ok := m.Load("key1")
	if ok {
		fmt.Println("key1:", val)
	}

	// 遍历
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %s\n", key, value)
		return true // 继续遍历
	})

	// 删除
	m.Delete("key2")
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: WaitGroup 的 Add() 不在 Goroutine 外**
```go
// 错误写法：Add() 在 Goroutine 内部
go func() {
	wg.Add(1) // 太晚了！可能 Wait() 已经执行了
	defer wg.Done()
}()

// 正确写法：Add() 在 Goroutine 外
wg.Add(1)
go func() {
	defer wg.Done()
}()
```

⚠️ **坑点 2: sync.Map 在写多读少时性能差**
```go
// 不推荐：写多读少的场景用 sync.Map
var m sync.Map
// 写多读少时，性能比 RWMutex+map 差

// 推荐：写多读少时用 RWMutex+map
var (
	mu sync.RWMutex
	m  = make(map[string]string)
)
```

⚠️ **坑点 3: sync.Cond 的 Wait() 必须在循环中**
```go
// 错误写法：Wait() 不在循环中
mu.Lock()
if !condition {
    cond.Wait() // 虚假唤醒时会出问题！
}
mu.Unlock()

// 正确写法：Wait() 在循环中
mu.Lock()
for !condition { // 用 for 而不是 if
    cond.Wait()
}
mu.Unlock()
```

---

## 5. 练习题

### 练习 2.1: 使用 WaitGroup 等待多个 Goroutine
**目标**: 启动 5 个 Goroutine，每个打印自己的 ID，然后用 WaitGroup 等待所有 Goroutine 完成。

**验收标准**:
- 正确使用 `wg.Add()`、`wg.Done()`、`wg.Wait()`
- `wg.Add()` 在 Goroutine 外部调用
- 所有 5 个 Goroutine 都正常执行
- 主 Goroutine 正确等待

### 练习 2.2: 理解 sync.Once（概念题）
**问题**:
1. `sync.Once` 是如何保证只执行一次的？
2. 什么时候用 `sync.Map`，什么时候用 `RWMutex+map`？

**验收标准**: 能够清晰解释这两个问题。

---

查看 [code/sync_atomic.go](./code/sync_atomic.go) 中的 `DemoSyncOther()` 函数来运行示例代码。
