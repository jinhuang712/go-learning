package collections

import "fmt"

// Run 运行集合类型课程代码
func Run() {
	fmt.Println("\n=== Lesson 2: Arrays, Slices, and Maps ===")

	// 1. 数组 (Array)
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Printf("Array: %v, Length: %d\n", arr, len(arr))

	// 2. 切片 (Slice)
	// 创建方式 A: make
	slice1 := make([]string, 0, 5)
	slice1 = append(slice1, "Java")
	slice1 = append(slice1, "Go")
	slice1 = append(slice1, "K8s")
	fmt.Printf("Slice1: %v, Len: %d, Cap: %d\n", slice1, len(slice1), cap(slice1))

	// 创建方式 B: 字面量
	nums := []int{10, 20, 30, 40, 50}

	// 切片操作
	subSlice := nums[1:3] // [20, 30]
	fmt.Printf("Original: %v, SubSlice[1:3]: %v\n", nums, subSlice)

	// 修改切片影响底层数组
	subSlice[0] = 999
	fmt.Printf("After modification - Original: %v (See index 1 changed!)\n", nums)

	// 3. Map (映射)
	// 创建
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 88

	// 字面量
	techStack := map[string]string{
		"Language": "Go",
		"Cloud":    "Kubernetes",
	}
	fmt.Printf("Map: %v\n", techStack)

	// 检查 Key
	score, exists := scores["Charlie"]
	if exists {
		fmt.Printf("Charlie's score: %d\n", score)
	} else {
		fmt.Println("Charlie not found")
	}

	// 遍历
	for k, v := range techStack {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
}
