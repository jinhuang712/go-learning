# 4. Java 与 Go 核心关键字对照表

很多 Java 里的关键字在 Go 里并不存在，Go 更倾向于用简洁的语法实现类似的功能：

| Java 关键字/概念 | Go 中的对应实现 | 说明 |
| :--- | :--- | :--- |
| `class` | `struct` | Go 没有类的概念，通过结构体（struct）来组合数据。 |
| `public` / `private` | **首字母大小写** | 大写字母开头即公开（Public），小写即私有（包内私有）。 |
| `extends` (继承) | **结构体嵌套 (组合)** | Go 崇尚组合优于继承，没有 `extends`。 |
| `implements` | **隐式实现** | Go 的接口不需要显式声明，实现了接口的方法即自动实现该接口（Duck Typing）。 |
| `this` | **方法接收者 (Receiver)** | Go 必须在方法签名里显式声明接收者变量（如 `func (u *User) ...`），不使用隐式的 `this`。 |
| `static` | **包级变量/函数** | Go 没有静态方法或属性，直接定义在包（Package）级别即可。 |
| `final` | `const` (部分对应) | Go 的 `const` 只支持基本类型（数字、字符串、布尔），不支持不可变对象/引用。 |
| `try` / `catch` / `throw` | 多返回值 `error` + `panic`/`recover` | Go 主要通过函数返回 `error` 值进行显式错误处理；`panic` 仅用于极其严重的不可恢复故障。 |
| `finally` | `defer` | Go 用 `defer` 保证在函数退出前执行清理逻辑。 |
| `void` | **无返回值** | Go 函数若无返回值则直接留空。 |
| `while` / `do...while`| `for` | Go 只有 `for` 关键字，`for {}` 等价于 `while(true)`。 |
| `enum` | `const` + `iota` | Go 没有原生枚举类型，通常用常量组配合自增生成器 `iota` 模拟。 |