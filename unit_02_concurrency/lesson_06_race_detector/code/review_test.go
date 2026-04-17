package lesson_06_race_detector

import (
	"sync"
	"testing"
)

type RaceDetectorPractice struct{}

func (p *RaceDetectorPractice) Exercise1_FixRace() int {
	var counter int
	var wg sync.WaitGroup
	const numGoroutines = 1000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			counter++ // 这里有数据竞争！
		}()
	}
	wg.Wait()

	return counter
}

func TestRaceDetectorPractice(t *testing.T) {
	p := &RaceDetectorPractice{}
	result := p.Exercise1_FixRace()
	t.Logf("Result: %d (note: this test has data race)", result)
}
