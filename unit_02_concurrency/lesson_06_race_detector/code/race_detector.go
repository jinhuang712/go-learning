package lesson_06_race_detector

import "fmt"

func DemoDataRace() {
	fmt.Println("\n--- 数据竞争演示 ---")
	fmt.Println("请查看 Section 1 文档了解详细内容")
	fmt.Println("注意：本示例有数据竞争，运行 go test -race 可以检测到")
}

func DemoRaceDetector() {
	fmt.Println("\n--- Race Detector 使用演示 ---")
	fmt.Println("请查看 Section 2 文档了解详细内容")
	fmt.Println("使用方式: go test -race ./...")
}

func DemoFixingRaces() {
	fmt.Println("\n--- 修复数据竞争演示 ---")
	fmt.Println("请查看 Section 3 文档了解详细内容")
	fmt.Println("修复方案包括：")
	fmt.Println("  1. 使用 Mutex")
	fmt.Println("  2. 使用 Channel")
	fmt.Println("  3. 使用 Atomic")
}
