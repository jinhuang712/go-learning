package lesson_02_channel_patterns

import (
	"testing"
	"time"
)

func TestChannelBasics(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()

	select {
	case val := <-ch:
		if val != 42 {
			t.Errorf("期望 42，实际 %d", val)
		}
	case <-time.After(time.Second):
		t.Fatal("超时")
	}
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	if len(ch) != 3 {
		t.Errorf("期望缓冲区长度 3，实际 %d", len(ch))
	}
}

func TestChannelClose(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)

	val, ok := <-ch
	if !ok {
		t.Error("期望 Channel 还有数据")
	}
	if val != 1 {
		t.Errorf("期望 1，实际 %d", val)
	}

	val, ok = <-ch
	if ok {
		t.Error("期望 Channel 已关闭")
	}
}
