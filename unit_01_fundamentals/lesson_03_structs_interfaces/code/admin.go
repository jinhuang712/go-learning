package lesson_03_structs_interfaces

import "fmt"

// Admin 嵌入了 User
type Admin struct {
	User // 匿名嵌入 (Anonymous Embedding)
	Role string
}

// Introduce 重写 (Shadowing) User 的 Introduce 方法
// Java: @Override public void introduce() { ... }
func (a Admin) Introduce() {
	fmt.Printf("Hi, I'm Admin %s. Role: %s\n", a.Name, a.Role)
}
