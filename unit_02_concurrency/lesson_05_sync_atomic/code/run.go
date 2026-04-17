package lesson_05_sync_atomic

import "fmt"

func Run() {
	fmt.Println("\n=== Lesson 5: Sync 包与 Atomic 原子操作 ===")

	DemoSyncMutex()
	DemoSyncOther()
	DemoAtomic()
}
