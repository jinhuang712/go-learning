# 4. 并发模型坑点

Goroutine和Channel是Go并发模型的核心，但Java开发者习惯了线程模型，转Go时很容易踩Goroutine泄露、Channel操作不当等坑。

## 坑点1：Goroutine泄露
```go
// ❌ 错误写法：Goroutine泄露，select里没有default或超时
func leakGoroutine() {
    ch := make(chan int)
    go func() {
        // 这个Goroutine会永远等待ch的读取，永远不会结束！
        ch <- 100
    }()
    // 没有读取ch，Goroutine永远阻塞
    return
}
```

✅ 正确做法：用context控制Goroutine生命周期，或者用超时
```go
func noLeakGoroutine(ctx context.Context) {
    ch := make(chan int, 1) // 用缓冲channel，避免发送方阻塞
    go func() {
        ch <- 100
    }()
    
    select {
    case <-ctx.Done():
        return
    case val := <-ch:
        fmt.Println("received:", val)
    }
}
```

## 坑点2：Channel关闭不当导致panic
```go
// ❌ 错误写法1：关闭已关闭的channel
func closeClosedChannel() {
    ch := make(chan int)
    close(ch)
    close(ch) // panic: close of closed channel
}

// ❌ 错误写法2：向已关闭的channel发送数据
func sendToClosedChannel() {
    ch := make(chan int)
    close(ch)
    ch <- 1 // panic: send on closed channel
}
```

✅ 正确做法：由发送方关闭channel，接收方不要关闭channel
```go
func safeChannel() {
    ch := make(chan int)
    // 发送方负责关闭
    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
        }
        close(ch) // 发送完成后关闭
    }()
    // 接收方用range读取，channel关闭后自动退出循环
    for val := range ch {
        fmt.Println("received:", val)
    }
}
```

## 坑点3：数据竞争（Race Condition）
Java有synchronized，Go忘记加锁很容易出现数据竞争：
```go
// ❌ 错误写法：并发读写共享变量，不加锁
var count int

func unsafeIncrement() {
    for i := 0; i < 10000; i++ {
        count++ // 并发读写，有数据竞争
    }
}

func main() {
    go unsafeIncrement()
    go unsafeIncrement()
    time.Sleep(1 * time.Second)
    fmt.Println(count) // 大概率不是20000
}
```

✅ 正确做法：用sync.Mutex或sync.RWMutex加锁，或者用atomic
```go
var count int
var mu sync.Mutex

func safeIncrement() {
    for i := 0; i < 10000; i++ {
        mu.Lock()
        count++
        mu.Unlock()
    }
}
```

## 练习题
1. 修复leakGoroutine函数，防止Goroutine泄露
2. 用sync.Mutex和atomic两种方式修复数据竞争问题
3. 用`go test -race ./...`检查上面的代码是否有数据竞争