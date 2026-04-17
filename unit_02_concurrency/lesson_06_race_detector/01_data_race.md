# Section 1: 什么是数据竞争

## 1. 知识点核心说明

### 数据竞争的定义
当两个或多个 Goroutine 同时访问同一个变量，且至少有一个是写操作时，就会发生数据竞争。

```go
var counter int

// Goroutine 1
go func() {
    counter++ // 写操作
}()

// Goroutine 2
go func() {
    counter++ // 写操作
}()
// 数据竞争！
```

### 为什么数据竞争危险
- 结果不确定
- 可能导致 panic
- 难以调试和复现
- 可能污染内存

---

## 2. Java 与 Go 的对比说明

| 特性 | Java | Go |
|------|------|----|
| **数据竞争问题** | 同样存在 | 同样存在 |
| **检测工具** | ThreadSanitizer, FindBugs | 内置 `-race` 检测器 |
| **内存模型** | Java Memory Model (JMM) | Go Memory Model |
| **Happens-Before** | 有明确定义 | 有明确定义 |

### 代码对照

**Java (数据竞争示例)**:
```java
public class Counter {
    private static int count = 0;
    
    public static void main(String[] args) throws InterruptedException {
        Thread t1 = new Thread(() -&gt; {
            for (int i = 0; i &lt; 10000; i++) {
                count++;
            }
        });
        
        Thread t2 = new Thread(() -&gt; {
            for (int i = 0; i &lt; 10000; i++) {
                count++;
            }
        });
        
        t1.start();
        t2.start();
        t1.join();
        t2.join();
        
        // 结果不一定是 20000！
        System.out.println("Count: " + count);
    }
}
```

**Go (数据竞争示例)**:
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i &lt; 10000; i++ {
			counter++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i &lt; 10000; i++ {
			counter++
		}
	}()

	wg.Wait()
	fmt.Println("Count:", counter) // 结果不一定是 20000！
}
```

**运行检测**:
```bash
# Go 内置检测
go run -race main.go

# Java 需要外部工具
java -javaagent:tsan.jar Main
```

**心智迁移要点**:
- Java 和 Go 都有数据竞争问题
- Java 需要使用外部工具（ThreadSanitizer、FindBugs）
- Go 内置了 race detector，使用非常方便
- 两者的内存模型都有 Happens-Before 关系的明确定义
- 数据竞争问题在两种语言中都非常危险，必须避免

---

## 3. 可运行代码示例

### 示例 1: 简单的数据竞争

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
	fmt.Println("Counter:", counter) // 不一定是 1000！
}
```

运行检测：
```bash
go run -race main.go
```

Race Detector 会输出类似这样的信息：
```
WARNING: DATA RACE
Write at 0x... by goroutine 7:
  main.main.func1()
      .../main.go:12 +0x3a

Previous write at 0x... by goroutine 6:
  main.main.func1()
      .../main.go:12 +0x3a
```

### 示例 2: 读写数据竞争

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var data string
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		data = "updated" // 写操作
	}()

	// 读操作，可能读到旧值或新值
	fmt.Println("Reading data:", data)

	wg.Wait()
	fmt.Println("Final data:", data)
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 忽视数据竞争**
```go
// 错误做法：假设数据竞争不会发生
var config string

go func() {
    config = "new config" // 写
}()

fmt.Println(config) // 读，可能读到旧值！
```

⚠️ **坑点 2: map 的并发读写**
```go
// 错误做法：map 不是并发安全的
var m = make(map[string]string)

go func() {
    m["key"] = "value" // 写
}()

fmt.Println(m["key"]) // 读，可能 panic！
```

⚠️ **坑点 3: slice 的并发修改**
```go
// 错误做法：slice 的 append 可能重新分配内存
var s []int

go func() {
    s = append(s, 1) // 可能重新分配内存
}()

fmt.Println(s[0]) // 可能 panic！
```

---

## 5. 练习题

### 练习 1.1: 识别数据竞争
**目标**: 找出下面代码中的数据竞争，并解释为什么会有数据竞争。

```go
package main

import (
	"sync"
	"time"
)

type User struct {
	Name string
	Age  int
}

var user *User

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		user = &amp;User{Name: "Alice", Age: 30}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if user != nil {
			println("User:", user.Name)
		}
	}()

	wg.Wait()
}
```

**验收标准**:
- 找出所有数据竞争点
- 解释每一个数据竞争为什么会发生
- 说明可能导致的问题

### 练习 1.2: 理解数据竞争（概念题）
**问题**:
1. 什么是数据竞争？它为什么危险？
2. Go 的内存模型中，"Happens-Before" 关系是什么意思？

**验收标准**: 能够清晰解释这两个问题。

---

查看 [code/race_detector.go](./code/race_detector.go) 中的 `DemoDataRace()` 函数来运行示例代码（注意：该示例有数据竞争，使用 `go run -race` 可以检测到）。
