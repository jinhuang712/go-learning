# Lesson 3: 结构体与接口 (Structs & Interfaces)

本课程主要介绍 Go 的面向对象编程 (OOP) 核心：结构体、方法、接口以及组合。

## 核心知识点

### 1. 结构体 (Struct)
Go 没有 `class`，只有 `struct`。
- **定义**: `type User struct { ... }`
- **实例化**: `u := User{Name: "Alice"}`

### 2. 方法 (Method)
方法是绑定到特定类型上的函数。
- **值接收者 (Value Receiver)**: `func (u User) Foo()`
  - 类似于传值调用，方法内修改不会影响原对象。
- **指针接收者 (Pointer Receiver)**: `func (u *User) Bar()`
  - 类似于传递引用，方法内可以修改原对象。
  - **规则**: 如果需要修改状态，必须用指针接收者。

### 3. 组合与嵌入 (Composition & Embedding)
Go 提倡**组合优于继承**。
- **嵌入**: 在 Struct 中匿名包含另一个 Struct。
  ```go
  type Admin struct {
      User // 嵌入 User
      Role string
  }
  ```
- **效果**: `Admin` 可以直接访问 `User` 的字段 (`admin.Name`) 和方法 (`admin.Introduce()`)，类似于继承，但本质是组合。

### 4. 接口 (Interface)
Go 的接口是**隐式实现**的 (Duck Typing)。
- **定义**: `type Speaker interface { Speak() }`
- **实现**: 只要类型 `T` 实现了 `Speak()` 方法，它就自动实现了 `Speaker` 接口，无需 `implements` 关键字。
- **Java 对比**: 
  - Java: `class User implements Speaker` (显式)
  - Go: `func (u User) Speak() { ... }` (隐式)

## 练习
1. 给 `Admin` 添加一个新的方法 `BanUser()`。
2. 定义一个新的结构体 `Bot`，也实现 `Speaker` 接口，并放入切片 `[]Speaker` 中进行统一遍历调用。
