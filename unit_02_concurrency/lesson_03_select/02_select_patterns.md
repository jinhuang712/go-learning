# Section 2: Select 模式

## 1. 知识点核心说明

Select 不仅可以用于多路等待，还可以用来实现各种实用的并发模式。本节介绍几种最常用的模式。

### 模式 1: 超时控制
使用 `time.After` 来实现超时控制。

### 模式 2: 非阻塞读写
使用 `default` case 来实现非阻塞的 Channel 操作。

### 模式 3: 多路等待
同时等待多个 Channel，哪个先就绪就处理哪个。

---

## 2. Java 与 Go 的对比说明

| 模式 | Java | Go |
|------|------|----|
| **超时控制** | `CompletableFuture.orTimeout()` / `Future.get(timeout)` | `select` + `time.After` |
| **非阻塞操作** | `BlockingQueue.poll()` / `offer()` | `select` + `default` |
| **多路等待** | `CompletableFuture.anyOf()` | `select` 多个 case |

### 代码对照

**Java (超时控制)**:
```java
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeUnit;

public class TimeoutExample {
    public static void main(String[] args) {
        CompletableFuture&lt;String&gt; future = CompletableFuture.supplyAsync(() -&gt; {
            try {
                Thread.sleep(2000);
                return "任务完成";
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                return "任务中断";
            }
        });
        
        try {
            String result = future.orTimeout(1, TimeUnit.SECONDS).get();
            System.out.println("收到结果: " + result);
        } catch (Exception e) {
            System.out.println("超时或异常: " + e.getMessage());
        }
    }
}
```

**Go (超时控制)**:
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch &lt;- "任务完成"
	}()

	select {
	case result := &lt;-ch:
		fmt.Println("收到结果:", result)
	case &lt;-time.After(1 * time.Second):
		fmt.Println("超时")
	}
}
```

**心智迁移要点**:
- Java: 使用 `CompletableFuture` 或显式的 `Future.get(timeout)`
- Go: 使用 `select` + `time.After`，更简洁
- Java: 超时会抛出异常
- Go: 超时是正常的控制流，通过 case 处理
- Java: `CompletableFuture` 功能强大但复杂
- Go: `select` 简单但灵活，可以组合多种模式

---

## 3. 可运行代码示例

### 模式 1: 超时控制

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	resultCh := make(chan string, 1)

	// 模拟一个慢查询
	go func() {
		time.Sleep(2 * time.Second)
		resultCh &lt;- "查询结果: 用户数据"
	}()

	// 等待结果，但最多等 1 秒
	select {
	case result := &lt;-resultCh:
		fmt.Println("成功:", result)
	case &lt;-time.After(1 * time.Second):
		fmt.Println("错误: 查询超时")
	}
}
```

### 模式 2: 非阻塞读写

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// 非阻塞发送
	select {
	case ch &lt;- 42:
		fmt.Println("发送成功")
	default:
		fmt.Println("发送失败（缓冲区满）")
	}

	// 非阻塞接收
	select {
	case val := &lt;-ch:
		fmt.Println("接收成功:", val)
	default:
		fmt.Println("接收失败（没有数据）")
	}

	// 再次非阻塞接收（这次应该有数据）
	select {
	case val := &lt;-ch:
		fmt.Println("接收成功:", val)
	default:
		fmt.Println("接收失败（没有数据）")
	}
}
```

### 模式 3: 多路等待

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	stop := make(chan struct{})

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 &lt;- "来自服务 A"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 &lt;- "来自服务 B"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		close(stop)
	}()

	for {
		select {
		case msg := &lt;-ch1:
			fmt.Println("收到:", msg)
		case msg := &lt;-ch2:
			fmt.Println("收到:", msg)
		case &lt;-stop:
			fmt.Println("收到停止信号，退出")
			return
		}
	}
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 在循环中使用 time.After 导致内存泄露**
```go
// 错误写法：每次循环创建新的 Timer，且不会被 GC 回收！
for {
	select {
	case &lt;-dataCh:
		// 处理数据
	case &lt;-time.After(5 * time.Second):
		// 超时
	}
}

// 正确写法：复用 Timer
timer := time.NewTimer(5 * time.Second)
defer timer.Stop()

for {
	select {
	case &lt;-dataCh:
		// 处理数据
		timer.Reset(5 * time.Second)
	case &lt;-timer.C:
		// 超时
		return
	}
}
```

⚠️ **坑点 2: 非阻塞操作的误用**
```go
// 不推荐：过度使用非阻塞，可能导致忙等待
for {
	select {
	case &lt;-ch:
		// 处理数据
	default:
		// 没有数据，继续循环
		// 这会导致 CPU 100%！
	}
}

// 推荐：使用阻塞，或者添加 sleep
for {
	select {
	case &lt;-ch:
		// 处理数据
	default:
		time.Sleep(10 * time.Millisecond) // 短暂 sleep，避免忙等待
	}
}
```

⚠️ **坑点 3: 超时控制的精度问题**
```go
// time.After 的精度受系统调度影响，可能不是精确的 5 秒
select {
case &lt;-time.After(5 * time.Second):
    // 可能在 5.001 秒或 5.1 秒后触发
}
```

---

## 5. 练习题

### 练习 2.1: 实现多路超时控制
**目标**: 写一个函数，同时请求 3 个服务，等待所有结果，但最多等待 3 秒。如果某个服务超时或失败，其他服务的结果仍然可用。

**验收标准**:
- 使用 `select` 多路等待
- 有整体超时控制
- 能处理部分服务成功、部分失败的情况
- 没有内存泄露

### 练习 2.2: 理解非阻塞操作（概念题）
**问题**:
1. 什么时候应该使用非阻塞 Channel 操作？什么时候应该使用阻塞？
2. `time.After` 在循环中为什么会导致内存泄露？如何避免？

**验收标准**: 能够清晰解释这两个问题。

---

查看 [code/select.go](./code/select.go) 来了解更多示例，或查看 `DemoSelectPatterns()` 函数。
