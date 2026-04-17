# Lesson 4: Context 传递与级联取消

本课程主要涵盖 Go Context 的使用、派生方式以及在微服务开发中的传递最佳实践。

## 课程目录

1. [Context 基础](./01_context_basics.md)
   - Context 接口
   - `context.Background()` vs `context.TODO()`
   - 与 Java 的对比（为什么 Go 需要 Context）

2. [派生 Context](./02_context_derived.md)
   - `WithCancel`
   - `WithTimeout` / `WithDeadline`
   - `WithValue`
   - 与 Java ThreadLocal 的对比

3. [Context 传递最佳实践](./03_context_propagation.md)
   - 如何在函数间传递
   - 不要存在结构体里
   - `WithValue` 的 key 类型最佳实践

## 运行方式
在 `main.go` 中调用 `lesson_04_context.Run()` 即可运行本课程代码。
