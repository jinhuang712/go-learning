# Java-Go 并发库对照

## 前言
本文档对比 Java 的 `java.util.concurrent` 包和 Go 的 `sync`/`channel` 在并发编程中的差异。

---

## 1. 基本并发单元

| 特性 | Java | Go |
|------|------|----|
| 并发执行单元 | `Thread` | `Goroutine` |
| 启动方式 | `new Thread().start()` | `go func()` |
| 栈大小 | 默认 ~1MB | 默认 2KB，动态扩容 |
| 调度者 | 操作系统 | Go Runtime |

---

## 2. 互斥锁

### Java: ReentrantLock
```java
import java.util.concurrent.locks.ReentrantLock;

ReentrantLock lock = new ReentrantLock();

lock.lock();
try {
    // 临界区
} finally {
    lock.unlock();
}
```

### Go: sync.Mutex
```go
import "sync"

var mu sync.Mutex

mu.Lock()
defer mu.Unlock()  // defer 确保解锁
// 临界区
```

**读写锁**:
```java
// Java
ReentrantReadWriteLock rwLock = new ReentrantReadWriteLock();
rwLock.readLock().lock();
rwLock.writeLock().lock();
```

```go
// Go
var rwMu sync.RWMutex
rwMu.RLock()  // 读锁
rwMu.Lock()   // 写锁
```

---

## 3. 等待多个任务完成

### Java: CountDownLatch
```java
import java.util.concurrent.CountDownLatch;

CountDownLatch latch = new CountDownLatch(3);

for (int i = 0; i &lt; 3; i++) {
    new Thread(() -&gt; {
        try {
            // 工作
        } finally {
            latch.countDown();
        }
    }).start();
}

latch.await();  // 等待所有完成
```

### Go: sync.WaitGroup
```go
import "sync"

var wg sync.WaitGroup

wg.Add(3)
for i := 0; i &lt; 3; i++ {
    go func() {
        defer wg.Done()
        // 工作
    }()
}

wg.Wait()  // 等待所有完成
```

---

## 4. 只执行一次

### Java: 双重检查锁
```java
public class Singleton {
    private static volatile Singleton instance;
    
    public static Singleton getInstance() {
        if (instance == null) {
            synchronized (Singleton.class) {
                if (instance == null) {
                    instance = new Singleton();
                }
            }
        }
        return instance;
    }
}
```

### Go: sync.Once
```go
import "sync"

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &amp;Singleton{}
    })
    return instance
}
```

---

## 5. 条件变量

### Java: Condition
```java
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.ReentrantLock;

ReentrantLock lock = new ReentrantLock();
Condition notFull = lock.newCondition();
Condition notEmpty = lock.newCondition();

// 生产者
lock.lock();
try {
    while (queue.size() == capacity) {
        notFull.await();
    }
    queue.put(item);
    notEmpty.signal();
} finally {
    lock.unlock();
}
```

### Go: sync.Cond
```go
import "sync"

var (
    mu      sync.Mutex
    cond    = sync.NewCond(&amp;mu)
    queue   []int
    maxSize = 10
)

// 生产者
mu.Lock()
for len(queue) == maxSize {
    cond.Wait()
}
queue = append(queue, item)
cond.Signal()
mu.Unlock()
```

**Go 的另一种方式：使用 Channel**
```go
// Channel 通常比 Cond 更简单易用
queue := make(chan int, 10)

// 生产者
queue &lt;- item

// 消费者
item := &lt;-queue
```

---

## 6. 并发安全的 Map

### Java: ConcurrentHashMap
```java
import java.util.concurrent.ConcurrentHashMap;

ConcurrentHashMap&lt;String, String&gt; map = new ConcurrentHashMap&lt;&gt;();

map.put("key", "value");
String val = map.get("key");
```

### Go: sync.Map
```go
import "sync"

var m sync.Map

m.Store("key", "value")
val, ok := m.Load("key")
```

**Go 的另一种方式：Mutex + Map**
```go
var (
    mu sync.RWMutex
    m  = make(map[string]string)
)

// 读
mu.RLock()
val := m["key"]
mu.RUnlock()

// 写
mu.Lock()
m["key"] = "value"
mu.Unlock()
```

**选择建议**:
- 读多写少：`sync.RWMutex + map`
- 写多或不确定：`sync.Map`

---

## 7. 线程池 vs Goroutine 池

### Java: ExecutorService
```java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

ExecutorService pool = Executors.newFixedThreadPool(10);

for (int i = 0; i &lt; 100; i++) {
    pool.submit(() -&gt; {
        // 任务
    });
}

pool.shutdown();
```

### Go: 通常不需要池！Goroutine 足够轻量
```go
// 直接启动 Goroutine，通常不需要池
for i := 0; i &lt; 100000; i++ {
    go func() {
        // 任务
    }()
}

// 需要限制并发时，用 Channel 做信号量
sem := make(chan struct{}, 100)
for i := 0; i &lt; 100000; i++ {
    sem &lt;- struct{}{}
    go func() {
        defer func() { &lt;-sem }()
        // 任务
    }()
}
```

**Goroutine 池实现（可选）**:
```go
type Pool struct {
    tasks chan func()
    wg    sync.WaitGroup
}

func NewPool(size int) *Pool {
    p := &amp;Pool{tasks: make(chan func())}
    p.wg.Add(size)
    for i := 0; i &lt; size; i++ {
        go func() {
            defer p.wg.Done()
            for task := range p.tasks {
                task()
            }
        }()
    }
    return p
}
```

---

## 8. Future/Promise vs Channel

### Java: CompletableFuture
```java
import java.util.concurrent.CompletableFuture;

CompletableFuture&lt;String&gt; future = CompletableFuture.supplyAsync(() -&gt; {
    return doSomething();
});

future.thenAccept(result -&gt; {
    System.out.println(result);
});

String result = future.get();  // 阻塞等待
```

### Go: Channel
```go
ch := make(chan string, 1)

go func() {
    ch &lt;- doSomething()
}()

// 阻塞等待
result := &lt;-ch

// 或带超时
select {
case result := &lt;-ch:
    fmt.Println(result)
case &lt;-time.After(5 * time.Second):
    fmt.Println("timeout")
}
```

---

## 9. 原子操作

### Java: AtomicInteger
```java
import java.util.concurrent.atomic.AtomicInteger;

AtomicInteger counter = new AtomicInteger(0);

counter.incrementAndGet();
int val = counter.get();
```

### Go: sync/atomic
```go
import "sync/atomic"

var counter int64

atomic.AddInt64(&amp;counter, 1)
val := atomic.LoadInt64(&amp;counter)
```

---

## 10. 总结对比表

| Java (java.util.concurrent) | Go (sync/channel) | 说明 |
|---------------------------|-------------------|------|
| `Thread` | `Goroutine` | Goroutine 更轻量 |
| `ReentrantLock` | `sync.Mutex` | Go 用 defer 确保解锁 |
| `ReentrantReadWriteLock` | `sync.RWMutex` | 功能类似 |
| `CountDownLatch` | `sync.WaitGroup` | WaitGroup 更易用 |
| 双重检查锁 | `sync.Once` | Once 更安全简洁 |
| `Condition` | `sync.Cond` / Channel | 优先用 Channel |
| `ConcurrentHashMap` | `sync.Map` / RWMutex+map | 根据场景选择 |
| `ExecutorService` | 通常不需要 / 工作池 | Goroutine 足够轻量 |
| `CompletableFuture` | Channel | Channel 更灵活 |
| `AtomicInteger` | `sync/atomic` | 功能类似 |

**Go 并发编程的核心**:
- Goroutine: 轻量级执行单元
- Channel: 用于同步和通信
- sync 包: 基本同步原语（锁、条件变量等）
- 优先使用 Channel，其次使用 sync 包
