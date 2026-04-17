# Section 3: Goroutine 泄露初体验

## 1. 知识点核心说明

### 什么是 Goroutine 泄露
Goroutine 泄露是指一个 Goroutine 启动后，因为某种原因永远无法结束，导致其占用的资源（栈内存、文件描述符等）无法被回收。

### 常见的 Goroutine 泄露场景

1. **Channel 阻塞**: Goroutine 在等待一个永远不会收到数据的 Channel
2. **锁未释放**: Goroutine 持有锁后崩溃或忘记释放，导致其他 Goroutine 永远等待
3. **循环引用**: Goroutine 之间互相等待形成死锁
4. **缺少退出机制**: Goroutine 运行在无限循环中，没有收到退出信号

---

## 2. Java 与 Go 的对比说明

| 特性 | Java Thread 泄露 | Go Goroutine 泄露 |
|------|-----------------|-------------------|
| 内存影响 | 每个泄露线程占用 ~1MB+ 栈 | 每个泄露 Goroutine 占用 ~2KB+ 栈（初始） |
| 发现难度 | 线程数量有限，较易发现 | Goroutine 可以很多，较难发现 |
| 回收机制 | 线程结束后自动回收 | Goroutine 结束后自动回收 |
| 常见原因 | 线程池未关闭、ThreadLocal 未清理 | Channel 阻塞、缺少退出机制 |

### 代码对比

**Java (Thread 泄露):**
```java
// 错误示例：线程永远等待
new Thread(() -&gt; {
    CountDownLatch latch = new CountDownLatch(1);
    try {
        latch.await(); // 永远等待，没有人 countDown
    } catch (InterruptedException e) {
        Thread.currentThread().interrupt();
    }
}).start();
```

**Go (Goroutine 泄露):**
```go
// 错误示例：Goroutine 永远等待
ch := make(chan int)
go func() {
    val := &lt;-ch // 永远等待，没有人发送数据
}()
```

---

## 3. 可运行代码示例

### 错误写法：Channel 接收导致泄露

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("初始 Goroutine 数量: %d\n", runtime.NumGoroutine())

	// 启动 100 个会泄露的 Goroutine
	for i := 0; i &lt; 100; i++ {
		ch := make(chan int)
		go func() {
			&lt;-ch // 永远等待
		}()
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("泄露后 Goroutine 数量: %d\n", runtime.NumGoroutine())
}
```

### 正确写法：使用带缓冲的 Channel 或确保发送

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("初始 Goroutine 数量: %d\n", runtime.NumGoroutine())

	// 方案1: 使用带缓冲的 Channel
	for i := 0; i &lt; 10; i++ {
		ch := make(chan int, 1) // 缓冲为1
		go func() {
			&lt;-ch
		}()
		ch &lt;- 1 // 发送数据
	}

	// 方案2: 使用 Context 或 done channel
	done := make(chan struct{})
	go func() {
		select {
		case &lt;-done:
			return
		}
	}()
	close(done) // 发送关闭信号

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("修复后 Goroutine 数量: %d\n", runtime.NumGoroutine())
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 忽略返回值的 Channel 发送**
```go
// 错误写法
go func() {
    result := doSomething()
    ch &lt;- result // 如果没有人接收，会阻塞
}()

// 正确写法：使用带缓冲的 Channel 或确保有接收者
ch := make(chan Result, 1)
go func() {
    result := doSomething()
    select {
    case ch &lt;- result:
    default: // 非阻塞发送
    }
}()
```

⚠️ **坑点 2: 忘记关闭导致 range 阻塞**
```go
// 错误写法
ch := make(chan int)
go func() {
    for i := 0; i &lt; 10; i++ {
        ch &lt;- i
    }
    // 忘记 close(ch)
}()

for val := range ch { // 永远阻塞
    fmt.Println(val)
}

// 正确写法：记得关闭 Channel
go func() {
    defer close(ch)
    for i := 0; i &lt; 10; i++ {
        ch &lt;- i
    }
}()
```

⚠️ **坑点 3: 在 HTTP Handler 中启动 Goroutine 但不加控制**
```go
// 错误写法
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    go func() {
        // 一些后台工作，没有超时控制
        time.Sleep(time.Hour)
    }()
})

// 正确写法：使用 Context 控制
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()
    
    go func() {
        select {
        case &lt;-ctx.Done():
            return
        case &lt;-time.After(time.Hour):
            // 超时会被 ctx 取消
        }
    }()
})
```

---

## 5. 练习题

### 练习 3.1: 识别并修复 Goroutine 泄露
**目标**: 找出下面代码中的 Goroutine 泄露问题并修复。

```go
package main

import (
	"fmt"
	"time"
)

func processTask(taskID int) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("Result of task %d", taskID)
}

func main() {
	results := make([]string, 0)
	
	for i := 0; i &lt; 10; i++ {
		go func(taskID int) {
			result := processTask(taskID)
			results = append(results, result) // 这里有问题！
		}(i)
	}
	
	time.Sleep(1 * time.Second)
	fmt.Printf("Got %d results\n", len(results))
}
```

**验收标准**:
- 找出所有问题（不止一个！）
- 修复后的程序能正确收集所有 10 个结果
- 没有 Goroutine 泄露

### 练习 3.2: 实现一个可取消的 Goroutine
**目标**: 写一个 Goroutine，它能做一些工作，同时也能接收取消信号并安全退出。

**验收标准**:
- Goroutine 能正常执行工作
- 能通过外部信号取消 Goroutine
- 取消后 Goroutine 能安全退出，没有泄露

---

查看 [goroutine.go](./goroutine.go) 中的 `DemoGoroutineLeak()` 函数来了解更多示例。
