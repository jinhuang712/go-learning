package lesson_02_channel

import (
	"fmt"
	"sync"
	"testing"
)

type ChannelPractice struct{}

func (p *ChannelPractice) Exercise1_AlternatePrint() {
	var wg sync.WaitGroup
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-ch1
			fmt.Print("A")
			ch2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-ch2
			fmt.Print("B")
			if i < 4 {
				ch1 <- struct{}{}
			}
		}
	}()

	ch1 <- struct{}{}
	wg.Wait()
	fmt.Println()
}

func TestChannelPractice(t *testing.T) {
	p := &ChannelPractice{}
	p.Exercise1_AlternatePrint()
}
