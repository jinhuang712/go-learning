package lesson_02_collections

import (
	"fmt"
	"sync"
)

// SafeCache 演示方案 A: 原生 map + 读写锁
type SafeCache struct {
	mu sync.RWMutex
	m  map[string]string
}

// NewSafeCache 初始化 SafeCache
func NewSafeCache() *SafeCache {
	return &SafeCache{
		m: make(map[string]string),
	}
}

// Set 写操作加写锁 (独占)
func (c *SafeCache) Set(key, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = val
}

// Get 读操作加读锁 (共享)
func (c *SafeCache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

// DemoMapInternals 演示 Map 的初始化、查找与遍历特性
func DemoMapInternals() {
	fmt.Println("\n--- 02: Map Internals ---")

	// 1. 初始化
	// var m map[string]int // 这是 nil map，写数据会 panic
	m := make(map[string]int)
	m["Alice"] = 95
	m["Bob"] = 88

	// 2. 查找与 Comma-ok 模式
	score, ok := m["Charlie"] // Charlie 不存在，返回 int 的零值 0
	fmt.Printf("Charlie's score: %d, exists: %t\n", score, ok)

	// 典型的 Go 判断写法
	if val, exists := m["Alice"]; exists {
		fmt.Printf("Alice is in the map with score: %d\n", val)
	}

	// 3. 删除元素
	delete(m, "Bob")
	fmt.Printf("After deleting Bob, map: %v\n", m)

	// 4. 遍历是无序的 (你可以多次运行，观察输出顺序)
	techStack := map[string]string{
		"Language": "Go",
		"Cloud":    "Kubernetes",
		"DB":       "MySQL",
		"Cache":    "Redis",
	}
	
	fmt.Println("Iterating map (Order is random):")
	for k, v := range techStack {
		fmt.Printf("  %s -> %s\n", k, v)
	}

	// 5. 并发安全演示 (SafeCache)
	fmt.Println("\n--- Safe Map with sync.RWMutex ---")
	cache := NewSafeCache()
	cache.Set("status", "running")
	if val, ok := cache.Get("status"); ok {
		fmt.Printf("SafeCache got: %s\n", val)
	}

	// 6. 并发安全演示 (sync.Map)
	fmt.Println("\n--- Safe Map with sync.Map ---")
	var sm sync.Map
	sm.Store("framework", "Gin") // 存入不需要强类型
	
	// 读取时返回的是 any (interface{})，需要类型断言
	if valAny, ok := sm.Load("framework"); ok {
		// valAny.(string) 就是类型断言，把 any 转回 string
		valStr := valAny.(string) 
		fmt.Printf("sync.Map got: %s\n", valStr)
	}
}
