# Lesson 2: Channel 模式与底层原理

本课程主要涵盖 Go Channel 的核心概念、常用模式以及底层实现原理，帮助 Java 开发者建立 CSP 并发模型的完整心智。

## 课程目录

1. [Channel 基础](./01_channel_basics.md)
   - 无缓冲 vs 有缓冲 Channel
   - 发送、接收、关闭操作
   - 与 Java BlockingQueue 的对比

2. [Channel 常用模式](./02_channel_patterns.md)
   - 生产者-消费者模式
   - 扇入扇出 (Fan-in/Fan-out)
   - 信号量模式

3. [Channel 底层原理](./03_channel_underhood.md)
   - `hchan` 结构
   - 环形缓冲区
   - 等待队列

## 运行方式
在 `main.go` 中调用 `lesson_02_channel.Run()` 即可运行本课程代码。
