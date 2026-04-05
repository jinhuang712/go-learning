# Lesson 6: Java 转 Go 常见坑点专项

本课程专门汇总资深 Java 开发者转向 Go 时最容易踩的典型坑点，避免在生产环境中踩雷。

## 课程目录

1. [值/指针误用坑点](./01_value_pointer_pitfalls.md)
   - 方法接收者混用值和指针导致的状态修改失效
   - 大结构体传值导致的性能开销
   - 值类型的nil判断误区

2. [接口nil判断坑点](./02_interface_nil_pitfalls.md)
   - 接口nil判断的底层原理
   - 错误返回时返回nil指针导致的错误吞掉问题
   - 防御性编程最佳实践

3. [错误处理坑点](./03_error_handling_pitfalls.md)
   - 模仿Java异常把error当异常抛
   - 错误缺少上下文导致排查困难
   - wrap和unwrap错误的正确用法

4. [并发模型坑点](./04_concurrency_pitfalls.md)
   - Goroutine泄露
   - Channel阻塞/关闭导致的panic
   - 数据竞争检测的正确使用

## 运行方式
在 `main.go` 中调用 `lesson_06_java_go_pitfalls.Run()` 即可运行本课程代码示例。
也可以执行 `go test ./unit_01_fundamentals/lesson_06_java_go_pitfalls` 验证练习题。