package packages

import "fmt"

func Run() {
	fmt.Println("\n=== Lesson 0: Go Package 系统详解 ===")

	DemoPackageBasics()
	DemoImportMechanism()
	DemoGoModules()
	DemoVisibilityRules()
}
