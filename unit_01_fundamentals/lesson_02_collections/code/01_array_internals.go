package lesson_02_collections

import "fmt"

// modifyArray 演示数组的值传递特性
func modifyArray(a [3]int) {
	a[0] = 999
}

// DemoArrayInternals 演示数组的声明与值传递
func DemoArrayInternals() {
	fmt.Println("--- 01: Array Internals ---")

	// 1. 数组声明与初始化
	var arr1 [3]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5}

	fmt.Printf("arr1: %v, type: %T\n", arr1, arr1)
	fmt.Printf("arr2: %v, type: %T\n", arr2, arr2)
	fmt.Printf("arr3: %v, type: %T\n", arr3, arr3)

	// 2. 数组是值传递
	fmt.Printf("Before modifyArray: %v\n", arr2)
	modifyArray(arr2)
	fmt.Printf("After modifyArray : %v (No change!)\n", arr2)
}
