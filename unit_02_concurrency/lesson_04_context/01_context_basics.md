# Section 1: Context 基础

## 1. 知识点核心说明

### 什么是 Context
Context 是 Go 语言中用于在 Goroutine 之间传递请求作用域数据、取消信号、超时信号等的标准接口。

### Context 接口
```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() &lt;-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

### 根 Context
- `context.Background()`: 返回一个空的 Context，通常用作根 Context
- `context.TODO()`: 返回一个空的 Context，用于不确定使用什么 Context 时

---

## 2. Java 与 Go 的对比说明

| 特性 | Java | Go |
|------|------|----|
| **请求上下文** | `ThreadLocal` + 手动传递 | `context.Context` + 显式传递 |
| **取消机制** | `Future.cancel()` / `Thread.interrupt()` | `context.WithCancel()` + `ctx.Done()` |
| **超时机制** | `Future.get(timeout)` | `context.WithTimeout()` / `WithDeadline()` |
| **数据传递** | `ThreadLocal` | `context.WithValue()` |

### 为什么 Go 需要 Context？
Java 中没有内置的 Context 概念，通常使用：
- `ThreadLocal` 来传递请求数据
- `Future` 来处理取消和超时
- 但 ThreadLocal 在异步编程中难以传递，且容易内存泄露

Go 的 Context 提供了：
- 显式传递，更清晰
- 支持级联取消
- 标准库广泛使用（HTTP、数据库、gRPC 等）

### 代码对照

**Java (ThreadLocal + Future)**:
```java
import java.util.concurrent.*;

public class JavaContextExample {
    private static final ThreadLocal&lt;String&gt; USER_ID = new ThreadLocal&lt;&gt;();
    
    public static void main(String[] args) throws Exception {
        ExecutorService executor = Executors.newSingleThreadExecutor();
        
        USER_ID.set("user-123");
        
        Future&lt;String&gt; future = executor.submit(() -&gt; {
            // ThreadLocal 在线程池中需要小心处理！
            String userId = USER_ID.get(); // 可能是 null！
            return "处理用户: " + userId;
        });
        
        try {
            String result = future.get(5, TimeUnit.SECONDS);
            System.out.println(result);
        } catch (TimeoutException e) {
            future.cancel(true);
            System.out.println("超时");
        }
        
        executor.shutdown();
    }
}
```

**Go (Context)**:
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	
	// 设置用户 ID
	type userIDKey struct{}
	ctx = context.WithValue(ctx, userIDKey{}, "user-123")
	
	// 设置超时
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	done := make(chan string, 1)
	go func() {
		// Context 显式传递，清晰明了
		userID := ctx.Value(userIDKey{}).(string)
		done &lt;- fmt.Sprintf("处理用户: %s", userID)
	}()
	
	select {
	case result := &lt;-done:
		fmt.Println(result)
	case &lt;-ctx.Done():
		fmt.Println("超时或取消:", ctx.Err())
	}
}
```

**心智迁移要点**:
- Java: 隐式通过 ThreadLocal 传递，在线程池中容易出问题
- Go: 显式传递 Context，每个函数接收 Context 作为第一个参数
- Java: 取消和超时是分开的机制
- Go: Context 统一处理取消、超时、数据传递
- Java: ThreadLocal 容易内存泄露
- Go: Context 不会有内存泄露（正确使用的情况下）

---

## 3. 可运行代码示例

### 错误写法：忽略 Context

```go
// 错误写法：函数不接收 Context，无法取消或超时
func doWork() {
	for {
		// 永远循环，无法取消！
		time.Sleep(1 * time.Second)
		fmt.Println("工作中...")
	}
}

func main() {
	go doWork()
	time.Sleep(3 * time.Second)
	// 无法取消 doWork！
}
```

### 正确写法：使用 Context

```go
// 正确写法：函数接收 Context，监听 Done()
func doWork(ctx context.Context) {
	for {
		select {
		case &lt;-ctx.Done():
			fmt.Println("收到取消信号，退出")
			return
		case &lt;-time.After(1 * time.Second):
			fmt.Println("工作中...")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	go doWork(ctx)
	
	time.Sleep(3 * time.Second)
	fmt.Println("发送取消信号")
	cancel()
	time.Sleep(1 * time.Second)
}
```

### Context.Background() vs context.TODO()

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	// Background() 通常用作根 Context
	ctx1 := context.Background()
	fmt.Println("Background():", ctx1)
	
	// TODO() 用于不确定用什么 Context 时
	ctx2 := context.TODO()
	fmt.Println("TODO():", ctx2)
	
	// 两者都是空 Context，但语义不同
	// - Background(): 明确是根 Context
	// - TODO(): 临时使用，后续应该替换为合适的 Context
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: Context 不作为第一个参数**
```go
// 错误写法：Context 不是第一个参数
func doWork(userID string, ctx context.Context) { /* ... */ }

// 正确写法：Context 总是第一个参数
func doWork(ctx context.Context, userID string) { /* ... */ }
```

⚠️ **坑点 2: 存储 Context 在结构体中**
```go
// 错误写法：把 Context 存在结构体里
type Worker struct {
    ctx context.Context  // 不要这样做！
}

// 正确写法：每次调用方法时传递 Context
type Worker struct{}

func (w *Worker) DoWork(ctx context.Context) { /* ... */ }
```

⚠️ **坑点 3: 忘记调用 cancel()**
```go
// 错误写法：WithCancel/WithTimeout 后忘记 cancel
ctx, _ := context.WithTimeout(ctx, 5*time.Second)
// 忘记 defer cancel()！

// 正确写法：总是 defer cancel()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel() // 确保资源释放
```

---

## 5. 练习题

### 练习 1.1: 实现可取消的工作
**目标**: 写一个函数，接收 Context，做一些工作（比如循环打印），并能在 Context 取消时优雅退出。

**验收标准**:
- 函数接收 Context 作为第一个参数
- 监听 `ctx.Done()` 来检测取消
- 收到取消信号时能优雅退出
- 使用 `context.WithCancel()` 来测试取消功能

### 练习 1.2: 理解 Context 的设计（概念题）
**问题**:
1. Go 的 Context 和 Java 的 ThreadLocal 有什么本质区别？
2. 为什么 Go 选择显式传递 Context，而不是像 Java 那样隐式传递？

**验收标准**: 能够清晰解释这两个问题。

---

查看 [code/context.go](./code/context.go) 中的 `DemoContextBasics()` 函数来运行示例代码。
