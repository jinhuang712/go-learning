package lesson_03_structs_interfaces

import "fmt"

// Run 运行结构体与接口课程代码
func Run() {
	fmt.Println("\n=== Lesson 3: Structs & Interfaces ===")
	
	// 运行结构体相关演示 (在 user.go 和 admin.go 中定义)
	demoStructs()

	// 运行接口底层相关演示
	demoInterfaces()
}

func demoStructs() {
	fmt.Println("\n--- 01: Structs & Methods ---")
	
	// 方式 1: 字面量初始化 (推荐)
	user := User{
		Name: "Alice",
		Age:  30,
	}
	fmt.Printf("User: %+v\n", user)

	// 方式 2: 通过模拟构造函数初始化
	user2 := NewUser("Bob", 25)
	fmt.Printf("User2 (from NewUser): %+v\n", user2)

	// 值接收者方法
	user.Introduce()
	
	// 指针接收者方法
	user.Birthday() 
	fmt.Printf("User after birthday: %+v\n", user)

	fmt.Println("\n--- Composition ---")
	admin := Admin{
		User: User{Name: "Bob", Age: 40},
		Role: "SuperAdmin",
	}
	
	fmt.Printf("Admin Name: %s (accessed directly)\n", admin.Name)
	admin.Introduce()

	fmt.Println("\n--- Anonymous Struct & Comparison ---")
	// 匿名结构体
	resp := struct {
		Code int
		Msg  string
	}{
		Code: 200,
		Msg:  "success",
	}
	fmt.Printf("Anonymous struct: %+v\n", resp)

	// 结构体比较
	type Point struct{ X, Y int }
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	fmt.Printf("p1 == p2? %t\n", p1 == p2)
}

func demoInterfaces() {
	fmt.Println("\n--- 02: Interface Internals ---")
	
	var speaker Speaker

	user := User{Name: "Alice", Age: 30}
	admin := Admin{
		User: User{Name: "Bob", Age: 40},
		Role: "SuperAdmin",
	}

	speaker = user
	fmt.Print("User as Speaker: ")
	speaker.Speak()

	speaker = admin
	fmt.Print("Admin as Speaker: ")
	speaker.Speak()

	// 接口的 nil 坑点演示
	fmt.Println("\n--- Interface Nil Pitfall ---")
	var uPtr *User = nil
	var anyInterface any = uPtr

	fmt.Printf("uPtr == nil? %t\n", uPtr == nil)
	fmt.Printf("anyInterface == nil? %t (Danger!)\n", anyInterface == nil)
	fmt.Printf("anyInterface type: %T, value: %v\n", anyInterface, anyInterface)
}
