package lesson_03_select

import "fmt"

func Run() {
	fmt.Println("\n=== Lesson 3: Select 多路复用与超时控制 ===")

	DemoSelectBasics()
	DemoSelectPatterns()
	DemoSelectPitfalls()
}
