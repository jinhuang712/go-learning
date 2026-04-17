# Section 3: Context 传递最佳实践

## 1. 知识点核心说明

Context 的传递方式对代码的可维护性和性能都有重要影响。本节介绍在微服务开发中 Context 传递的最佳实践。

---

## 2. Java 与 Go 的对比说明

| 实践 | Java | Go |
|------|------|----|
| **Context 参数位置** | 灵活，通常靠后 | 强制作为第一个参数 |
| **存储方式** | ThreadLocal (隐式) | 显式传递 |
| **中间件传递** | Servlet Filter + ThreadLocal | HTTP Handler + Context 参数 |
| **跨 Goroutine 传递** | 困难 (线程池 + ThreadLocal) | 容易 (显式传递) |

### HTTP 请求处理对比

**Java (Spring Boot + ThreadLocal)**:
```java
@Component
public class UserInterceptor implements HandlerInterceptor {
    private static final ThreadLocal&lt;String&gt; USER_ID = new ThreadLocal&lt;&gt;();
    
    @Override
    public boolean preHandle(HttpServletRequest request, 
                            HttpServletResponse response, 
                            Object handler) {
        String userId = request.getHeader("X-User-ID");
        USER_ID.set(userId);
        return true;
    }
    
    @Override
    public void afterCompletion(HttpServletRequest request, 
                               HttpServletResponse response, 
                               Object handler, Exception ex) {
        USER_ID.remove(); // 重要：清理 ThreadLocal！
    }
}

@RestController
public class UserController {
    @GetMapping("/user")
    public String getUser() {
        String userId = UserInterceptor.USER_ID.get();
        return "User: " + userId;
    }
}
```

**Go (net/http + Context)**:
```go
package main

import (
    "context"
    "net/http"
)

type userIDKey struct{}

func userMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := r.Header.Get("X-User-ID")
        ctx := context.WithValue(r.Context(), userIDKey{}, userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value(userIDKey{}).(string)
    w.Write([]byte("User: " + userID))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/user", userHandler)
    
    http.ListenAndServe(":8080", userMiddleware(mux))
}
```

**心智迁移要点**:
- Java: 依靠 ThreadLocal + Filter，容易忘记清理导致内存泄露
- Go: 显式传递 Context，更清晰，不会泄露
- Java: 在线程池中传递 ThreadLocal 很麻烦
- Go: 跨 Goroutine 传递 Context 很简单（显式传递即可）
- Java: Servlet 规范没有标准的 Context 传递方式
- Go: 标准库（net/http, database/sql, gRPC 等）都支持 Context

---

## 3. 可运行代码示例

### 最佳实践 1: Context 总是作为第一个参数

```go
package main

import (
    "context"
    "fmt"
)

// 正确写法：Context 作为第一个参数
func processRequest(ctx context.Context, reqID string) {
    fmt.Println("处理请求:", reqID)
    // ... 处理逻辑
}

// 错误写法：Context 不是第一个参数
func badProcessRequest(reqID string, ctx context.Context) {
    // 不要这样写！
}

func main() {
    ctx := context.Background()
    processRequest(ctx, "req-123")
}
```

### 最佳实践 2: 不要把 Context 存在结构体里

```go
package main

import (
    "context"
    "fmt"
)

// 错误写法：把 Context 存在结构体里
type BadWorker struct {
    ctx context.Context // 不要这样做！
}

func (w *BadWorker) DoWork() {
    // 使用 w.ctx
}

// 正确写法：每次调用时传递 Context
type GoodWorker struct{}

func (w *GoodWorker) DoWork(ctx context.Context) {
    // 使用传入的 ctx
    fmt.Println("工作中...")
}

func main() {
    ctx := context.Background()
    
    worker := &amp;GoodWorker{}
    worker.DoWork(ctx)
}
```

### 最佳实践 3: WithValue 的 key 使用自定义类型

```go
package main

import (
    "context"
    "fmt"
)

// 正确写法：定义包私有的自定义类型作为 key
type (
    userIDKey struct{}
    traceIDKey struct{}
)

func main() {
    ctx := context.Background()
    
    // 设置值
    ctx = context.WithValue(ctx, userIDKey{}, "user-123")
    ctx = context.WithValue(ctx, traceIDKey{}, "trace-456")
    
    // 获取值
    userID := ctx.Value(userIDKey{})
    traceID := ctx.Value(traceIDKey{})
    
    fmt.Printf("User: %v, Trace: %v\n", userID, traceID)
}
```

### 最佳实践 4: 派生 Context 后立即 defer cancel()

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func doWork(ctx context.Context) {
    select {
    case &lt;-time.After(1 * time.Second):
        fmt.Println("工作完成")
    case &lt;-ctx.Done():
        fmt.Println("工作取消:", ctx.Err())
    }
}

func main() {
    ctx := context.Background()
    
    // 正确写法：立即 defer cancel()
    ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
    defer cancel() // 确保即使提前返回也能调用 cancel()
    
    doWork(ctx)
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 把 Context 存在结构体里**
```go
// 错误写法！
type Service struct {
    ctx context.Context  // 不要这样做！
    db  *Database
}

func (s *Service) DoSomething() {
    // 使用 s.ctx
    // 问题：一个 Service 实例可能被多个请求共享！
}

// 正确写法
type Service struct {
    db *Database
}

func (s *Service) DoSomething(ctx context.Context) {
    // 每次调用都传入独立的 ctx
}
```

⚠️ **坑点 2: 忽略 Context 直接使用 Background()**
```go
// 错误写法：忽略传入的 ctx，直接用 Background()
func handler(ctx context.Context) {
    // 不要这样做！会丢失取消信号、超时等
    doWork(context.Background())
}

// 正确写法：使用传入的 ctx 派生
func handler(ctx context.Context) {
    // 正确：从传入的 ctx 派生
    childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    doWork(childCtx)
}
```

⚠️ **坑点 3: WithValue 存储可变对象**
```go
// 错误写法：存储可变对象，可能被其他 Goroutine 修改
type Data struct {
    Value string
}
data := &amp;Data{Value: "old"}
ctx = context.WithValue(ctx, "data", data)

// 另一个 Goroutine 修改了 data
data.Value = "new" // 会影响 ctx 中的值！

// 正确写法：只存储不可变值或深拷贝
ctx = context.WithValue(ctx, "data", "old") // 不可变
// 或者深拷贝后再存储
```

---

## 5. 练习题

### 练习 3.1: 实现一个 HTTP 中间件
**目标**: 写一个 HTTP 中间件，从请求头中读取 `X-User-ID` 和 `X-Trace-ID`，将它们存入 Context，然后传递给下游 Handler。

**验收标准**:
- 使用自定义类型作为 WithValue 的 key
- 正确使用 `r.WithContext()`
- 下游 Handler 能正确从 Context 中取出值
- 代码结构清晰，符合 Go 风格

### 练习 3.2: 理解 Context 传递（概念题）
**问题**:
1. 为什么 Go 约定 Context 总是作为函数的第一个参数？
2. 为什么不要把 Context 存在结构体里？请举一个会出问题的场景。

**验收标准**: 能够清晰解释这两个问题。

---

Context 是 Go 微服务开发的核心，遵循最佳实践能让代码更清晰、更易维护、更少 bug！
