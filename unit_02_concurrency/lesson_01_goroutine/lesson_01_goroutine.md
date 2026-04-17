# Lesson 1: Goroutine &amp; CSP 模型基础

本课程主要涵盖 Go 并发编程的基础：Goroutine 的概念、CSP 模型、调度原理以及常见的 Goroutine 泄露问题。

## 课程目录

1. [Goroutine 基础与底层原理](./01_goroutine_intro.md)
   - `go` 关键字
   - M-P-G 模型初体验
   - 与 Java Thread 的对比

2. [Goroutine 调度与生命周期](./02_goroutine_scheduling.md)
   - 启动、运行、阻塞、退出
   - GOMAXPROCS 的作用

3. [Goroutine 泄露初体验](./03_goroutine_leak.md)
   - 泄露场景演示
   - 为什么会泄露

## 运行方式
在 `main.go` 中调用 `lesson_01_goroutine.Run()` 即可运行本课程代码。
