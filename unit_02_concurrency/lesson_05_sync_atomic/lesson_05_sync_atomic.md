# Lesson 5: Sync 包与 Atomic 原子操作

本课程主要涵盖 Go sync 包的基本同步原语以及 atomic 包的原子操作，帮助 Java 开发者理解 Go 的并发同步机制。

## 课程目录

1. [Mutex &amp; RWMutex](./01_sync_mutex.md)
   - 互斥锁基础
   - 读写锁适用场景
   - 与 Java ReentrantLock 的对比

2. [WaitGroup, Once, Cond, Map](./02_sync_other.md)
   - WaitGroup: 等待多个 Goroutine 完成
   - Once: 只执行一次
   - Cond: 条件变量
   - sync.Map: 并发安全 Map
   - 与 Java 并发工具的对比

3. [Atomic 原子操作](./03_atomic.md)
   - 基本原子操作
   - 原子类型
   - 什么时候用 atomic，什么时候用 Mutex
   - 与 Java AtomicInteger 的对比

## 运行方式
在 `main.go` 中调用 `lesson_05_sync_atomic.Run()` 即可运行本课程代码。
