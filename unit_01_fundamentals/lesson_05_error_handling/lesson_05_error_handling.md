# Lesson 5: 错误处理 (Error Handling & Panic/Recover)

本课程涵盖 Go 语言最核心的错误处理心智模型。由于 Go 抛弃了 Java 中的异常 (Exception) 机制，我们需要建立全新的防御式编程思维。

## 课程目录

1. [错误接口与显式错误处理](./01_error_interface.md)
   - `error` 接口的本质
   - 多返回值与 `if err != nil`
   - 自定义错误类型
2. [崩溃与恢复 (Panic & Recover)](./02_panic_recover.md)
   - 什么时候用 `panic`
   - 为什么不要用 `panic` 代替异常
   - 如何使用 `defer` 和 `recover` 兜底严重故障

## 运行方式
在 `main.go` 中调用 `lesson_05_error_handling.Run()` 即可运行本课程代码。也可以执行 `go test ./unit_01_fundamentals/lesson_05_error_handling` 验证你的练习题。