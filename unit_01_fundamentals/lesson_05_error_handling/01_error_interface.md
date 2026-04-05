# 1. 错误接口与显式错误处理 (Error Interface)

## 核心设计哲学对比
| 维度 | Java 异常体系 | Go 错误处理 |
|------|--------------|-------------|
| 本质 | 控制流跳跃机制 | 普通返回值 |
| 处理方式 | 隐式，异常可以跨多层自动向上抛，直到被catch | 显式，必须每层显式处理或向上返回 |
| 性能开销 | 抛异常会生成栈快照，开销大 | 错误就是普通值，无额外开销 |
| 代码可读性 | try-catch嵌套导致主业务逻辑被淹没 | 错误处理和主逻辑并行，清晰可见 |
| 生产环境稳定性 | 未捕获的RuntimeException会直接导致进程崩溃 | 未处理的error只会作为返回值，不会导致进程崩溃 |

在 Java 中，我们习惯于使用 `try-catch` 捕获 `Exception`。这是一种**控制流的跳跃**。
而在 Go 语言中，错误（Error）被视为一种**普通的值（Value）**。它不是异常，你需要像处理其他返回值一样去显式地处理它。

### Java 转 Go 典型误区：
❌ 错误做法：模仿Java的异常，写一个统一的catch-all中间件把所有error捕获，返回统一HTTP响应
✅ 正确做法：每层显式处理错误，要么就地处理（重试、降级），要么给错误添加上下文信息后再向上返回：
```go
// 给错误添加上下文，方便排查问题
user, err := userService.Find(id)
if err != nil {
    // 用fmt.Errorf的%w包装原始错误，上层可以用errors.Is/As判断错误类型
    return nil, fmt.Errorf("get user failed: %w", err)
}
```

## 1.1 `error` 本质上只是一个接口
Go 底层对错误的定义极其简单，它只是一个带 `Error()` 方法的接口：
```go
// 源码定义
type error interface {
    Error() string
}
```

## 1.2 创建错误
- **基础错误**: 使用 `errors.New("错误信息")`
- **格式化错误**: 使用 `fmt.Errorf("找不到用户: %d", id)`

## 1.3 核心心智模型：多返回值与 `if err != nil`
在 Go 的微服务开发中，几乎所有可能失败的函数，**最后的一个返回值一定是 `error`**。
调用方**必须**显式检查 `err` 是否为 `nil`（表示成功）。

**Java 对应对比**:
- Java: `User user = userService.find(id);` (如果找不到抛出 `UserNotFoundException`)
- Go:
  ```go
  user, err := userService.Find(id)
  if err != nil {
      // 错误处理，比如记录日志、或者把错误继续向上层返回 (return nil, err)
      return nil, err 
  }
  // 走到这里说明 err == nil，业务继续
  ```

## 1.4 自定义错误类型 (Custom Errors)
既然 `error` 是个接口，我们就可以用结构体（Struct）来实现它，从而携带更多的上下文信息（比如错误码）。

```go
type NotFoundError struct {
    Resource string
    ID       int
}

// 只要实现了 Error() string，它就是个 error
func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s with ID %d not found", e.Resource, e.ID)
}
```
