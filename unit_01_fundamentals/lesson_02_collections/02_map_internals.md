# 2. 映射 (Map) 核心剖析与并发陷阱

Go 的 `map` 等同于 Java 中的 `HashMap`，但它在设计上有几个非常显著的特点和“坑点”，在微服务开发中极易中招。

## 2.1 Map 的声明与初始化
- `map` 是引用类型，**零值是 `nil`**。
- **坑点 1**: 往 `nil` 的 map 中写数据会导致程序直接崩溃 (`panic`)。
- **正确姿势**: 必须用 `make` 或者字面量初始化。
  ```go
  // 错误 (会 panic)
  var m1 map[string]int
  // m1["A"] = 1 

  // 正确
  m2 := make(map[string]int) // 也可以指定初始容量 make(..., 10)
  ```

## 2.2 查找与 "Comma ok" 模式
在 Java 中，如果 `map.get(key)` 找不到，会返回 `null`。
但在 Go 中，如果 key 不存在，**它会返回 Value 类型的零值**（比如 int 返回 0，string 返回 `""`）。

那么，你怎么知道这个 `0` 是因为没找到，还是因为别人存的就是 `0` 呢？
这就是 Go 特有的 `comma ok` 语法：
```go
val, ok := m["Alice"]
if !ok {
    fmt.Println("Alice 不在 map 里")
}
```

## 2.3 遍历的随机性
**坑点 2**: Go 的 `map` 遍历是**绝对无序的**（底层甚至故意加了随机种子）。
千万不要在代码中依赖 `for k, v := range m` 的输出顺序。如果需要有序，只能把 keys 拿出来放到切片里排序，再遍历切片。

## 2.4 并发不安全 (Fatal Error)
**最致命的坑**: Go 的原生 `map` **不支持并发读写**。
- 如果一个 goroutine 在写 map，另一个 goroutine 在读/写 map，Go 运行时会直接抛出致命错误 `fatal error: concurrent map read and map write`，这会导致整个微服务进程立刻挂掉，连 `recover` 都救不回来！
- **Java 对比**: 类似 Java 的 `HashMap`（也是线程不安全的），但在 Java 里最多是数据错乱或者死循环，Go 里是直接 kill 进程。
- **解决方案**:
  1. 使用 `sync.RWMutex` 自己加锁保护。
  2. 使用并发安全的 `sync.Map`（通常用于读多写少的缓存场景）。