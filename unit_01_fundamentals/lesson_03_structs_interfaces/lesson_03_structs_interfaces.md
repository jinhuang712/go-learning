# Lesson 3: 结构体与接口 (Structs & Interfaces)

本课程主要介绍 Go 的面向对象编程 (OOP) 核心：结构体、方法、接口以及组合。这与 Java 基于类的继承体系有着本质的区别。

## 课程目录

1. [结构体 (Structs) 深度剖析](./01_struct_internals.md)
   - 内存布局与值类型本质
   - 方法接收者：值传递 vs 指针传递
   - 组合优于继承 (匿名嵌入)
   - 结构体标签 (Struct Tags) 的微服务应用
2. [接口 (Interfaces) 底层与设计哲学](./02_interface_internals.md)
   - 隐式实现 (Duck Typing) 与解耦优势
   - 接口的底层胖指针：`iface` 与 `eface` (`any`)
   - 致命坑点：接口的 `nil` 判断逻辑
   - 类型断言 (Type Assertion)

## 运行方式
在 `main.go` 中调用 `lesson_03_structs_interfaces.Run()` 即可运行本课程代码。
也可以执行 `go test ./unit_01_fundamentals/lesson_03_structs_interfaces` 验证练习题。
