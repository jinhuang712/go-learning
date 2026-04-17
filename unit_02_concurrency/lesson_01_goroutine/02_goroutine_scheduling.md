# Section 2: Goroutine 调度与生命周期

## 1. 知识点核心说明

### Goroutine 生命周期

Goroutine 有以下几个状态：
1. **Gidle**: 刚刚分配，还未初始化
2. **Grunnable**: 在运行队列中，等待运行
3. **Grunning**: 正在运行中
4. **Gsyscall**: 正在执行系统调用
5. **Gwaiting**: 等待某个条件（如 Channel、锁等）
6. **Gdead**: 已结束，等待被复用

### GOMAXPROCS 的作用

`GOMAXPROCS` 决定了同时执行用户级 Go 代码的操作系统线程（M）的最大数量。

```go
import "runtime"

runtime.GOMAXPROCS(4) // 设置为 4
```

默认值是机器的 CPU 核心数。

### Goroutine 调度触发时机

Goroutine 调度会在以下情况发生：
- 函数调用时（尤其是栈扩容时）
- 系统调用后
- 显式调用 `runtime.Gosched()`
- 使用锁、Channel 等同步原语时

---

## 2. Java 与 Go 的对比说明

| 特性 | Java Thread Scheduling | Go Goroutine Scheduling |
|------|-----------------------|-------------------------|
| 调度器 | 操作系统内核调度器 | Go Runtime M-P-G 调度器 |
| 调度时机 | 由操作系统决定，不可控 | 在 Go 运行时中可控 |
| 线程数 | 通常受限于操作系统 | `GOMAXPROCS` 控制活跃 M 的数量 |
| 上下文切换 | 昂贵（内核态） | 廉价（用户态） |
| 优先级 | 支持线程优先级 | 不支持 Goroutine 优先级 |

### 线程数设置对比

**Java:**
```java
// Java 没有直接限制线程数的机制
// 通常通过线程池来管理
ExecutorService pool = Executors.newFixedThreadPool(4);
```

**Go:**
```go
import "runtime"

// 控制同时执行 Go 代码的 M 数量
runtime.GOMAXPROCS(4)
```

---

## 3. 可运行代码示例

### 示例 1: GOMAXPROCS 和 Goroutine 数量监控

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("默认 GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("初始 Goroutine 数量: %d\n", runtime.NumGoroutine())

	// 改变 GOMAXPROCS
	runtime.GOMAXPROCS(2)
	fmt.Printf("设置后 GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i &lt; 10; i++ {
		go func() {
			defer wg.Done()
			// 模拟一些工作
			for j := 0; j &lt; 1000000; j++ {
			}
		}()
	}

	fmt.Printf("启动后 Goroutine 数量: %d\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("结束后 Goroutine 数量: %d\n", runtime.NumGoroutine())
}
```

### 示例 2: 使用 runtime.Gosched() 主动让出 CPU

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i &lt; 5; i++ {
			fmt.Println("Goroutine 1:", i)
			runtime.Gosched() // 主动让出 CPU
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i &lt; 5; i++ {
			fmt.Println("Goroutine 2:", i)
			runtime.Gosched()
		}
	}()

	wg.Wait()
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 盲目修改 GOMAXPROCS**
- 默认值通常是最优的，不要随意修改
- 只有在特殊场景下（如容器环境 CPU 限制）才需要调整
- 在容器中运行时，注意 Go 1.19+ 能自动识别 cgroup CPU 限制

⚠️ **坑点 2: 过度使用 runtime.Gosched()**
- `runtime.Gosched()` 会让当前 Goroutine 让出 CPU，但不保证立即被调度回来
- 不要在性能关键路径上频繁调用
- 大多数情况下，使用 Channel 或 sync 包的同步原语更好

⚠️ **坑点 3: 用 Goroutine 数量衡量负载**
- Goroutine 数量多 ≠ 负载高
- 应该关注 CPU、内存、Goroutine 阻塞等待等指标
- 可以用 `pprof` 或 `runtime.NumGoroutine()` 监控，但不要仅依赖这一个指标

---

## 5. 练习题

### 练习 2.1: 观察 GOMAXPROCS 的影响
**目标**: 编写一个程序，分别在 GOMAXPROCS=1 和 GOMAXPROCS=4 的情况下运行多个 CPU 密集型 Goroutine，观察并对比运行时间。

**验收标准**:
- 程序能正确设置 GOMAXPROCS
- 能输出不同设置下的运行时间
- 能解释为什么会有这样的差异

### 练习 2.2: 实现一个简单的 Goroutine 池
**目标**: 用固定数量的 Goroutine 来处理一批任务，避免创建过多 Goroutine。

**验收标准**:
- Goroutine 数量可配置
- 任务可以通过某种方式提交
- 所有任务都能被处理
- 使用 sync.WaitGroup 等待所有任务完成

---

查看 [goroutine.go](./goroutine.go) 中的 `DemoGoroutineScheduling()` 函数来运行示例代码。
