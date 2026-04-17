# Section 4: 可见性规则

## 1. 知识点核心说明

### 可见性规则
Go 通过标识符的首字母大小写来控制可见性：

| 首字母 | 可见性 | 范围 |
|--------|--------|------|
| **大写** | Public | 任何包都可以访问 |
| **小写** | Private | 只能在当前包内访问 |

### 适用范围
可见性规则适用于：
- 函数
- 变量
- 常量
- 类型
- 结构体字段
- 接口方法

---

## 2. Java 与 Go 的对比说明

| 特性 | Java | Go |
|------|------|----|
| **Public** | `public` | 首字母大写 |
| **Private** | `private` | 首字母小写 |
| **Protected** | `protected` | 无类似机制 |
| **Package-Private** | 无修饰符 | 同小写（仅当前包） |
| **细粒度控制** | 方法级、字段级 | 仅通过首字母 |

### 代码对照

**Java**:
```java
package com.example;

public class User {
    public String name;        // 公开
    private int age;           // 私有
    protected String email;    // 受保护
    
    public void publicMethod() {}    // 公开方法
    private void privateMethod() {}  // 私有方法
}
```

**Go**:
```go
package user

type User struct {
    Name string  // 公开（大写）
    age  int     // 私有（小写）
}

func (u *User) PublicMethod() {}   // 公开方法
func (u *User) privateMethod() {} // 私有方法
```

---

## 3. 可运行代码示例

### 示例 1: 可见性规则演示

```go
// package: user
package user

type User struct {
    Name string  // 公开
    age  int     // 私有
}

func NewUser(name string, age int) *User {
    return &amp;User{Name: name, age: age}
}

func (u *User) GetAge() int {  // 公开方法访问私有字段
    return u.age
}

func (u *User) validate() {     // 私有方法
    // 验证逻辑
}

// package: main
package main

import "example.com/user"

func main() {
    u := user.NewUser("Alice", 30)
    
    u.Name = "Bob"           // OK: 访问公开字段
    // u.age = 25            // 编译错误：age 是私有的
    
    age := u.GetAge()        // OK: 通过公开方法访问
    // u.validate()           // 编译错误：validate 是私有的
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 忘记首字母大写导致无法访问**
```go
// package: utils
package utils

func helper() string {  // 小写，其他包无法访问
    return "help"
}

// package: main
package main

import "example.com/utils"

func main() {
    // utils.helper()  // 编译错误！
}

// 修复：首字母大写
func Helper() string {
    return "help"
}
```

⚠️ **坑点 2: 过度暴露内部实现**
```go
// 不推荐 - 暴露过多内部细节
type User struct {
    ID        string
    Password  string  // 不应该公开！
    Salt      string  // 不应该公开！
}

// 推荐 - 只暴露必要的 API
type User struct {
    ID string
    // password 和 salt 是私有的
}

func (u *User) SetPassword(pwd string) error { /* ... */ }
func (u *User) CheckPassword(pwd string) bool { /* ... */ }
```

⚠️ **坑点 3: 结构体字段小写导致 JSON 序列化失败**
```go
// 错误 - 小写字段不会被序列化
type User struct {
    name string  // 不会出现在 JSON 中
    age  int     // 不会出现在 JSON 中
}

// 正确 - 使用 tag 控制 JSON 字段名
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

---

## 5. 练习题

### 练习 4.1: 设计一个合理的 API
**目标**: 设计一个 `user` 包，包含 `User` 类型和相关方法，确保只暴露必要的公开 API，隐藏内部实现细节。

**验收标准**:
- `User` 结构体只暴露必要的公开字段
- 提供合理的公开方法
- 私有字段和方法得到保护
- 包含 JSON 序列化支持（如果需要）

### 练习 4.2: 理解可见性规则（概念题）
**问题**:
1. Go 为什么没有 `protected` 修饰符？这是设计缺陷还是有意为之？
2. Java 的 package-private 和 Go 的小写可见性有什么区别？

**验收标准**: 能够清晰解释这两个问题。
