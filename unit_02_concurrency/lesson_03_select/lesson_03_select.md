# Lesson 3: Select 多路复用与超时控制

本课程主要涵盖 Go Select 语句的使用、常用模式以及生产环境中的常见坑点。

## 课程目录

1. [Select 基础](./01_select_basics.md)
   - Select 语法与随机选择
   - 与 Switch 的区别
   - 与 Java NIO Selector 的对比

2. [Select 模式](./02_select_patterns.md)
   - 超时控制 (time.After)
   - 非阻塞读写
   - 多路等待

3. [Select 坑点](./03_select_pitfalls.md)
   - nil Channel 的行为
   - 空 `select {}` 的问题
   - 与 Java 超时机制的对比

## 运行方式
在 `main.go` 中调用 `lesson_03_select.Run()` 即可运行本课程代码。
