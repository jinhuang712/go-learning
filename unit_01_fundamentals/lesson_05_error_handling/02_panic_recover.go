package lesson_05_error_handling

import "fmt"

// simulatePanic 模拟一个发生严重错误的函数
func simulatePanic() {
	fmt.Println("simulatePanic: Doing some work...")
	
	// 触发恐慌
	panic("simulatePanic: Oh no, a critical bug occurred (e.g. nil pointer)!")
	
	// 这行代码永远不会被执行
	// fmt.Println("simulatePanic: This will never print")
}

// safeCall 包装模拟恐慌的函数，保证程序不会因为它的恐慌而彻底崩溃
func safeCall() {
	// 1. 必须在 defer 内调用 recover
	// 相当于 Java 的 catch (Exception e)
	defer func() {
		if r := recover(); r != nil {
			// r 就是 panic 传入的值
			fmt.Printf("safeCall: Recovered from panic! Error was: %v\n", r)
		}
	}()

	fmt.Println("safeCall: Calling risky function...")
	
	// 2. 调用可能发生 panic 的函数
	simulatePanic()

	// 3. 发生 panic 后，这行也不会执行。
	// 控制流会直接跳到 defer 里，执行完 defer 后 safeCall() 直接返回。
	fmt.Println("safeCall: This will not print if panic occurs.")
}

// DemoPanicRecover 运行 02 的演示
func DemoPanicRecover() {
	fmt.Println("\n--- 02: Panic & Recover ---")
	
	fmt.Println("Main: Start")
	
	safeCall()
	
	// 因为 safeCall 内部 recover 了，所以外层可以继续正常执行
	fmt.Println("Main: Program survived and continues normally!")
}
