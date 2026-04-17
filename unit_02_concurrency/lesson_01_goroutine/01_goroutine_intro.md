# Section 1: Goroutine 基础与底层原理

## 1. 知识点核心说明

### 什么是 Goroutine
Goroutine 是 Go 语言中轻量级的线程实现，由 Go 运行时（runtime）调度，而非操作系统调度。

### `go` 关键字
使用 `go` 关键字即可启动一个新的 Goroutine：

```go
go func() {
    fmt.Println("Hello from goroutine!")
}()
```

### M-P-G 模型初体验
Go 采用独特的 M-P-G 调度模型：
- **M (Machine)**: 操作系统线程
- **P (Processor)**: 处理器，包含执行 Go 代码所需的资源
- **G (Goroutine)**: Goroutine，Go 语言的协程

核心思想：多个 Goroutine (G) 在少量的操作系统线程 (M) 上通过 P 进行调度，实现高效的并发。

---

## 2. Java 与 Go 的对比说明

| 特性 | Java Thread | Go Goroutine |
|------|-------------|--------------|
| 调度者 | 操作系统内核 | Go Runtime (用户态) |
| 初始栈大小 | 通常 1MB | 仅 2KB |
| 栈增长方式 | 固定大小 | 动态扩容/收缩 |
| 上下文切换 | 内核态切换（开销大） | 用户态切换（开销极小） |
| 内存占用 | 每个线程 ~1MB+ | 每个 Goroutine 仅几 KB |
| 并发量级 | 数千线程已很吃力 | 轻松支持数十万 Goroutine |

### 代码对比

**Java (Thread):**
```java
new Thread(() -&gt; {
    System.out.println("Hello from Java thread!");
}).start();
```

**Go (Goroutine):**
```go
go func() {
    fmt.Println("Hello from goroutine!")
}()
```

---

## 3. 可运行代码示例

### 错误写法：忘记等待 Goroutine 完成

```go
package main

import "fmt"

func main() {
    go func() {
        fmt.Println("Hello from goroutine!")
    }()
    // 主 Goroutine 立即退出，子 Goroutine 可能来不及执行
}
```

### 正确写法：使用 WaitGroup 等待

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(1)

    go func() {
        defer wg.Done()
        fmt.Println("Hello from goroutine!")
    }()

    wg.Wait() // 等待 Goroutine 完成
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1：Goroutine 泄露**
- 如果 Goroutine 因为等待 Channel、锁等永远阻塞，就会造成泄露
- 解决方案：使用 Context 进行级联取消，或者确保有退出机制

⚠️ **坑点 2：过度使用 Goroutine**
- 虽然 Goroutine 很轻量，但也不是越多越好
- 过多的 Goroutine 会增加调度开销和内存占用
- 考虑使用工作池 (Worker Pool) 来限制并发数

⚠️ **坑点 3：在 Goroutine 中使用循环变量**
```go
// 错误写法：所有 Goroutine 共享同一个变量 i
for i := 0; i &lt; 10; i++ {
    go func() {
        fmt.Println(i) // 可能输出相同的值
    }()
}

// 正确写法：每次迭代创建一个新变量
for i := 0; i &lt; 10; i++ {
    i := i // 或者使用参数传递
    go func() {
        fmt.Println(i)
    }()
}
```

---

## 5. 练习题

### 练习 1.1: 启动多个 Goroutine
**目标**: 启动 10 个 Goroutine，每个打印自己的编号（0-9）。

**验收标准**:
- 所有 10 个编号都被打印
- 使用 sync.WaitGroup 确保所有 Goroutine 都完成

### 练习 1.2: 对比 Goroutine 和 Thread（概念题）
**问题**: 
1. 为什么 Goroutine 的初始栈大小可以设为 2KB，而 Java Thread 需要 1MB？
2. 上下文切换在用户态和内核态有什么区别？

**验收标准**: 能够用自己的话解释清楚这两个问题。

---

查看 [goroutine.go](./goroutine.go) 中的 `DemoGoroutineBasics()` 函数来运行示例代码。
