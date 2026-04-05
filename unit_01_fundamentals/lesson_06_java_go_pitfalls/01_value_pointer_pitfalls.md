# 1. 值/指针误用坑点

## 坑点1：方法接收者混用值和指针，导致状态修改失效
```go
// ❌ 错误写法：值接收者修改状态无效
type User struct {
    Name string
    Age  int
}

func (u User) SetAge(age int) {
    u.Age = age // 修改的是副本，原对象不变
}

func main() {
    u := User{Name: "Bob", Age: 20}
    u.SetAge(30)
    fmt.Println(u.Age) // 输出 20，不是30！
}
```

✅ 正确写法：需要修改状态的方法必须用指针接收者
```go
func (u *User) SetAge(age int) {
    u.Age = age // 修改原对象
}
```

## 坑点2：大结构体传值导致严重性能开销
Java中对象默认传引用，没有这个问题，但Go中传值会拷贝整个结构体：
```go
// ❌ 错误写法：大结构体传值，每次调用都拷贝几十KB内存
type BigStruct struct {
    Field1 [1024]int
    Field2 [1024]string
    // 还有几十个其他字段...
}

func Process(s BigStruct) {
    // 处理逻辑
}
```

✅ 正确写法：大结构体传指针，只拷贝8字节地址
```go
func Process(s *BigStruct) {
    // 处理逻辑
}
```

## 坑点3：值类型不能为nil，误用导致判断错误
Java中所有对象都可以为null，但Go中值类型（int, string, struct等）默认有零值，不能为nil：
```go
// ❌ 错误写法：无法区分"没传值"和"传了零值"
type UserRequest struct {
    Age int // 前端没传的话默认是0，无法区分是真的0还是没传
}
```

✅ 正确写法：需要区分零值和未传的场景用指针
```go
type UserRequest struct {
    Age *int // 没传的话是nil，传了0的话是*int=0
}
```