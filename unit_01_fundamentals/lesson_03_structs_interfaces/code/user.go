package lesson_03_structs_interfaces

import "fmt"

// Speaker 接口定义 (为了解耦，通常定义在调用方，但这里为了演示放在这)
// Java: public interface Speaker { void speak(); }
type Speaker interface {
	Speak()
}

// User 定义一个结构体
// Java: public class User { public String name; public int age; ... }
type User struct {
	Name string
	Age  int
}

// NewUser 模拟构造函数 (大厂常用模式)
func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

// Introduce 是 User 的一个方法
// (u User) 是接收者 (Receiver)，类似于 Java 中的 this
// 这是一个 "值接收者" (Value Receiver)，通过拷贝调用，无法修改 u 的字段
func (u User) Introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", u.Name, u.Age)
}

// Birthday 是 User 的另一个方法
// (u *User) 是 "指针接收者" (Pointer Receiver)
// 可以修改 u 指向的对象
func (u *User) Birthday() {
	u.Age++
	fmt.Println("Happy Birthday! Age increased.")
}

// Speak 实现 Speaker 接口
func (u User) Speak() {
	fmt.Println("Hello!")
}
