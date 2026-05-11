package lesson_03_select_timeout

import (
	"testing"
	"time"
)

func TestSelectTimeout(t *testing.T) {
	ch := make(chan int)

	done := make(chan bool)
	go func() {
		select {
		case <-ch:
			t.Error("不应收到数据")
		case <-time.After(100 * time.Millisecond):
			done <- true
		}
	}()

	select {
	case <-done:
		// 超时正常
	case <-time.After(1 * time.Second):
		t.Fatal("超时机制失效")
	}
}

func TestNonBlocking(t *testing.T) {
	ch := make(chan int, 1)

	// 非阻塞发送应该成功
	select {
	case ch <- 1:
		// 成功
	default:
		t.Error("非阻塞发送失败")
	}

	// 非阻塞发送应该失败
	select {
	case ch <- 2:
		t.Error("不应发送成功")
	default:
		// 正常
	}
}
