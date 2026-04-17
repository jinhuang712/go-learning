# Lesson 6: Race Detector 数据竞争检测实战

本课程主要涵盖数据竞争的概念、Go Race Detector 的使用以及如何修复数据竞争问题。

## 课程目录

1. [什么是数据竞争](./01_data_race.md)
   - 数据竞争的定义
   - 常见场景
   - 与 Java 的对比

2. [Race Detector 使用](./02_race_detector.md)
   - `-race` 标志
   - 如何解读输出
   - 与 Java 工具的对比

3. [修复数据竞争](./03_fixing_races.md)
   - 几种修复方案对比
   - 性能 trade-off
   - 与 Java 修复方案的对比

## 运行方式
在 `main.go` 中调用 `lesson_06_race_detector.Run()` 即可运行本课程代码。
