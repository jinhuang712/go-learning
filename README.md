# Go Learning

面向 Java 资深后端工程师的 Go 学习项目，目标是能在高并发、分布式微服务场景中独立阅读、设计、评审和优化 Go 代码。

## 📋 项目规范

**重要**: 请先阅读 [PROJECT_RULES.md](./PROJECT_RULES.md) 了解项目开发规范和目录结构要求。

## 🏗️ 项目结构

```
go-learning/
├── PROJECT_RULES.md      # 项目开发规范（必读）
├── study-plan.md          # 学习计划
├── README.md              # 本文件
├── docs/                  # 架构思路对比文档
│   ├── java_go_design_patterns.md    # 设计模式对照
│   ├── java_go_microservices.md      # 微服务组件对照
│   ├── java_go_architecture.md       # 架构思路对比
│   └── java_go_concurrency.md        # 并发库对照
├── main.go                # 主程序入口
├── go.mod                 # Go Module 定义
├── pkg/                   # 公共包
├── unit_01_fundamentals/  # Unit 1: Go 基础
├── unit_02_concurrency/   # Unit 2: 并发
└── ...
```

**目录结构说明**:
- 每个 Lesson 下的 `.md` 文档直接放在 Lesson 根目录
- 所有 `.go` 文件放在 `code/` 子目录中

## 环境配置要求

- Go 版本 &gt;= 1.20
- 开发工具推荐：GoLand / VSCode + Go 插件
- 建议配置：开启 `gofmt` 自动格式化、`golangci-lint` 静态检查

## Java-Go 核心概念速查表

| Java 概念 | Go 对应概念 | 本质差异 |
|----------|------------|----------|
| Class 类 | Struct 结构体 + 方法 | Go没有继承，只有组合 |
| Interface 接口 | Interface 接口 | Go是隐式实现，鸭子类型 |
| extends 继承 | 匿名嵌入 (Anonymous Embedding) | 只是语法糖，不是父子类关系 |
| Thread 线程 | Goroutine 协程 | 用户态调度，栈动态扩容，开销极小 |
| Synchronized 同步锁 | sync.Mutex / sync.RWMutex | 更轻量，性能更高 |
| Future / CompletableFuture | Channel | CSP并发模型，同步通信 |
| Exception 异常 | error 接口 | 错误是普通值，必须显式处理 |
| NullPointerException | 空指针解引用 panic | 只有未初始化的指针才会触发 |
| 线程池 | Goroutine 池（可选，如ants） | Goroutine本身足够轻量，大部分场景不需要池 |

## 学习路线图 (Roadmap)

以下是基于 `Unit -&gt; Lesson -&gt; Section` 三级架构规划的进阶学习路径，目标是达到大厂资深 Golang 微服务开发者水平。

### Unit 0: Go Package 系统 (Package &amp; Module System)
- [x] [Lesson 0: Go Package 系统详解](./unit_01_fundamentals/lesson_00_packages/lesson_00_packages.md)

### Unit 1: Go 基础与心智模型 (Fundamentals &amp; Mental Model)
- [x] [Lesson 1: Basics (基础语法、控制流、可见性)](./unit_01_fundamentals/lesson_01_basics/lesson_01_basics.md)
- [x] [Lesson 2: Collections (数组、切片与映射)](./unit_01_fundamentals/lesson_02_collections/lesson_02_collections.md)
- [x] [Lesson 3: Structs &amp; Interfaces (结构体、方法、接口隐式实现与组合)](./unit_01_fundamentals/lesson_03_structs_interfaces/lesson_03_structs_interfaces.md)
- [x] [Lesson 4: Pointers &amp; Value Semantics (指针与值语义)](./unit_01_fundamentals/lesson_04_pointers/lesson_04_pointers.md)
- [x] [Lesson 5: Error Handling (错误处理、panic/recover、多返回值实战)](./unit_01_fundamentals/lesson_05_error_handling/lesson_05_error_handling.md)
- [x] [Lesson 6: Java 转 Go 常见坑点专项](./unit_01_fundamentals/lesson_06_java_go_pitfalls/lesson_06_java_go_pitfalls.md)

### Unit 2: 并发与并行核心 (Concurrency &amp; Parallelism)
- [x] [Lesson 1: Goroutine &amp; CSP 模型基础](./unit_02_concurrency/lesson_01_goroutine/lesson_01_goroutine.md)
- [ ] Lesson 2: Channel 模式与底层原理 (无缓冲/有缓冲、关闭、广播)
- [ ] Lesson 3: Select 多路复用与超时控制
- [ ] Lesson 4: Context 传递与级联取消
- [ ] Lesson 5: Sync 包 (Mutex, WaitGroup, Once, Map) 与 Atomic 原子操作
- [ ] Lesson 6: Race Detector 数据竞争检测实战

### Unit 3: 底层原理与性能调优 (Under the Hood &amp; Performance)
- [ ] Lesson 1: G-M-P 调度器模型与阻塞原理
- [ ] Lesson 2: GC 垃圾回收机制与三色标记法
- [ ] Lesson 3: 内存逃逸分析 (Escape Analysis)
- [ ] Lesson 4: pprof 性能调优实战 (CPU, Memory, Goroutine 泄露排查)

### Unit 4: 工程化与测试 (Engineering &amp; Testing)
- [ ] Lesson 1: Go 测试基础与 TDD 实战
- [ ] Lesson 2: 依赖注入与 Mock 测试
- [ ] Lesson 3: Benchmark 基准测试与性能对比
- [ ] Lesson 4: Fuzzing 模糊测试

### Unit 5: 大厂微服务实战 (Microservices in Production)
- [ ] Lesson 1: RPC 设计与 Protobuf/gRPC 基础
- [ ] Lesson 2: 中间件模式 (Middleware) 与请求拦截
- [ ] Lesson 3: 高可用架构：超时、重试与幂等性设计
- [ ] Lesson 4: 限流与熔断 (Rate Limiting &amp; Circuit Breaking)
- [ ] Lesson 5: 可观测性 (Metrics, Tracing, Logging) 集成

## 复习方式

- 代码回放：`go run .`
- 单元复习测试：`go test ./...`
- 并发阶段开始后建议固定加入：`go test -race ./...`
