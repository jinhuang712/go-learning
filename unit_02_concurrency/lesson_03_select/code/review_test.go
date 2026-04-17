package lesson_03_select

import (
	"testing"
	"time"
)

type SelectPractice struct{}

func (p *SelectPractice) Exercise1_TimeoutRequest() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "请求完成"
	}()

	select {
	case result := <-ch:
		println("收到结果:", result)
	case <-time.After(100 * time.Millisecond):
		println("请求超时")
	}
}

func TestSelectPractice(t *testing.T) {
	p := &SelectPractice{}
	p.Exercise1_TimeoutRequest()
}
