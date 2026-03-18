# Lesson 1: Go 基础语法 (Basics)

本课程主要涵盖 Go 语言的核心基础概念，并与 Java 进行对比。

## 核心知识点

### 1. 变量声明
Go 提供了多种声明变量的方式，旨在简化代码：
- **显式声明**: `var name string = "Go"` (类似 Java `String name = "Go";`)
- **类型推断**: `var name = "Go"` (类似 Java `var name = "Go";`)
- **短变量声明**: `name := "Go"` (**Go 特有**，最常用，仅限函数内)

### 2. 函数与多返回值
- **类型后置**: `func add(a int, b int) int`
- **多返回值**: Go 支持原生多返回值，常用于返回结果和错误。
  ```go
  func divMod(a, b int) (int, int) {
      return a / b, a % b
  }
  ```

### 3. 控制流
- **只有 `for`**: Go 只有 `for` 循环，没有 `while`。
  ```go
  for i := 0; i < 3; i++ { ... }
  ```

### 4. 可见性 (Visibility)
Go 通过首字母大小写控制访问权限：
- **大写开头 (Public)**: 可被其他包访问 (e.g., `fmt.Println`, `calc.PublicFunc`)
- **小写开头 (Private)**: 仅当前包内可见

## 运行方式
在 `main.go` 中调用 `basics.Run()` 即可运行本课程代码。
