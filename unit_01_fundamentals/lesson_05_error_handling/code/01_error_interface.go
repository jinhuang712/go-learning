package lesson_05_error_handling

import (
	"errors"
	"fmt"
)

// User 模拟数据库实体
type User struct {
	ID   int
	Name string
}

// 1. 自定义错误类型 (模拟业务错误)
// Java 对标: class UserNotFoundException extends RuntimeException
type NotFoundError struct {
	Resource string
	ID       int
}

// 只要实现了 Error() string，NotFoundError 就隐式实现了 error 接口
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Error [404]: %s with ID %d not found", e.Resource, e.ID)
}

// 2. 模拟微服务中的查找逻辑，返回 (结果, error)
func findUser(id int) (*User, error) {
	if id <= 0 {
		// 返回标准库的简单错误
		return nil, errors.New("id must be strictly positive")
	}

	if id != 1 {
		// 返回自定义错误 (注意要返回指针，因为接收者可以是指针)
		return nil, &NotFoundError{Resource: "User", ID: id}
	}

	// 成功时，error 返回 nil
	return &User{ID: 1, Name: "Alice"}, nil
}

// DemoErrorInterface 运行 01 的演示
func DemoErrorInterface() {
	fmt.Println("--- 01: Error Interface ---")

	// 场景 A: 正常流
	user, err := findUser(1)
	if err != nil {
		fmt.Println("Failed:", err)
	} else {
		fmt.Printf("Success: found user %s\n", user.Name)
	}

	// 场景 B: 触发 NotFoundError
	_, err = findUser(2)
	if err != nil {
		// 这里 err 会打印出我们自定义的 Error() string
		fmt.Println("Failed:", err)
	}
}
