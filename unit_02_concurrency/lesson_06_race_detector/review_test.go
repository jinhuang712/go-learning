package lesson_06_race_detector

import (
	"sync"
	"testing"
)

func TestNoRaceCondition(t *testing.T) {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	const iterations = 1000
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	if count != iterations {
		t.Errorf("期望 %d，实际 %d", iterations, count)
	}
}
