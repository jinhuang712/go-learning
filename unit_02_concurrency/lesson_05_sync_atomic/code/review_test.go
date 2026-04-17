package lesson_05_sync_atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

type SyncAtomicPractice struct{}

func (p *SyncAtomicPractice) Exercise1_SafeCounter() int64 {
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

func TestSyncAtomicPractice(t *testing.T) {
	p := &SyncAtomicPractice{}
	result := p.Exercise1_SafeCounter()
	if result != 1000 {
		t.Errorf("Expected 1000, got %d", result)
	}
}
