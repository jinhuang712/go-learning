# 学习计划 (study-plan.md)

**重要**: 本文件引用 [PROJECT_RULES.md](./PROJECT_RULES.md) 作为主要开发规范，请先阅读该文件。

---

## 学员背景与目标

（详见 [PROJECT_RULES.md](./PROJECT_RULES.md)）

---

## 工作方式（必须严格遵守 Unit -&gt; Lesson -&gt; Section 三级架构）

（详见 [PROJECT_RULES.md](./PROJECT_RULES.md)）

---

## 目录结构规范（强制）

**新规范**: 所有 `.go` 文件放入 `code/` 子目录，`.md` 文档直接放在 Lesson 根目录。

**新结构示例**:
```
unit_01_fundamentals/
  lesson_01_basics/
    lesson_01_basics.md    # Lesson 总览
    01_variables_types.md   # Section 文档
    ...
    code/                   # ✨ 所有 .go 文件在这里
      basics.go
      run.go
      review_test.go
```

（详见 [PROJECT_RULES.md](./PROJECT_RULES.md) 第 3 节）

---

## Section 内容规范（强制）

（详见 [PROJECT_RULES.md](./PROJECT_RULES.md) 第 4 节）

每个 Section 文档必须包含：
1. 知识点核心说明（底层原理）
2. **Java 与 Go 的对比说明（强制）**
3. 可运行代码示例（错误写法 + 正确写法）
4. 生产环境坑点提示
5. 练习题（明确验收标准）

---

## Lesson 验收标准（强制）

（详见 [PROJECT_RULES.md](./PROJECT_RULES.md) 第 5 节）

---

## 课程规划参考（按 Unit 体系进阶）

- Unit 0（新增）：Go Package 系统（Package、Import、Module、可见性）
- Unit 1：Go 基础与心智模型（类型、函数、结构体、接口隐式实现、值与指针、error处理、Java转Go专项坑点）
- Unit 2：并发与并行核心（Goroutine、Channel 模式与底层、Select、Context、Sync/Atomic、Race 检测）
- Unit 3：底层原理与性能调优（G-M-P 调度器、GC 垃圾回收机制、内存逃逸分析、pprof 实战）
- Unit 4：工程化与测试（TDD、Mock、Benchmark、Fuzzing）
- Unit 5：大厂微服务实战（RPC 设计、中间件、限流降级、高可用架构）

---

## 架构思路对比文档

项目包含 `docs/` 目录，存放概念性的 Java-Go 对比文档：
- [java_go_design_patterns.md](./docs/java_go_design_patterns.md) - 设计模式对照
- [java_go_microservices.md](./docs/java_go_microservices.md) - 微服务组件对照
- [java_go_architecture.md](./docs/java_go_architecture.md) - 架构思路对比
- [java_go_concurrency.md](./docs/java_go_concurrency.md) - 并发库对照

---

## 输出语言

默认中文；专业术语可中英并列。每次修改代码后，必须主动执行 `go test ./...` 和 `go run .` 并告知结果。
