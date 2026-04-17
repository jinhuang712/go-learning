# Section 1: Go Package 基础

## 1. 知识点核心说明

### 什么是 Go Package
Package 是 Go 代码的基本组织单元，用于将相关的 Go 源代码组织在一起。

**核心规则**:
- 一个目录下的所有 `.go` 文件必须声明属于同一个 package
- package 名通常是目录名的最后一段（但不一定必须相同）
- 同一个 package 下的文件可以互相访问未导出（小写开头）的标识符

### `package` 声明
每个 `.go` 文件的第一行必须是 `package` 声明：

```go
package mypackage
```

### 目录结构与 Package 的关系
Go 使用文件系统目录来组织 package，而不是像 Java 那样在文件中声明包路径。

```
example.com/
├── go.mod
├── main.go
└── utils/
    ├── string.go  // package utils
    └── math.go    // package utils
```

---

## 2. Java 与 Go 的对比说明

| 特性 | Java Package | Go Package |
|------|-------------|------------|
| **声明方式** | `package com.example.utils;` | `package utils` |
| **目录与包关系** | 目录结构必须与包名一致 | 目录结构决定 import 路径，package 名独立 |
| **包名规范** | 反向域名（com.example） | 小写、简洁、不包含下划线 |
| **文件组织** | 一个文件可以声明任意包 | 一个目录下的所有文件必须同属一个包 |

### 代码对照

**Java**:
```java
// 文件位置: com/example/utils/StringUtils.java
package com.example.utils;

public class StringUtils {
    public static String reverse(String s) {
        // ...
    }
}
```

**Go**:
```go
// 文件位置: utils/string.go
package utils

func Reverse(s string) string {
    // ...
}
```

---

## 3. 可运行代码示例

### 示例 1: 基本的 Package 组织

```go
// 目录结构:
// demo/
// ├── go.mod
// ├── main.go
// └── greeting/
//     ├── hello.go
//     └── hi.go

// greeting/hello.go
package greeting

import "fmt"

func Hello() {
	fmt.Println("Hello!")
}

// greeting/hi.go
package greeting  // 同一个目录，必须是同一个 package

import "fmt"

func Hi() {
	fmt.Println("Hi!")
}

// main.go
package main

import (
	"demo/greeting"
)

func main() {
	greeting.Hello()
	greeting.Hi()
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 同一个目录下不同的 package 声明**
```go
// 错误写法 - 同一个目录下不能有不同的 package
// file1.go
package utils

// file2.go (同一个目录)
package helpers // 编译错误！

// 正确写法 - 必须同一个 package
// file2.go
package utils
```

⚠️ **坑点 2: 循环 import**
```go
// 错误写法 - 循环导入
// a.go 导入 b，b.go 导入 a
// Go 不允许循环导入！

// 正确做法
// 重构代码，提取公共部分到第三个 package
```

⚠️ **坑点 3: package 名和目录名不一致**
```go
// 虽然允许，但不推荐
// 目录: stringutils/
// 文件: stringutils/str.go
package str // package 名与目录名不一致

// 推荐做法
package stringutils // package 名与目录名一致
```

---

## 5. 练习题

### 练习 1.1: 创建一个简单的 Package
**目标**: 创建一个 `mathutil` package，包含 `Add` 和 `Multiply` 两个函数，在 main 中调用它们。

**验收标准**:
- 正确的目录结构
- package 声明正确
- 能正常编译和运行

### 练习 1.2: 理解目录与 Package 的关系（概念题）
**问题**:
1. Go 中一个目录下可以有多个不同的 package 吗？为什么？
2. Java 和 Go 在 package 组织方式上有什么本质区别？

**验收标准**: 能够清晰解释这两个问题。
