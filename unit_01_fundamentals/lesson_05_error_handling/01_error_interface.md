# 1. 错误接口与显式错误处理 (Error Interface)

在 Java 中，我们习惯于使用 `try-catch` 捕获 `Exception`。这是一种**控制流的跳跃**。
而在 Go 语言中，错误（Error）被视为一种**普通的值（Value）**。它不是异常，你需要像处理其他返回值一样去显式地处理它。

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
