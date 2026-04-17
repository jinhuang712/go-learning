package pointers

import "fmt"

// Config 模拟一个配置结构体
type Config struct {
	Port int
	Host string
}

// passByValue 尝试修改值，但这只会修改传入的"副本"
func passByValue(n int) {
	n = 99
}

// passByPointer 接收一个整型指针，通过地址修改原始数据
func passByPointer(p *int) {
	*p = 99 // 解引用：修改这个地址里存放的值
}

// updateConfigValue 传值，无法修改原结构体
func updateConfigValue(c Config) {
	c.Port = 9090
}

// updateConfigPointer 传指针，可以修改原结构体
func updateConfigPointer(c *Config) {
	// 这里的 c.Port 实际上是 (*c).Port 的语法糖
	c.Port = 9090
}

// Run 运行指针课程演示
func Run() {
	fmt.Println("\n=== Lesson 4: Pointers & Value Semantics ===")

	// 1. 基本指针操作
	x := 10
	fmt.Printf("Initial x: %d, Address of x: %p\n", x, &x)

	// 2. 值传递演示
	passByValue(x)
	fmt.Printf("After passByValue, x: %d (No change!)\n", x)

	// 3. 指针传递演示
	p := &x // p 的类型是 *int (指向 x 的内存地址)
	passByPointer(p)
	// 或者直接写 passByPointer(&x)
	fmt.Printf("After passByPointer, x: %d (Changed!)\n", x)

	// 4. 结构体与指针
	// 初始化结构体并直接获取它的指针 (非常常用的写法)
	cfg := &Config{
		Port: 8080,
		Host: "localhost",
	}
	
	// 传值 (失败)
	updateConfigValue(*cfg) // *cfg 将指针解引用为值传递进去
	fmt.Printf("After updateConfigValue, Port: %d\n", cfg.Port)

	// 传指针 (成功)
	updateConfigPointer(cfg)
	fmt.Printf("After updateConfigPointer, Port: %d\n", cfg.Port)
}
