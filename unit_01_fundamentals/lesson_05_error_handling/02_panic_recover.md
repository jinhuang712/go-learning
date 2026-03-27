# 2. 崩溃与恢复 (Panic & Recover)

既然 Go 不提倡用异常 (`Exception`)，那如果程序遇到真的无法继续执行的严重错误（比如空指针、数组越界、配置加载失败）该怎么办？

答案是：**Panic (恐慌)**。

## 2.1 Panic (引发恐慌)
`panic` 类似于 Java 中抛出 `RuntimeException` 或 `Error`。一旦调用 `panic`，正常的控制流会被中断，当前 goroutine 中的所有 `defer` 语句会按倒序执行，最后程序崩溃并打印堆栈跟踪。

**什么时候用 Panic？**
在微服务中，**永远不要用 panic 处理业务逻辑错误**（如密码错误、找不到数据）。
只有在**不可恢复的程序 Bug**（如必须要初始化的配置缺失、明显的开发者错误）时，才使用 `panic`。

## 2.2 Recover (恢复)
如果在 Web 服务中，某个请求触发了 `panic`，我们显然不希望整个微服务进程挂掉。我们需要一种机制捕获这个 panic，这就是 `recover`。

- `recover()` 必须**只在 `defer` 函数内部调用**才有效。
- 如果没有发生 `panic`，`recover()` 返回 `nil`。
- 捕获到 `panic` 后，程序不会崩溃，而是从触发 panic 的函数正常返回，继续执行上层代码。

**Java 与 Go 崩溃恢复对比**:

| Java | Go | 适用场景 |
| :--- | :--- | :--- |
| `try { ... }` | 正常写代码 | 业务代码域 |
| `catch (Exception e) { ... }` | `defer func() { if r := recover(); r != nil { ... } }()` | 捕获意料之外的严重错误，防止进程退出 |
| `finally { ... }` | `defer` | 资源清理 |

## 2.3 微服务中的最佳实践
在 Gin 或 Hertz 这种 Web 框架中，通常会有一个全局的 `Recovery Middleware`（恢复中间件）。
它的作用就是在所有请求的最外层套一个 `defer recover()`，这样无论哪个业务开发人员写出了空指针 bug，都只会导致当前这个请求返回 `HTTP 500`，而不会导致整个微服务宕机。