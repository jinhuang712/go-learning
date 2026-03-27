# Lesson 1: Go 基础语法 (Basics)

本课程主要涵盖 Go 语言的核心基础概念，并与 Java 进行对比。为了方便阅读，我们将本课内容拆分成了多个子文档：

## 课程目录

1. [变量、常量与基本数据类型](./01_variables_types.md)
   - `var`, `:=`, `const`, `iota`
   - `int`, `string`, `bool`, `rune`, `byte`
2. [控制流与 Defer](./02_control_flow.md)
   - `for` (替代 `while`)
   - `if` 的特殊写法
   - `switch` (默认不穿透)
   - `defer` (替代 `finally`)
3. [函数与可见性](./03_functions_visibility.md)
   - 多返回值
   - 首字母大小写控制访问权限
4. [Java 与 Go 核心关键字对照表](./04_java_go_mapping.md)
   - Java 开发者迁移到 Go 时最重要的心智映射

## 运行方式
在 `main.go` 中调用 `basics.Run()` 即可运行本课程代码。
