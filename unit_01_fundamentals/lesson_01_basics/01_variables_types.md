# 1. 变量、常量与基本数据类型

## 1.1 变量声明
Go 提供了多种声明变量的方式，旨在简化代码：
- **显式声明**: `var name string = "Go"` (类似 Java `String name = "Go";`)
- **类型推断**: `var name = "Go"` (类似 Java `var name = "Go";`)
- **短变量声明**: `name := "Go"` (**Go 特有**，最常用，仅限函数内使用)

## 1.2 常量与枚举 (const & iota)
- **常量声明**: 使用 `const` 关键字。常量的可见性规则与变量完全一致，**由首字母大小写决定**：
  - `const pi = 3.14159` (小写字母开头，仅当前包内可见，类似 Java 的 `private`)
  - `const Pi = 3.14159` (大写字母开头，其他包也可导入使用，类似 Java 的 `public`)
  - 如果在函数内部声明 `const`，则只是该函数的局部常量。
  
- **枚举模拟 (`iota`)**: Go 没有 `enum` 关键字。`iota` 是 Go 语言的**内置常量生成器**。
  - 当代码中遇到 `const` 关键字时，`iota` 会被重置为 `0`。
  - `const` 块中每新增一行常量声明，`iota` 就会自动 `+1`。
  ```go
  const (
      Sunday = iota // 0 (iota 被重置为 0)
      Monday        // 1 (未写赋值表达式，默认延续上一行的规则，此时 iota=1)
      Tuesday       // 2 (此时 iota=2)
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