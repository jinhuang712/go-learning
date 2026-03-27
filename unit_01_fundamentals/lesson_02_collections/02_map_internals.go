package lesson_02_collections

import "fmt"

// DemoMapInternals 演示 Map 的初始化、查找与遍历特性
func DemoMapInternals() {
	fmt.Println("\n--- 02: Map Internals ---")

	// 1. 初始化
	// var m map[string]int // 这是 nil map，写数据会 panic
	m := make(map[string]int)
	m["Alice"] = 95
	m["Bob"] = 88

	// 2. 查找与 Comma-ok 模式
	score, ok := m["Charlie"] // Charlie 不存在，返回 int 的零值 0
	fmt.Printf("Charlie's score: %d, exists: %t\n", score, ok)

	// 典型的 Go 判断写法
	if val, exists := m["Alice"]; exists {
		fmt.Printf("Alice is in the map with score: %d\n", val)
	}

	// 3. 删除元素
	delete(m, "Bob")
	fmt.Printf("After deleting Bob, map: %v\n", m)

	// 4. 遍历是无序的 (你可以多次运行，观察输出顺序)
	techStack := map[string]string{
		"Language": "Go",
		"Cloud":    "Kubernetes",
		"DB":       "MySQL",
		"Cache":    "Redis",
	}
	
	fmt.Println("Iterating map (Order is random):")
	for k, v := range techStack {
		fmt.Printf("  %s -> %s\n", k, v)
	}
}
