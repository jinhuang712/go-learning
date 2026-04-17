package lesson_01_goroutine

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type GoroutinePractice struct{}

func (p *GoroutinePractice) Exercise1_ConcurrentCounter() int64 {
	var counter int64
	var wg sync.WaitGroup
	const numGoroutines = 1000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()

	return counter
}

func (p *GoroutinePractice) Exercise2_FixLeak(f func()) {
	f()
}

func TestExercise1_ConcurrentCounter(t *testing.T) {
	p := &GoroutinePractice{}
	result := p.Exercise1_ConcurrentCounter()

	if result != 1000 {
		t.Errorf("Expected 1000, got %d", result)
	}
}

func TestGoroutineLeakFix(t *testing.T) {
	p := &GoroutinePractice{}

	initialGoroutines := runtime.NumGoroutine()

	done := make(chan struct{})
	p.Exercise2_FixLeak(func() {
		go func() {
			<-done
		}()
	})

	close(done)
	time.Sleep(100 * time.Millisecond)

	finalGoroutines := runtime.NumGoroutine()
	if finalGoroutines > initialGoroutines+10 {
		t.Errorf("Possible goroutine leak: initial=%d, final=%d", initialGoroutines, finalGoroutines)
	}
}
