# Section 1: Channel 基础

## 1. 知识点核心说明

### 什么是 Channel
Channel 是 Go 语言中用于 Goroutine 之间通信和同步的核心原语，是 CSP (Communicating Sequential Processes) 并发模型的实现。

**核心思想**: 不要通过共享内存来通信，而要通过通信来共享内存。

### Channel 的类型

| 类型 | 语法 | 说明 |
|------|------|------|
| 无缓冲 Channel | `make(chan T)` | 发送和接收必须同时就绪，否则阻塞 |
| 有缓冲 Channel | `make(chan T, n)` | 缓冲区未满时发送不阻塞，缓冲区非空时接收不阻塞 |

### 基本操作

```go
ch := make(chan int)

// 发送
ch &lt;- 42

// 接收
val := &lt;-ch

// 关闭
close(ch)

// 遍历（必须关闭后才能用 range 遍历）
for v := range ch {
    fmt.Println(v)
}
```

---

## 2. Java 与 Go 的对比说明

| 特性 | Java | Go |
|------|------|----|
| **通信原语** | `BlockingQueue` | Channel |
| **同步方式** | 共享内存 + 锁 | CSP 模型（通过通信同步） |
| **无缓冲** | `SynchronousQueue` | 无缓冲 Channel |
| **有缓冲** | `ArrayBlockingQueue` / `LinkedBlockingQueue` | 有缓冲 Channel |
| **阻塞行为** | `put()` / `take()` 阻塞 | 发送/接收操作阻塞 |

### 代码对照

**Java (BlockingQueue)**:
```java
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.SynchronousQueue;

public class Main {
    public static void main(String[] args) throws InterruptedException {
        BlockingQueue&lt;Integer&gt; queue = new SynchronousQueue&lt;&gt;();
        
        new Thread(() -&gt; {
            try {
                System.out.println("Thread: 发送数据 42");
                queue.put(42);
                System.out.println("Thread: 发送完成");
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }).start();
        
        System.out.println("Main: 等待接收...");
        Integer val = queue.take();
        System.out.println("Main: 接收到 " + val);
    }
}
```

**Go (Channel)**:
```go
package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        fmt.Println("Goroutine: 发送数据 42")
        ch &lt;- 42
        fmt.Println("Goroutine: 发送完成")
    }()

    fmt.Println("Main: 等待接收...")
    val := &lt;-ch
    fmt.Printf("Main: 接收到 %d\n", val)
}
```

**心智迁移要点**:
- Java: 共享内存是默认方式，队列只是辅助
- Go: Channel 是并发通信的核心方式，优先考虑 Channel
- Java: 异常处理 InterruptedException
- Go: Channel 操作不会抛异常，阻塞是正常行为

---

## 3. 可运行代码示例

### 错误写法：从已关闭的 Channel 接收

```go
// 错误写法
ch := make(chan int)
close(ch)
val := &lt;-ch  // 会接收到零值，不会 panic，但可能不是你想要的
fmt.Println(val)  // 输出 0
```

### 正确写法：检查 Channel 是否已关闭

```go
// 正确写法：使用 comma-ok 惯用法
ch := make(chan int, 1)
ch &lt;- 42
close(ch)

val, ok := &lt;-ch
if ok {
    fmt.Println("接收到:", val)
} else {
    fmt.Println("Channel 已关闭")
}

// 或者使用 range（会自动处理关闭）
ch2 := make(chan int, 3)
ch2 &lt;- 1
ch2 &lt;- 2
ch2 &lt;- 3
close(ch2)

for v := range ch2 {
    fmt.Println("遍历收到:", v)
}
```

### 无缓冲 vs 有缓冲的区别

```go
// 无缓冲 Channel
unbuffered := make(chan int)
go func() {
    unbuffered &lt;- 1  // 会阻塞，直到有人接收
}()

// 有缓冲 Channel
buffered := make(chan int, 2)
buffered &lt;- 1  // 不会阻塞
buffered &lt;- 2  // 不会阻塞
// buffered &lt;- 3  // 会阻塞，缓冲区已满
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 向已关闭的 Channel 发送**
```go
ch := make(chan int)
close(ch)
ch &lt;- 1  // panic! 不要这样做！
```

⚠️ **坑点 2: 忘记关闭 Channel 导致 range 永远阻塞**
```go
// 错误写法
ch := make(chan int)
go func() {
    ch &lt;- 1
    ch &lt;- 2
    // 忘记 close(ch)
}()

for v := range ch {  // 永远阻塞！
    fmt.Println(v)
}
```

⚠️ **坑点 3: nil Channel 的行为**
```go
var ch chan int  // nil Channel

ch &lt;- 1  // 永远阻塞！
&lt;-ch     // 永远阻塞！

// nil Channel 在 select 中很有用
select {
case &lt;-ch:  // nil Channel，这个 case 永远不会被选中
default:
}
```

⚠️ **坑点 4: 关闭 Channel 多次**
```go
ch := make(chan int)
close(ch)
close(ch)  // panic! 不要这样做！
```

---

## 5. 练习题

### 练习 1.1: 交替打印
**目标**: 使用两个 Goroutine 和两个无缓冲 Channel，实现交替打印 "A" 和 "B"，各打印 5 次，输出应该是 "ABABABABAB"。

**验收标准**:
- 使用无缓冲 Channel 同步
- 输出顺序正确
- 所有资源正确释放

### 练习 1.2: 理解 Channel 阻塞（概念题）
**问题**:
1. 无缓冲 Channel 和有缓冲 Channel 在发送和接收行为上有什么区别？
2. 为什么 Go 说"不要通过共享内存来通信，而要通过通信来共享内存"？

**验收标准**: 能够清晰解释这两个问题。

---

查看 [code/channel.go](./code/channel.go) 中的 `DemoChannelBasics()` 函数来运行示例代码。
