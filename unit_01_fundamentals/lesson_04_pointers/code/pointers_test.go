package pointers

import "testing"

func TestPassByPointer(t *testing.T) {
	val := 10
	passByPointer(&val)
	if val != 99 {
		t.Fatalf("expected val to be 99, got %d", val)
	}
}

// swap 练习题：请实现这个函数，利用指针交换 x 和 y 的值
func swap(x, y *int) {
	// TODO: 实现交换逻辑
	// 提示：你需要一个临时变量来保存解引用后的值
	temp := *x
	*x = *y
	*y = temp
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	
	// 调用 swap，你需要传入 a 和 b 的地址
	swap(&a, &b)

	if a != 2 || b != 1 {
		t.Fatalf("expected a=2, b=1, but got a=%d, b=%d", a, b)
	}
}
