package packages

import (
	"fmt"
	"go-learning/pkg/calc"
)

func DemoPackageBasics() {
	fmt.Println("\n--- Package 基础演示 ---")

	fmt.Println("导入并使用 pkg/calc 包:")
	result := calc.PublicFunc(10, 20)
	fmt.Printf("calc.PublicFunc(10, 20) = %d\n", result)
	fmt.Printf("calc.PublicConst = %d\n", calc.PublicConst)
}

func DemoImportMechanism() {
	fmt.Println("\n--- Import 机制演示 ---")
	fmt.Println("请查看 02_import_mechanism.md 文档了解不同的导入方式")
}

func DemoGoModules() {
	fmt.Println("\n--- Go Module 系统演示 ---")
	fmt.Println("当前项目的 module 路径在 go.mod 中定义")
	fmt.Println("请查看 03_go_modules.md 文档了解详情")
}

func DemoVisibilityRules() {
	fmt.Println("\n--- 可见性规则演示 ---")
	fmt.Println("首字母大写 = Public，首字母小写 = Private")
	fmt.Println("请查看 04_visibility_rules.md 文档了解详情")
}
