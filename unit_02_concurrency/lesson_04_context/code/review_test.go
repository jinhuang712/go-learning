package lesson_04_context

import (
	"context"
	"testing"
	"time"
)

type ContextPractice struct{}

func (p *ContextPractice) Exercise1_CascadeCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})

	go func() {
		select {
		case <-ctx.Done():
			println("Goroutine 收到取消信号")
		}
		close(done)
	}()

	time.Sleep(100 * time.Millisecond)
	cancel()
	<-done
}

func TestContextPractice(t *testing.T) {
	p := &ContextPractice{}
	p.Exercise1_CascadeCancel()
}
