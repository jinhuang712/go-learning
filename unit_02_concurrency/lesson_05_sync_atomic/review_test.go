package lesson_05_sync_atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestMutex(t *testing.T) {
	counter := &SafeCounter{}
	var wg sync.WaitGroup

	const iterations = 1000
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	if counter.Get() != iterations {
		t.Errorf("期望 %d，实际 %d", iterations, counter.Get())
	}
}

func TestAtomic(t *testing.T) {
	var count int64
	var wg sync.WaitGroup

	const iterations = 1000
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)
		}()
	}

	wg.Wait()
	if count != iterations {
		t.Errorf("期望 %d，实际 %d", iterations, count)
	}
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var count int

	for i := 0; i < 10; i++ {
		once.Do(func() {
			count++
		})
	}

	if count != 1 {
		t.Errorf("期望执行 1 次，实际 %d 次", count)
	}
}
