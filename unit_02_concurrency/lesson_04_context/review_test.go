package lesson_04_context

import (
	"context"
	"testing"
	"time"
)

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan bool)
	go func() {
		<-ctx.Done()
		done <- true
	}()

	cancel()

	select {
	case <-done:
		// 正常
	case <-time.After(1 * time.Second):
		t.Fatal("Context 取消未生效")
	}
}

func TestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	select {
	case <-ctx.Done():
		if ctx.Err() != context.DeadlineExceeded {
			t.Errorf("期望 DeadlineExceeded，实际 %v", ctx.Err())
		}
	case <-time.After(1 * time.Second):
		t.Fatal("超时未触发")
	}
}

func TestContextValue(t *testing.T) {
	type key string
	testKey := key("test")

	ctx := context.WithValue(context.Background(), testKey, "hello")

	val := ctx.Value(testKey)
	if val != "hello" {
		t.Errorf("期望 hello，实际 %v", val)
	}
}
