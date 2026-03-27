# Go Learning

面向 Java 资深后端工程师的 Go 学习项目，目标是能在高并发、分布式微服务场景中独立阅读、设计、评审和优化 Go 代码。

## 学习路线图 (Roadmap)

以下是基于 `Unit -> Lesson -> Section` 三级架构规划的进阶学习路径，目标是达到大厂资深 Golang 微服务开发者水平。

### Unit 1: Go 基础与心智模型 (Fundamentals & Mental Model)
- [x] Lesson 1: Basics (基础语法、控制流、可见性)
- [x] Lesson 2: Collections (数组、切片与映射)
- [x] Lesson 3: Structs & Interfaces (结构体、方法、接口隐式实现与组合)
- [x] Lesson 4: Pointers & Value Semantics (指针与值语义)
- [ ] Lesson 5: Error Handling (错误处理、panic/recover、多返回值实战)

### Unit 2: 并发与并行核心 (Concurrency & Parallelism)
- [ ] Lesson 1: Goroutine & CSP 模型基础
- [ ] Lesson 2: Channel 模式与底层原理 (无缓冲/有缓冲、关闭、广播)
- [ ] Lesson 3: Select 多路复用与超时控制
- [ ] Lesson 4: Context 传递与级联取消
- [ ] Lesson 5: Sync 包 (Mutex, WaitGroup, Once, Map) 与 Atomic 原子操作
- [ ] Lesson 6: Race Detector 数据竞争检测实战

### Unit 3: 底层原理与性能调优 (Under the Hood & Performance)
- [ ] Lesson 1: G-M-P 调度器模型与阻塞原理
- [ ] Lesson 2: GC 垃圾回收机制与三色标记法
- [ ] Lesson 3: 内存逃逸分析 (Escape Analysis)
- [ ] Lesson 4: pprof 性能调优实战 (CPU, Memory, Goroutine 泄露排查)

### Unit 4: 工程化与测试 (Engineering & Testing)
- [ ] Lesson 1: Go 测试基础与 TDD 实战
- [ ] Lesson 2: 依赖注入与 Mock 测试
- [ ] Lesson 3: Benchmark 基准测试与性能对比
- [ ] Lesson 4: Fuzzing 模糊测试

### Unit 5: 大厂微服务实战 (Microservices in Production)
- [ ] Lesson 1: RPC 设计与 Protobuf/gRPC 基础
- [ ] Lesson 2: 中间件模式 (Middleware) 与请求拦截
- [ ] Lesson 3: 高可用架构：超时、重试与幂等性设计
- [ ] Lesson 4: 限流与熔断 (Rate Limiting & Circuit Breaking)
- [ ] Lesson 5: 可观测性 (Metrics, Tracing, Logging) 集成

## 复习方式

- 代码回放：`go run .`
- 单元复习测试：`go test ./...`
- 并发阶段开始后建议固定加入：`go test -race ./...`
