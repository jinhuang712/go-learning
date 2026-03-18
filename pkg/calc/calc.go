package calc

// PublicFunc 以大写字母开头
// 相当于 Java 的 public static int PublicFunc(...)
// 可以被其他包导入并使用
func PublicFunc(a, b int) int {
	return privateFunc(a, b) // 内部可以调用私有函数
}

// privateFunc 以小写字母开头
// 相当于 Java 的 private static int privateFunc(...) (或者 package-private)
// 只能在 calc 包内部使用，外部无法访问
func privateFunc(a, b int) int {
	return a + b
}

// Constant (常量) 和 Struct (结构体) 也是同样的规则

// PublicConst 外部可见
const PublicConst = 100

// privateConst 仅包内可见
const privateConst = 200
