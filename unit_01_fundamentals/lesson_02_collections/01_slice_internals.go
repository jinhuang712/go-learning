package lesson_02_collections

import "fmt"

// DemoSliceInternals 演示切片的底层共享与扩容机制
func DemoSliceInternals() {
	fmt.Println("--- 01: Slice Internals ---")

	// 1. 数组 vs 切片
	arr := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}
	fmt.Printf("Array type: %T, Slice type: %T\n", arr, slice)

	// 2. 切片截取与底层数组共享
	nums := []int{10, 20, 30, 40, 50}
	sub := nums[1:3] // 截取 index 1 和 2: [20, 30]
	fmt.Printf("Before modify: nums=%v, sub=%v\n", nums, sub)

	// 修改 sub 会影响 nums，因为它们共享底层数组
	sub[0] = 999 
	fmt.Printf("After modify sub[0]: nums=%v, sub=%v\n", nums, sub)

	// 3. Append 与扩容机制
	fmt.Println("\n--- Append & Capacity ---")
	// make([]type, len, cap)
	s := make([]int, 0, 3) 
	fmt.Printf("Initial  : len=%d, cap=%d, ptr=%p\n", len(s), cap(s), s)

	// 不触发扩容
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Printf("Appended 3: len=%d, cap=%d, ptr=%p\n", len(s), cap(s), s)

	// 触发扩容 (cap 从 3 变成 6)
	s = append(s, 4)
	fmt.Printf("Appended 4: len=%d, cap=%d, ptr=%p (Pointer changed!)\n", len(s), cap(s), s)
}
