package lesson_06_java_go_pitfalls

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 1. 值/指针误用坑点示例

type UserValue struct {
	Name string
	Age  int
}

// ❌ 错误写法：值接收者修改状态无效
func (u UserValue) SetAgeWrong(age int) {
	u.Age = age
}

// ✅ 正确写法：指针接收者修改状态
func (u *UserValue) SetAgeCorrect(age int) {
	u.Age = age
}

type BigStruct struct {
	Field1 [1024]int
	Field2 [1024]string
}

func processValue(s BigStruct) {
	// 大结构体传值，拷贝开销大
}

func processPointer(s *BigStruct) {
	// 大结构体传指针，开销小
}

type UserRequest struct {
	Age *int // 用指针区分零值和未传
}

// 2. 接口nil判断坑点示例

type CustomError struct {
	Code int
	Msg  string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

// ❌ 错误写法：返回nil的具体类型指针
func GetUserWrong(id int64) error {
	var err *CustomError
	if id <= 0 {
		err = &CustomError{Code: 400, Msg: "invalid id"}
	}
	return err
}

// ✅ 正确写法：直接返回nil
func GetUserCorrect(id int64) error {
	if id <= 0 {
		return &CustomError{Code: 400, Msg: "invalid id"}
	}
	return nil
}

// 3. 错误处理坑点示例

// ❌ 错误写法：用panic代替error
func GetUserPanic(id int64) *UserValue {
	if id <= 0 {
		panic("invalid id")
	}
	return &UserValue{Name: "Bob", Age: 30}
}

// ✅ 正确写法：返回error
func GetUserNoError(id int64) (*UserValue, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid id: %d", id)
	}
	return &UserValue{Name: "Bob", Age: 30}, nil
}

// 4. 并发模型坑点示例

var count int
var mu sync.Mutex

// ❌ 错误写法：Goroutine泄露
func leakGoroutine() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	// 没有读取ch，Goroutine永远阻塞
}

// ✅ 正确写法：Goroutine不泄露
func noLeakGoroutine(ctx context.Context) {
	ch := make(chan int, 1)
	go func() {
		ch <- 100
	}()

	select {
	case <-ctx.Done():
		return
	case val := <-ch:
		fmt.Println("noLeakGoroutine received:", val)
	}
}

// ❌ 错误写法：数据竞争
func unsafeIncrement() {
	for i := 0; i < 10000; i++ {
		count++
	}
}

// ✅ 正确写法：加锁避免数据竞争
func safeIncrement() {
	for i := 0; i < 10000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
}

// RunPitfalls 运行所有坑点示例
func RunPitfalls() {
	fmt.Println("--- 1. 值/指针误用坑点 ---")
	u := UserValue{Name: "Bob", Age: 20}
	u.SetAgeWrong(30)
	fmt.Println("SetAgeWrong (value receiver):", u.Age)
	u.SetAgeCorrect(30)
	fmt.Println("SetAgeCorrect (pointer receiver):", u.Age)

	fmt.Println("\n--- 2. 接口nil判断坑点 ---")
	errWrong := GetUserWrong(1)
	fmt.Println("GetUserWrong(1) == nil?", errWrong == nil)
	errCorrect := GetUserCorrect(1)
	fmt.Println("GetUserCorrect(1) == nil?", errCorrect == nil)

	fmt.Println("\n--- 3. 错误处理坑点 ---")
	user, err := GetUserNoError(1)
	if err != nil {
		fmt.Println("GetUserNoError error:", err)
	} else {
		fmt.Println("GetUserNoError success:", user)
	}

	fmt.Println("\n--- 4. 并发模型坑点 ---")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	noLeakGoroutine(ctx)

	count = 0
	go safeIncrement()
	go safeIncrement()
	time.Sleep(200 * time.Millisecond)
	fmt.Println("safeIncrement count:", count)
}
