# Section 2: Channel 常用模式

## 1. 知识点核心说明

Channel 不仅仅是通信工具，还可以用来实现各种常用的并发模式。本节介绍几种最常用的模式。

### 1. 生产者-消费者模式
生产者发送数据到 Channel，消费者从 Channel 接收数据处理。

### 2. 扇入 (Fan-in)
多个 Goroutine 发送数据到同一个 Channel。

### 3. 扇出 (Fan-out)
一个 Goroutine 发送数据到多个 Channel，或者多个 Goroutine 从同一个 Channel 接收数据。

### 4. 信号量模式
使用有缓冲 Channel 来限制并发数。

---

## 2. Java 与 Go 的对比说明

| 模式 | Java | Go |
|------|------|----|
| **生产者-消费者** | `BlockingQueue` + 线程 | Channel + Goroutine |
| **工作池** | `ExecutorService` | Goroutine + Channel（通常不需要池） |
| **信号量** | `Semaphore` | 有缓冲 Channel |
| **Future** | `CompletableFuture` | Channel |

### 代码对照

**Java (工作池 - ExecutorService)**:
```java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class WorkerPool {
    public static void main(String[] args) {
        ExecutorService pool = Executors.newFixedThreadPool(5);
        
        for (int i = 0; i &lt; 100; i++) {
            final int taskId = i;
            pool.submit(() -&gt; {
                System.out.println("Processing task " + taskId);
            });
        }
        
        pool.shutdown();
    }
}
```

**Go (信号量模式 - 限制并发)**:
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	const maxConcurrent = 5
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	for i := 0; i &lt; 100; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()
			sem &lt;- struct{}{} // 获取信号量
			defer func() { &lt;-sem }() // 释放信号量

			fmt.Printf("Processing task %d\n", taskID)
		}(i)
	}

	wg.Wait()
}
```

**心智迁移要点**:
- Java: 通常使用线程池来管理并发
- Go: Goroutine 足够轻量，通常不需要池；需要限制并发时用信号量模式
- Java: `ExecutorService` 是重量级的
- Go: 用 Channel 实现的信号量模式非常轻量

---

## 3. 可运行代码示例

### 模式 1: 生产者-消费者

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan&lt;- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i &lt; 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch &lt;- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch &lt;-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("Consuming: %d\n", val)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go producer(ch, &amp;wg)
	go consumer(ch, &amp;wg)

	wg.Wait()
}
```

### 模式 2: 扇入 (Fan-in)

```go
package main

import "fmt"

func producer(id int, ch chan&lt;- string) {
	for i := 0; i &lt; 3; i++ {
		ch &lt;- fmt.Sprintf("Producer %d: message %d", id, i)
	}
}

func fanIn(ch1, ch2 &lt;-chan string, out chan&lt;- string) {
	for {
		select {
		case msg := &lt;-ch1:
			out &lt;- msg
		case msg := &lt;-ch2:
			out &lt;- msg
		}
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	out := make(chan string)

	go producer(1, ch1)
	go producer(2, ch2)
	go fanIn(ch1, ch2, out)

	for i := 0; i &lt; 6; i++ {
		fmt.Println(&lt;-out)
	}
}
```

### 模式 3: 信号量（限制并发数）

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const maxConcurrent = 3
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	for i := 0; i &lt; 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			sem &lt;- struct{}{} // 获取
			defer func() { &lt;-sem }() // 释放
			
			fmt.Printf("Worker %d: 开始工作\n", id)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Worker %d: 完成工作\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("所有工作完成")
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 忘记关闭 Channel 导致消费者永远阻塞**
```go
// 错误写法
ch := make(chan int)
go func() {
    ch &lt;- 1
    ch &lt;- 2
    // 忘记 close(ch)
}()

for val := range ch { // 永远阻塞！
    fmt.Println(val)
}
```

⚠️ **坑点 2: 信号量获取和释放不配对**
```go
// 错误写法
go func() {
    sem &lt;- struct{}{}
    if someCondition {
        return // 提前返回，没有释放信号量！
    }
    &lt;-sem
}()

// 正确写法：用 defer 确保释放
go func() {
    sem &lt;- struct{}{}
    defer func() { &lt;-sem }()
    if someCondition {
        return
    }
}()
```

⚠️ **坑点 3: 扇出时启动过多 Goroutine**
```go
// 虽然 Goroutine 很轻量，但也不是无限的
// 如果任务数是 100 万，直接启动 100 万个 Goroutine 可能会有问题

// 更好的做法：结合信号量模式限制并发
```

---

## 5. 练习题

### 练习 2.1: 实现工作池
**目标**: 实现一个固定大小的工作池，使用 Channel 来分发任务给多个 Worker Goroutine。

**验收标准**:
- Worker 数量可配置
- 任务可以通过 Channel 提交
- 所有任务都能被处理
- 正确使用 sync.WaitGroup 等待完成

### 练习 2.2: 理解扇入扇出（概念题）
**问题**:
1. 什么是扇入？什么是扇出？
2. 在什么场景下应该使用扇入/扇出模式？

**验收标准**: 能够清晰解释这两个问题。

---

查看 [code/channel.go](./code/channel.go) 来了解更多示例，或查看 `DemoChannelPatterns()` 函数。
