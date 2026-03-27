# 3. 函数与可见性

## 3.1 函数与多返回值
Go 的函数签名设计与 Java 有所不同，类型写在变量名后面。
- **基本声明**: `func add(a int, b int) int`
- **简写**: `func add(a, b int) int`
- **多返回值**: Go 原生支持返回多个值，这在 Go 中极其常见，通常用来返回 `(结果, 错误)`。
  ```go
  func divMod(a, b int) (int, int) {
      return a / b, a % b
  }
  
  // 调用时接收
  q, r := divMod(10, 3)
  ```

## 3.2 可见性 (Visibility)
Go **没有** `public`、`private` 或 `protected` 关键字。
它通过标识符的**首字母大小写**来决定可见性（仅在包级别生效）。

- **大写字母开头 (Public/Exported)**: 可以被其他 package 导入和使用。
  - 例如: `fmt.Println()`, `calc.PublicFunc()`, 结构体的导出字段 `User.Name`。
- **小写字母开头 (Private/Unexported)**: 只能在当前 package 内部使用。
  - 例如: `calc.privateFunc()`, 结构体的私有字段 `User.age`。