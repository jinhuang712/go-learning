package lesson_05_error_handling

import (
	"errors"
	"testing"
)

// divide 练习题 1: 实现一个除法函数，如果 b 为 0，返回一个 errors.New 创建的错误
func divide(a, b int) (int, error) {
	// TODO: 实现除法并处理 b == 0 的情况
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func TestDivide(t *testing.T) {
	// 测试正常情况
	res, err := divide(10, 2)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if res != 5 {
		t.Fatalf("Expected 5, got: %d", res)
	}

	// 测试异常情况
	_, err = divide(10, 0)
	if err == nil {
		t.Fatal("Expected an error for division by zero, but got nil")
	}
}

// 练习题 2: 测试我们之前写的 findUser 自定义错误
func TestFindUserCustomError(t *testing.T) {
	_, err := findUser(2)
	
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	// 在 Go 中，我们可以用类型断言 (Type Assertion) 来检查 err 的具体类型
	// 类似 Java 的: if (err instanceof NotFoundError)
	customErr, ok := err.(*NotFoundError)
	if !ok {
		t.Fatalf("Expected error to be of type *NotFoundError, got %T", err)
	}

	if customErr.ID != 2 {
		t.Fatalf("Expected error ID to be 2, got %d", customErr.ID)
	}
}
