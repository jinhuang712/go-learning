package lesson_01_gmp_scheduler

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

type GMPPractice struct{}

func (p *GMPPractice) Exercise1_GoroutineCount() int {
	var wg sync.WaitGroup
	const numGoroutines = 100

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			runtime.Gosched()
		}()
	}
	wg.Wait()

	return runtime.NumGoroutine()
}

func TestGMPPractice(t *testing.T) {
	p := &GMPPractice{}
	result := p.Exercise1_GoroutineCount()
	t.Logf("Goroutine count: %d", result)
}
