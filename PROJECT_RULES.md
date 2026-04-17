# 项目开发规范 (PROJECT_RULES.md)

## 1. 学员背景与目标

**学员背景**: 资深 Java 后端工程师，长期从事国际化分布式微服务开发

**学习目标**: 达到"在字节跳动等大厂浸淫多年的资深 Golang 分布式微服务开发者"水平，能够独立设计、阅读、评审、优化极高并发与复杂度的 Go 核心代码。

**核心迁移**: 从 Java 生态迁移到 Go 生态，重点掌握：
- 心智模型：值语义/指针、接口、错误处理、并发模型
- 底层原理：G-M-P 调度器、GC、内存逃逸、性能调优
- 微服务实战：context 传递、超时重试、幂等、限流熔断、可观测性

---

## 2. 强制规范：Java 对比要求

### 2.1 基本原则
**所有知识点必须包含详细的 Java-Go 对照**，因为学员是 Java 资深开发者，需要建立清晰的迁移映射。

### 2.2 Java 对比的内容要求

每个 Section 文档必须包含：

| 对比类型 | 要求 |
|---------|------|
| **概念对照表** | 表格形式，列出 Java 概念与 Go 对应概念的差异 |
| **代码对照** | 同时提供 Java 代码示例和 Go 代码示例 |
| **心智迁移要点** | 重点说明 Java 开发者需要注意的思维转变 |
| **底层原理解析** | 解释"为什么 Go 这样设计" |

### 2.3 强化对比的主题

除了语法层面，还需要包含：
- **设计模式对照**: Singleton、Factory、Strategy 等在 Go 中的实现
- **微服务组件对照**: Spring Cloud vs Go 生态（gRPC、gin、zap、viper 等）
- **并发库对照**: `java.util.concurrent` vs `sync`/`channel`
- **架构思路对照**: 面向对象 vs 组合、异常 vs 值错误、IoC vs 依赖注入

---

## 3. 目录结构规范

### 3.1 整体架构
```
go-learning/
├── PROJECT_RULES.md      # ✅ 本文件：开发规范（单一真相源）
├── README.md              # 项目说明与进度
├── study-plan.md          # 学习计划（引用 PROJECT_RULES.md）
├── go.mod                 # Go Module 定义
├── main.go                # 主程序入口
├── pkg/                   # 公共包
├── docs/                  # ✅ 新增：架构思路对比文档
├── unit_01_fundamentals/  # Unit 1: 基础
├── unit_02_concurrency/   # Unit 2: 并发
└── ...
```

### 3.2 Lesson 目录结构（强制）

**问题**: .md 和 .go 文件混在一起，很乱

**解决方案**: 所有 `.go` 文件放入 `code/` 子目录

**新结构示例**:
```
lesson_01_goroutine/
├── lesson_01_goroutine.md    # Lesson 总览与目录
├── 01_goroutine_intro.md      # Section 1 文档
├── 02_goroutine_scheduling.md # Section 2 文档
├── 03_goroutine_leak.md       # Section 3 文档
└── code/                       # ✅ 新增：所有 .go 文件
    ├── goroutine.go            # 演示代码
    ├── run.go                  # 统一 Run() 入口
    └── review_test.go          # 单元测试
```

### 3.3 目录规范说明

| 位置 | 内容 | 说明 |
|------|------|------|
| `lesson_xx_*/` 根目录 | `.md` 文档 | 所有 Markdown 文档直接放在 Lesson 根目录 |
| `lesson_xx_*/code/` | `.go` 文件 | 所有 Go 代码文件（.go、_test.go）放在 code 子目录 |

**Import 路径变化**:
```go
// 重构前
import "go-learning/unit_02_concurrency/lesson_01_goroutine"

// 重构后
import "go-learning/unit_02_concurrency/lesson_01_goroutine/code"

// 调用方式不变（包名保持不变）
lesson_01_goroutine.Run()
```

---

## 4. Section 文档规范（强制）

每个 Section 文档必须包含以下 5 个部分：

### 4.1 知识点核心说明（底层原理）
- 解释该知识点的核心概念
- 说明 Go 为什么这样设计
- 必要时讲解底层实现原理

### 4.2 Java 与 Go 的对比说明（强制）
- 表格形式的概念对照
- 突出 Java 开发者需要注意的差异
- 心智迁移要点

### 4.3 可运行代码示例
- **错误写法** + **正确写法** 的对比
- Java 代码示例 + Go 代码示例 对照
- 代码必须可直接运行

### 4.4 生产环境坑点提示
- 列出该知识点在实际生产中常见的陷阱
- 提供解决方案或规避建议

### 4.5 练习题（明确验收标准）
- 设计与知识点相关的练习
- 明确验收标准（如何判断练习完成）

---

## 5. Lesson 验收标准（强制）

每个 Lesson 完成后必须满足：

1. ✅ `go run .` 能正常执行 Lesson 演示代码，无 panic
2. ✅ `go test ./<lesson目录>/code` 所有测试用例通过
3. ✅ 所有 Section 文档完整，包含 Java 对比和坑点提示
4. ✅ README.md 进度状态已更新
5. ✅ 代码符合 gofmt 规范
6. ✅ 目录结构符合本规范（.md 在根目录，.go 在 code/）

---

## 6. 代码风格与最佳实践

### 6.1 基本规范
- 使用 `gofmt` 格式化代码
- 优先使用标准库，避免不必要的第三方依赖
- 示例代码必须简洁、工程化、可运行

### 6.2 命名规范
- 包名：小写，简洁，不使用下划线
- 公开函数/类型：首字母大写
- 私有函数/类型：首字母小写

### 6.3 注释规范
- 公开的包、函数、类型必须有注释
- 注释以名称开头，如：`// Package calc provides ...`

---

## 7. 新增：架构思路文档规范

### 7.1 docs/ 目录
创建 `docs/` 目录存放概念性、思路性的对比文档，这些文档可能没有对应的代码。

### 7.2 架构文档清单

| 文档 | 内容 |
|------|------|
| `java_go_design_patterns.md` | 设计模式对照 |
| `java_go_microservices.md` | 微服务组件对照 |
| `java_go_architecture.md` | 架构思路对比 |
| `java_go_concurrency.md` | 并发库对照 |

### 7.3 架构文档内容要求
- 可以是纯概念讲解，不要求代码
- 重点对比 Java 生态和 Go 生态的思路差异
- 结合微服务场景讲解

---

## 8. 工作方式（必须严格遵守）

### 8.1 Unit -&gt; Lesson -&gt; Section 三级架构
教学推进必须按照"Unit (大单元) -&gt; Lesson (课程) -&gt; Section (小节知识点)"的层级进行。

### 8.2 串联执行机制
每个 Lesson 必须提供 `code/run.go` 封装当前 Lesson 所有 Section 的演示，并确保在项目根目录的 `main.go` 中被导入和调用。

### 8.3 进度追踪
每完成一个 Lesson，必须更新项目根目录 `README.md` 的学习进度与路线图状态。

### 8.4 文档同步
每次执行完任何调整后，必须同时检查 `PROJECT_RULES.md`、`study-plan.md` 和 `README.md`，确保内容一致。

---

## 9. 输出语言

默认中文；专业术语可中英并列。每次修改代码后，必须主动执行 `go test ./...` 和 `go run .` 并告知结果。

---

**本规范为项目开发的单一真相源，所有开发工作必须严格遵守。**
