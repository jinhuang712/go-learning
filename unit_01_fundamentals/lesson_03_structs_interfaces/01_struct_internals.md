# 1. 结构体 (Structs) 深度剖析

Go 语言抛弃了传统的基于 Class 的面向对象体系，而是退回到了类似 C 语言的 `struct`。但这并不意味着 Go 不支持面向对象，它只是换了一种更轻量、更组合化的方式。

## 1.1 结构体的本质与内存布局
在 Java 中，对象是分配在堆上的，变量里存的只是对象的引用（指针）。
但在 Go 中，**`struct` 默认是值类型，它的字段在内存中是连续分配的**。

```go
type User struct {
    Age  int8   // 1 byte
    Role string // 16 bytes (pointer + len)
    Name string // 16 bytes
}
```
**内存对齐与性能**：
在极高并发的微服务中，结构体的字段顺序会影响其在内存中的大小（由于 CPU 内存对齐机制）。为了节省内存，最佳实践是**将占用空间相同或相近的字段放在一起，按字段大小从大到小排序**。

## 1.2 方法接收者 (Receiver)：值 vs 指针
方法本质上只是一个带有特殊参数（接收者）的普通函数。

- **值接收者 `(u User)`**:
  - 调用方法时，会把整个 `User` 结构体**拷贝**一份。
  - 在方法内修改 `u.Name`，**不会**影响原对象。
  - 适用场景：不需要修改状态的小对象，或者为了保证并发安全的不可变对象。

- **指针接收者 `(u *User)`**:
  - 调用时只拷贝指针（8字节），效率极高。
  - 在方法内修改 `u.Name`，**会真正修改原对象**。
  - 适用场景：需要修改对象状态、对象很大（避免拷贝开销）、包含 `sync.Mutex` 锁的结构体（锁绝对不能被拷贝！）。
  - **大厂规范**：为了统一行为，如果一个结构体有一个方法是指针接收者，那么**它的所有方法都应该是指针接收者**。

## 1.3 组合优于继承 (Composition over Inheritance)
Go 没有 `extends` 关键字，它使用**匿名嵌入 (Anonymous Embedding)** 来实现代码复用。

```go
type Admin struct {
    User // 匿名嵌入
    Level int
}
```
- **字段提升 (Promoted Fields)**：你可以直接通过 `admin.Name` 访问内嵌的 `User` 的字段，就像继承一样。
- **覆盖 (Shadowing)**：如果 `Admin` 也定义了 `Name` 字段或同名方法，它会屏蔽 `User` 的。
- **本质区别**：这只是语法糖！`Admin` 并不是 `User` 的子类，你**不能**把 `Admin` 对象传给需要 `User` 参数的函数。

## 1.4 结构体标签 (Struct Tags)
这是微服务中最常用的特性（JSON 序列化、ORM 映射、参数校验全靠它）。
标签是在运行时通过反射 (Reflection) 读取的元数据。

```go
type Request struct {
    // json: 序列化为 my_name, omitempty: 如果为空则不输出该字段
    // binding: 框架(如 Gin)层的必填校验
    Name string `json:"my_name,omitempty" binding:"required"`
    Age  int    `json:"-"` // "-" 表示在 JSON 序列化时绝对忽略该字段
}
```