package lesson_02_channel_patterns

import (
	"fmt"
	"sync"
	"time"
)

// DemoProducerConsumer 演示生产者-消费者模式
func DemoProducerConsumer() {
	fmt.Println("\n--- Section 6: 生产者-消费者模式 ---")

	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	// 启动 2 个消费者
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// 生产者发送任务
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
			fmt.Printf("  发送任务: %d\n", j)
		}
		close(jobs)
	}()

	// 收集结果
	go func() {
		wg.Wait()
		close(results)
	}()

	// 输出结果
	for result := range results {
		fmt.Printf("  收到结果: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("  Worker %d 处理任务: %d\n", id, j)
		time.Sleep(100 * time.Millisecond)
		results <- j * 2
	}
}
