# 2. 接口nil判断坑点

这是Java转Go开发者踩坑率最高的Top1问题，无数人在生产环境中因为这个问题导致错误被吞、服务异常。

## 底层原理
一个接口变量等于`nil`，当且仅当它的**类型指针**和**数据指针**都为`nil`：
```go
// 空接口底层结构
type eface struct {
    _type *type // 类型指针
    data  unsafe.Pointer // 数据指针
}
```

## 坑点1：返回nil的具体类型指针给接口，导致错误判断失效
```go
// ❌ 错误写法：返回nil的*CustomError给error接口
type CustomError struct {
    Code int
    Msg  string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

func GetUser(id int64) error {
    var err *CustomError
    if id <= 0 {
        err = &CustomError{Code: 400, Msg: "invalid id"}
    }
    // 当id>0时，err的类型是*CustomError（不是nil），数据是nil
    // 所以返回的error接口不是nil！
    return err
}

func main() {
    err := GetUser(1)
    if err != nil { // 这里永远为true！即使没有错误
        log.Fatal("unexpected error: ", err) // 会输出：unexpected error: <nil>
    }
}
```

✅ 正确写法：永远直接返回`nil`，不要返回nil的具体类型指针
```go
func GetUser(id int64) error {
    if id <= 0 {
        return &CustomError{Code: 400, Msg: "invalid id"}
    }
    // 直接返回nil，类型和数据都是nil
    return nil
}
```

## 坑点2：把nil的接口赋值给具体类型，导致panic
```go
var i any = nil
var s string = i.(string) // panic: interface conversion: interface {} is nil, not string
```

✅ 正确写法：使用comma-ok断言
```go
s, ok := i.(string)
if !ok {
    // 处理类型不匹配的情况
}
```

## 坑点3：interface{}和具体类型的`==`判断问题
```go
var a int = 0
var b int64 = 0
var ia any = a
var ib any = b

fmt.Println(ia == ib) // false！因为类型不同
```

## 练习题
1. 修复`GetUser`函数的错误，让它能正确返回nil error
2. 编写代码验证：只有当接口的类型和数据都为nil时，接口才等于nil
3. 解释为什么`fmt.Println(err)`会输出`<nil>`，但`err != nil`却为true