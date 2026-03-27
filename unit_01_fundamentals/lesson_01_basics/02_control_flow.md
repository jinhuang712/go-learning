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

## 2.2 `if/else` 的特殊写法 (带初始化语句)
Go 的 `if` 语句允许在进行条件判断之前，**先执行一条简短的初始化语句**。这两个部分用分号 `;` 隔开。

**基本语法**: `if 初始化语句; 条件表达式 { ... }`

**Java 里的痛点**:
在 Java 中，如果你想调用一个方法，判断它的结果，并且只在 `if` 块内使用这个结果，你通常得这么写：
```java
// res 的作用域污染了整个外部方法
Result res = doSomething();
if (res != null) {
    System.out.println(res);
}
// 这里依然可以访问到 res
```

**Go 的优雅解法**:
通过带初始化的 `if`，我们可以把变量的作用域严格限制在 `if`（以及随后的 `else`）块内部，避免变量污染外部作用域。这种写法在 Go 的错误处理和 Map 查找中极其常见。

```go
// 1. err 变量在这里被声明和赋值
// 2. 紧接着判断 err != nil
if err := doSomething(); err != nil {
    // 3. 只有在 if 块内，err 变量才有效
    fmt.Println("出错了:", err)
    return err
}
// 4. 离开 if 块后，err 变量就被销毁了，外部无法访问
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