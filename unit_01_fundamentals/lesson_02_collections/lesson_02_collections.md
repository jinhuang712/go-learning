# Lesson 2: 集合类型 (Collections & Internals)

本课程主要介绍 Go 语言中的核心集合类型：数组、切片 (Slice) 和 映射 (Map)。
不要被它们简单的语法欺骗，作为资深开发者，你必须深入理解它们在内存中的底层结构，才能避开性能和并发陷阱。

## 课程目录

1. [数组 (Array) 核心剖析](./01_array_internals.md)
   - 长度是类型的一部分
   - 纯粹的值传递
2. [切片 (Slice) 核心剖析与坑点](./02_slice_internals.md)
   - 切片底层结构 (Pointer, len, cap)
   - make 创建语法
   - 截取的副作用与安全 Clone 技巧
   - append 扩容机制与必写 `=` 的原因
3. [映射 (Map) 核心剖析与并发陷阱](./03_map_internals.md)
   - `nil` map 恐慌
   - `comma ok` 检查模式
   - 遍历的绝对无序性
   - 致命坑点：原生 Map 并发读写引发 `fatal error`

## 运行方式
在 `main.go` 中调用 `lesson_02_collections.Run()` 即可运行本课程代码。
也可以执行 `go test ./unit_01_fundamentals/lesson_02_collections` 验证练习题。
