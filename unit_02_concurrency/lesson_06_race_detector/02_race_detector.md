# Section 2: Race Detector 使用

## 1. 知识点核心说明

### 启用 Race Detector
使用 `-race` 标志来启用 Race Detector：

```bash
# 运行程序
go run -race main.go

# 运行测试
go test -race ./...

# 编译程序（包含 race detector）
go build -race
```

### Race Detector 输出解读
Race Detector 会输出：
- 竞争的内存地址
- 发生竞争的 Goroutine
- 发生竞争的代码位置
- 调用栈

---

## 2. Java 与 Go 的对比说明

| 特性 | Java | Go |
|------|------|----|
| **数据竞争检测** | ThreadSanitizer (外部工具) | 内置 `-race` 标志 |
| **易用性** | 需要额外配置 | 非常简单，加个标志即可 |
| **性能开销** | 高 (约 10x-100x) | 中 (约 2x-10x) |
| **集成度** | 独立工具 | 集成到 go 命令中 |

### 代码对照

**Java (ThreadSanitizer)**:
```bash
# 需要特殊编译和运行
javac Main.java
java -javaagent:tsan.jar Main
```

**Go (内置 Race Detector)**:
```bash
# 非常简单，直接加 -race 标志
go run -race main.go
go test -race ./...
```

**心智迁移要点**:
- Java: 数据竞争检测需要外部工具，配置复杂
- Go: 数据竞争检测内置在工具链中，使用非常简单
- Java: ThreadSanitizer 性能开销很大
- Go: Race Detector 性能开销适中，适合在测试中使用
- 两者的输出信息类似，都包含竞争位置、调用栈等
- Go 的 Race Detector 是开发工作流的一部分，应该在测试中定期运行

---

## 3. 可运行代码示例

### 示例 1: 使用 Race Detector 检测数据竞争

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i &lt; 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // 数据竞争！
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
```

运行检测：
```bash
go run -race main.go
```

输出示例：
```
WARNING: DATA RACE
Write at 0x140000a2018 by goroutine 8:
  main.main.func1()
      .../main.go:15 +0x38

Previous write at 0x140000a2018 by goroutine 7:
  main.main.func1()
      .../main.go:15 +0x38
```

### 示例 2: 在测试中使用 Race Detector

```go
package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i &lt; 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // 数据竞争！
		}()
	}

	wg.Wait()
	if counter != 1000 {
		t.Errorf("Expected 1000, got %d", counter)
	}
}
```

运行测试：
```bash
go test -race -v
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 不在测试中运行 race detector**
```bash
# 错误做法：忘记运行 race detector
go test ./...

# 正确做法：定期运行 race detector
go test -race ./...
```

⚠️ **坑点 2: 认为 race detector 能找到所有问题**
```go
// Race Detector 只能找到实际发生的数据竞争
// 如果竞争没有在测试中触发，就不会被检测到
// 解决：充分的测试覆盖，尤其是并发测试
```

⚠️ **坑点 3: 在生产环境使用 race detector**
```bash
# 错误做法：生产环境使用 -race 编译
go build -race -o server

# 正确做法：只在开发和测试环境使用
# 生产环境使用普通编译
go build -o server
```

---

## 5. 练习题

### 练习 2.1: 使用 Race Detector
**目标**: 写一个有数据竞争的程序，然后使用 Race Detector 来检测它，并解读 Race Detector 的输出。

**验收标准**:
- 程序有明确的数据竞争
- 能使用 `go run -race` 或 `go test -race` 检测到竞争
- 能解读 Race Detector 的输出，解释发生了什么

### 练习 2.2: 理解 Race Detector（概念题）
**问题**:
1. Race Detector 是如何工作的？它能保证找到所有的数据竞争吗？
2. 为什么不应该在生产环境使用包含 Race Detector 的程序？

**验收标准**: 能够清晰解释这两个问题。

---

查看 [code/race_detector.go](./code/race_detector.go) 中的 `DemoRaceDetector()` 函数来了解更多，或使用 `go test -race ./unit_02_concurrency/lesson_06_race_detector/code` 来运行检测。
