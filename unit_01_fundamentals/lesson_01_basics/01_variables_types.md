# 1. 变量、常量与基本数据类型

## 1.1 变量声明
Go 提供了多种声明变量的方式，旨在简化代码：
- **显式声明**: `var name string = "Go"` (类似 Java `String name = "Go";`)
- **类型推断**: `var name = "Go"` (类似 Java `var name = "Go";`)
- **短变量声明**: `name := "Go"` (**Go 特有**，最常用，仅限函数内使用)

## 1.2 常量与枚举 (const & iota)
- **常量声明**: 使用 `const` 关键字。
  ```go
  const pi = 3.14159
  ```
- **枚举模拟 (`iota`)**: Go 没有 `enum` 关键字，通常使用 `const` 配合 `iota` 实现。
  ```go
  const (
      Sunday = iota // 0
      Monday        // 1
      Tuesday       // 2
  )
  ```

## 1.3 基本数据类型
Go 的基本数据类型与 Java 类似，但有一些自己的特点：
- **整型**: `int`, `int8`, `int16`, `int32`, `int64` (以及对应的无符号类型 `uint` 等)。
  - `int` 的大小取决于操作系统（32位或64位），类似于 Java 的 `int` 或 `long`。
- **浮点型**: `float32`, `float64` (推荐使用 `float64`，对应 Java 的 `double`)。
- **布尔型**: `bool` (`true` / `false`)。
- **字符串**: `string` (Go 的字符串是不可变的字节切片，默认 UTF-8 编码)。
- **字符**:
  - `byte`: 等同于 `uint8`，通常用于处理 ASCII 字符或二进制数据。
  - `rune`: 等同于 `int32`，用于处理 Unicode 字符 (对应 Java 的 `char`)。