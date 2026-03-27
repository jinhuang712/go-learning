# Lesson 2: 集合类型 (Collections & Internals)

本课程主要介绍 Go 语言中的核心集合类型：数组、切片 (Slice) 和 映射 (Map)。
不要被它们简单的语法欺骗，作为资深开发者，你必须深入理解它们在内存中的底层结构，才能避开性能和并发陷阱。

## 课程目录

1. [数组与切片核心剖析](./01_slice_internals.md)
   - 数组 (定长、值传递)
   - 切片底层结构 (Pointer, len, cap)
   - 截取的副作用与 `append` 扩容机制
2. [映射 (Map) 核心剖析与并发陷阱](./02_map_internals.md)
   - `nil` map 恐慌
   - `comma ok` 检查模式
   - 遍历的绝对无序性
   - 致命坑点：原生 Map 并发读写引发 `fatal error`

## 运行方式
在 `main.go` 中调用 `lesson_02_collections.Run()` 即可运行本课程代码。
也可以执行 `go test ./unit_01_fundamentals/lesson_02_collections` 验证练习题。
