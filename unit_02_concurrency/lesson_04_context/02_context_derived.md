# Section 2: 派生 Context

## 1. 知识点核心说明

Context 是不可变的，需要通过派生（Derive）来创建新的 Context。

### 派生方式

| 函数 | 说明 |
|------|------|
| `context.WithCancel(parent)` | 创建可取消的 Context |
| `context.WithTimeout(parent, timeout)` | 创建带超时的 Context |
| `context.WithDeadline(parent, deadline)` | 创建带截止时间的 Context |
| `context.WithValue(parent, key, value)` | 创建带键值对的 Context |

### 级联取消
当父 Context 取消时，所有由它派生的子 Context 也会被取消。

---

## 2. Java 与 Go 的对比说明

| 特性 | Java | Go |
|------|------|----|
| **取消** | `Future.cancel()` / `Thread.interrupt()` | `context.WithCancel()` + `ctx.Done()` |
| **超时** | `Future.get(timeout)` | `context.WithTimeout()` / `WithDeadline()` |
| **数据传递** | `ThreadLocal` | `context.WithValue()` |
| **级联取消** | 需要手动实现 | 自动支持 |

### ThreadLocal vs WithValue 的关键区别

| 特性 | ThreadLocal | WithValue |
|------|-------------|-----------|
| **存储位置** | 线程本地 | Context 树中 |
| **异步传递** | 困难（线程池中易丢失） | 容易（显式传递） |
| **内存泄露** | 风险高 | 风险低（随 Context 回收） |
| **作用域** | 线程级 | 请求级 |

### 代码对照

**Java (ThreadLocal + Future)**:
```java
import java.util.concurrent.*;

public class JavaDerivedContext {
    private static final ThreadLocal&lt;String&gt; USER_ID = new ThreadLocal&lt;&gt;();
    private static final ThreadLocal&lt;String&gt; TRACE_ID = new ThreadLocal&lt;&gt;();
    
    public static void main(String[] args) throws Exception {
        ExecutorService executor = Executors.newFixedThreadPool(3);
        
        USER_ID.set("user-123");
        TRACE_ID.set("trace-456");
        
        // 超时控制
        Future&lt;String&gt; future = executor.submit(() -&gt; {
            // 在线程池中，ThreadLocal 需要手动传递！
            String userId = USER_ID.get(); // 可能是 null！
            String traceId = TRACE_ID.get(); // 可能是 null！
            
            try {
                Thread.sleep(2000);
                return "处理完成: " + userId;
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                return "被中断";
            }
        });
        
        try {
            String result = future.get(1, TimeUnit.SECONDS);
            System.out.println(result);
        } catch (TimeoutException e) {
            future.cancel(true);
            System.out.println("超时");
        }
        
        executor.shutdown();
    }
}
```

**Go (派生 Context)**:
```go
package main

import (
	"context"
	"fmt"
	"time"
)

type userIDKey struct{}
type traceIDKey struct{}

func main() {
	ctx := context.Background()
	
	// 设置值
	ctx = context.WithValue(ctx, userIDKey{}, "user-123")
	ctx = context.WithValue(ctx, traceIDKey{}, "trace-456")
	
	// 设置超时
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	
	done := make(chan string, 1)
	go func() {
		// Context 显式传递，值不会丢失
		userID := ctx.Value(userIDKey{}).(string)
		traceID := ctx.Value(traceIDKey{}).(string)
		
		select {
		case &lt;-time.After(2 * time.Second):
			done &lt;- fmt.Sprintf("处理完成: %s (trace: %s)", userID, traceID)
		case &lt;-ctx.Done():
			done &lt;- fmt.Sprintf("被取消或超时: %v", ctx.Err())
		}
	}()
	
	select {
	case result := &lt;-done:
		fmt.Println(result)
	case &lt;-ctx.Done():
		fmt.Println("主 Context 超时或取消")
	}
}
```

**心智迁移要点**:
- Java: ThreadLocal 在异步编程中容易丢失，需要手动传递
- Go: Context 显式传递，值不会丢失
- Java: 超时和取消需要通过 Future 处理
- Go: Context 统一处理超时和取消
- Java: 级联取消需要手动实现
- Go: Context 自动支持级联取消
- Java: ThreadLocal 容易内存泄露
- Go: WithValue 随 Context 回收，不会泄露

---

## 3. 可运行代码示例

### 示例 1: WithCancel 级联取消

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case &lt;-ctx.Done():
			fmt.Printf("Worker %d: 收到取消信号，退出\n", id)
			return
		case &lt;-time.After(500 * time.Millisecond):
			fmt.Printf("Worker %d: 工作中...\n", id)
		}
	}
}

func main() {
	ctx := context.Background()
	
	// 创建可取消的父 Context
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	
	// 启动多个 Worker
	for i := 1; i &lt;= 3; i++ {
		go worker(ctx, i)
	}
	
	// 运行一会儿
	time.Sleep(2 * time.Second)
	
	// 取消父 Context，所有子 Context 都会被取消！
	fmt.Println("发送取消信号...")
	cancel()
	
	// 等待一下，让 Worker 退出
	time.Sleep(1 * time.Second)
}
```

### 示例 2: WithTimeout / WithDeadline

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context) string {
	select {
	case &lt;-time.After(2 * time.Second):
		return "操作完成"
	case &lt;-ctx.Done():
		return fmt.Sprintf("操作被取消: %v", ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	
	// WithTimeout: 相对超时
	fmt.Println("--- WithTimeout ---")
	ctx1, cancel1 := context.WithTimeout(ctx, 1*time.Second)
	defer cancel1()
	fmt.Println(slowOperation(ctx1))
	
	// WithDeadline: 绝对截止时间
	fmt.Println("\n--- WithDeadline ---")
	deadline := time.Now().Add(1 * time.Second)
	ctx2, cancel2 := context.WithDeadline(ctx, deadline)
	defer cancel2()
	fmt.Println(slowOperation(ctx2))
}
```

### 示例 3: WithValue

```go
package main

import (
	"context"
	"fmt"
)

// 定义 key 类型，避免与其他包冲突
type userKey struct{}
type traceKey struct{}

func handleRequest(ctx context.Context) {
	// 从 Context 中获取值
	userID := ctx.Value(userKey{})
	traceID := ctx.Value(traceKey{})
	
	fmt.Printf("处理请求: user=%v, trace=%v\n", userID, traceID)
	
	// 派生新 Context 传递给下游
	downstream(ctx)
}

func downstream(ctx context.Context) {
	// 下游也能获取到值
	userID := ctx.Value(userKey{})
	fmt.Printf("下游: user=%v\n", userID)
}

func main() {
	ctx := context.Background()
	
	// 设置值
	ctx = context.WithValue(ctx, userKey{}, "user-123")
	ctx = context.WithValue(ctx, traceKey{}, "trace-456")
	
	// 处理请求
	handleRequest(ctx)
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: WithValue 使用 string 作为 key**
```go
// 错误写法：使用 string 作为 key，容易与其他包冲突
ctx = context.WithValue(ctx, "user-id", "123")

// 正确写法：使用自定义类型作为 key
type userIDKey struct{}
ctx = context.WithValue(ctx, userIDKey{}, "123")
```

⚠️ **坑点 2: 忘记 defer cancel()**
```go
// 错误写法：忘记 cancel，可能导致资源泄露
ctx, _ := context.WithTimeout(parent, 5*time.Second)

// 正确写法：总是 defer cancel()
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()
```

⚠️ **坑点 3: 在 WithValue 中存储大对象**
```go
// 不推荐：WithValue 应该只存储小的请求作用域数据
type BigData struct { /* 很多字段 */ }
ctx = context.WithValue(ctx, "data", &amp;BigData{}) // 不推荐

// 推荐：只存储小的元数据
ctx = context.WithValue(ctx, userIDKey{}, "user-123")
ctx = context.WithValue(ctx, traceKey{}, "trace-456")
```

---

## 5. 练习题

### 练习 2.1: 实现级联取消
**目标**: 创建一个父 Context，派生出 3 个子 Context，每个子 Context 启动一个 Goroutine。当取消父 Context 时，所有子 Goroutine 都应该能收到取消信号并退出。

**验收标准**:
- 使用 `context.WithCancel()`
- 派生出多个子 Context
- 每个子 Goroutine 监听 `ctx.Done()`
- 取消父 Context 时，所有子 Goroutine 都能退出
- 正确使用 `defer cancel()`

### 练习 2.2: 理解 WithValue（概念题）
**问题**:
1. 为什么 WithValue 的 key 应该使用自定义类型，而不是 string？
2. ThreadLocal 和 WithValue 在设计上有什么本质区别？各自适用于什么场景？

**验收标准**: 能够清晰解释这两个问题。

---

查看 [code/context.go](./code/context.go) 来了解更多示例，或查看 `DemoContextDerived()` 函数。
