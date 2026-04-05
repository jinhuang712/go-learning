# 3. 错误处理坑点

Java的异常体系是控制流跳跃，而Go的错误是普通值，思维转换不当很容易写出问题代码。

## 坑点1：模仿Java异常，用panic代替错误返回
```go
// ❌ 错误写法：业务错误用panic抛出，性能差且容易导致进程崩溃
func GetUser(id int64) *User {
    user, err := db.Query(id)
    if err != nil {
        // 模仿Java抛异常，直接panic
        panic(err)
    }
    return user
}
```

✅ 正确做法：业务错误必须用error返回，只有不可恢复的致命错误才用panic
```go
func GetUser(id int64) (*User, error) {
    user, err := db.Query(id)
    if err != nil {
        // 给错误添加上下文后返回
        return nil, fmt.Errorf("query user failed: %w", err)
    }
    return user, nil
}
```

## 坑点2：忽略错误，导致隐形问题
Java中不catch异常会向上抛，而Go中忽略错误会导致问题被隐藏：
```go
// ❌ 错误写法：忽略error返回值
user, _ := GetUser(1) // 万一GetUser返回错误，user是nil，后续操作会panic
fmt.Println(user.Name)
```

✅ 正确做法：必须显式处理所有error，实在不需要处理也要用`_`明确表示忽略，并且加注释说明原因
```go
// 忽略错误是因为这里即使失败也不影响主流程
user, _ := GetUser(1) 
if user != nil {
    fmt.Println(user.Name)
}
```

## 坑点3：错误没有上下文，排查困难
```go
// ❌ 错误写法：直接返回原始错误，不知道是哪里发生的
func CreateOrder(userID int64) error {
    user, err := GetUser(userID)
    if err != nil {
        return err // 调用方收到错误不知道是获取用户失败导致的
    }
    // ...
}
```

✅ 正确做法：用`fmt.Errorf("%w", err)`包装错误，添加上下文信息
```go
func CreateOrder(userID int64) error {
    user, err := GetUser(userID)
    if err != nil {
        return fmt.Errorf("create order failed: get user %d error: %w", userID, err)
    }
    // ...
}
```

## 坑点4：滥用errors.Is/As，性能下降
```go
// ❌ 错误写法：每次都用As判断错误类型，不要在循环中大量使用
if errors.As(err, &customErr) {
    // 处理
}
```

## 练习题
1. 修复上面的`GetUser`函数的panic问题，改为返回error
2. 给一个三层调用的函数（controller->service->dao），每一层都给错误添加上下文
3. 编写代码演示如何用`errors.Is`和`errors.As`判断被包装的错误类型