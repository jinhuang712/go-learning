# Lesson 0: Go Package 系统详解

本课程为 Java 开发者详细讲解 Go 的 Package、Import 和 Module 系统，建立清晰的心智模型。

## 课程目录

1. [01_package_basics.md](./01_package_basics.md) - Go Package 基础
   - 什么是 Package
   - `package` 声明
   - 目录结构与 Package 的关系

2. [02_import_mechanism.md](./02_import_mechanism.md) - Import 机制
   - `import` 语法
   - 相对导入 vs 绝对导入
   - 命名导入、点导入、下划线导入
   - 与 Java import 的对比

3. [03_go_modules.md](./03_go_modules.md) - Go Module 系统
   - `go.mod` 文件
   - `module` 声明
   - 与 Java Maven/Gradle 的对比

4. [04_visibility_rules.md](./04_visibility_rules.md) - 可见性规则
   - 首字母大小写的意义
   - 与 Java public/private/protected 的对比

## 运行方式
在 `main.go` 中调用 `packages.Run()` 即可运行本课程代码。
