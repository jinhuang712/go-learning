# Lesson 1: G-M-P 调度器模型与阻塞原理

本课程主要涵盖 Go G-M-P 调度器模型的详细讲解、Goroutine 调度时机以及阻塞原理与处理。

## 课程目录

1. [G-M-P 模型详解](./01_gmp_model.md)
   - M、P、G 分别是什么
   - 三者之间的关系
   - 与 Java 线程调度的对比

2. [Goroutine 调度时机](./02_scheduling_triggers.md)
   - 函数调用/栈扩容
   - 系统调用
   - runtime.Gosched()
   - 同步原语

3. [阻塞原理与处理](./03_blocking_principles.md)
   - Goroutine 阻塞时会发生什么
   - M 的阻塞与非阻塞
   - 与 Java 线程阻塞的对比

## 运行方式
在 `main.go` 中调用 `lesson_01_gmp_scheduler.Run()` 即可运行本课程代码。
