# Section 2: Import 机制

## 1. 知识点核心说明

### `import` 语法
Go 使用 `import` 关键字导入其他 package：

```go
import "fmt"
import "example.com/utils"
```

批量导入：
```go
import (
    "fmt"
    "example.com/utils"
)
```

### Import 路径
Import 路径是基于 module 路径的，而不是文件系统相对路径。

### 导入方式

| 方式 | 语法 | 说明 |
|------|------|------|
| **默认导入** | `import "fmt"` | 导入后使用 `fmt.Println()` |
| **命名导入** | `import f "fmt"` | 导入后使用 `f.Println()` |
| **点导入** | `import . "fmt"` | 导入后直接使用 `Println()`（不推荐） |
| **下划线导入** | `import _ "example.com/driver"` | 只执行 init 函数，不使用包 |

---

## 2. Java 与 Go 的对比说明

| 特性 | Java Import | Go Import |
|------|-------------|-----------|
| **导入粒度** | 可以导入类或整个包 | 只能导入整个包 |
| **静态导入** | `import static` | 点导入 `import .` |
| **通配符导入** | `import java.util.*` | 不支持，必须显式导入 |
| **初始化** | 类加载时初始化 | 导入时执行 init 函数 |
| **循环导入** | 允许 | 不允许 |

### 代码对照

**Java**:
```java
import java.util.List;
import java.util.ArrayList;
import static java.lang.Math.PI;

public class Main {
    public static void main(String[] args) {
        List&lt;String&gt; list = new ArrayList&lt;&gt;();
        double radius = 2 * PI;
    }
}
```

**Go**:
```go
import (
    "fmt"
    m "math"  // 命名导入
)

func main() {
    fmt.Println("Pi:", m.Pi)
}
```

---

## 3. 可运行代码示例

### 示例 1: 各种导入方式

```go
package main

import (
    "fmt"
    m "math"           // 命名导入
    . "strings"        // 点导入（不推荐在生产环境使用）
    _ "net/http/pprof" // 下划线导入，只触发 init
)

func main() {
    fmt.Println("Pi:", m.Pi)
    fmt.Println("Upper:", ToUpper("hello")) // 点导入后可直接使用
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 滥用点导入**
```go
// 不推荐 - 点导入会导致命名空间污染
import . "fmt"
import . "strings"

func main() {
    Println("Hello") // Println 是 fmt 还是 strings 的？
}

// 推荐做法
import "fmt"

func main() {
    fmt.Println("Hello")
}
```

⚠️ **坑点 2: 循环导入**
```go
// 错误 - Go 不允许循环导入
// a.go 导入 b，b.go 导入 a

// 解决方法：提取公共接口或类型到第三个包
```

⚠️ **坑点 3: 未使用的导入**
```go
// 错误 - Go 不允许未使用的导入
import (
    "fmt"
    "unused" // 编译错误！
)

// 解决方法：删除未使用的导入，或使用下划线导入
import _ "unused"
```

---

## 5. 练习题

### 练习 2.1: 使用不同的导入方式
**目标**: 编写一个程序，演示默认导入、命名导入、下划线导入三种方式的使用。

**验收标准**:
- 使用命名导入给 `math` 包起别名
- 使用下划线导入一个包（如 `time`）
- 程序能正常编译和运行

### 练习 2.2: 理解导入机制（概念题）
**问题**:
1. 为什么 Go 不支持通配符导入（如 `import . "*"`）？
2. 下划线导入有什么用途？请举一个实际场景。

**验收标准**: 能够清晰解释这两个问题。
