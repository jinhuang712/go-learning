# 2. 控制流与 Defer

Go 的控制流比 Java 更精简，去掉了冗余的括号，并增强了部分功能。

## 2.1 只有 `for` 循环
Go 没有 `while` 或 `do-while`，一切皆 `for`。

- **标准循环**:
  ```go
  for i := 0; i < 3; i++ { ... }
  ```
- **`while` 替代品**:
  ```go
  n := 0
  for n < 5 {
      n++
  }
  ```
- **`while true` (无限循环)**:
  ```go
  for {
      // 业务逻辑
      if needStop {
          break
      }
  }
  ```

## 2.2 `if/else` 的特殊写法
Go 的 `if` 语句可以在条件判断前执行一个简短的初始化语句（常用于错误检查）。
```go
if err := doSomething(); err != nil {
    // 处理错误，err 变量只在这个 if 块内有效
    return err
}
```

## 2.3 `switch` (默认不穿透)
与 Java 最大的不同是，Go 的 `switch` **默认不会穿透 (fallthrough)**，匹配到一个 `case` 执行完后自动退出，**不需要写 `break`**。
```go
switch day {
case "Monday", "Tuesday": // 可以匹配多个值
    fmt.Println("Workday")
case "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Unknown")
}
```

## 2.4 `defer` (延迟执行)
`defer` 是 Go 的一大特色，用于确保函数在退出前一定会执行某些清理操作（类似于 Java 的 `finally`，但写法更就地、更优雅）。
- **执行时机**: 当前函数 `return` 之前。
- **常见场景**: 关闭文件、释放锁、关闭数据库连接。
```go
func processFile() {
    file := openFile()
    defer file.Close() // 无论后面发生什么，函数退出时都会执行这句

    // 业务逻辑...
}
```