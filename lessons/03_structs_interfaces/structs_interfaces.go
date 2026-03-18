package structs_interfaces

import "fmt"

// Speaker 接口定义
// Java: public interface Speaker { void speak(); }
type Speaker interface {
	Speak()
}

// Run 运行结构体与接口课程代码
func Run() {
	fmt.Println("\n=== Lesson 3: Structs, Methods, and Interfaces ===")

	// 1. 结构体 (Struct)
	// Java: User user = new User("Alice", 30);
	user := User{
		Name: "Alice",
		Age:  30,
	}
	fmt.Printf("User: %+v\n", user)

	// 2. 方法 (Method) - 值接收者 vs 指针接收者
	// 调用值接收者方法 (不修改原对象)
	user.Introduce()
	
	// 调用指针接收者方法 (修改原对象)
	// Java: user.setAge(31);
	user.Birthday() 
	fmt.Printf("User after birthday: %+v\n", user)

	// 3. 组合与嵌入 (Composition & Embedding)
	// Java: public class Admin extends User { ... }
	// Go: 使用结构体嵌入模拟继承
	admin := Admin{
		User: User{Name: "Bob", Age: 40}, // 初始化嵌入的结构体
		Role: "SuperAdmin",
	}
	
	// 可以直接访问 User 的字段和方法 (Promoted fields/methods)
	fmt.Printf("Admin Name: %s (accessed directly)\n", admin.Name)
	admin.Introduce() // 调用 User 的 Introduce 方法

	// 4. 接口 (Interface)
	// Go 的接口是隐式实现的 (Duck Typing)
	// 只要实现了接口的所有方法，就自动实现了该接口
	var speaker Speaker

	speaker = user
	fmt.Print("User as Speaker: ")
	speaker.Speak()

	speaker = admin
	fmt.Print("Admin as Speaker: ")
	speaker.Speak() // Admin 继承了 User 的 Speak 方法 (如果有) 或者实现了自己的
}
