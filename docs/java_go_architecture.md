# Java-Go 架构思路对比

## 前言
本文档从架构设计思路的层面对比 Java 和 Go 的差异，帮助 Java 开发者理解 Go 的设计哲学，避免用 Java 思维写 Go 代码。

---

## 1. 面向对象 vs 组合

### Java: 基于继承的 OOP
```java
// Java: 继承
public class Animal {
    protected String name;
    public void eat() { /* ... */ }
}

public class Dog extends Animal {
    public void bark() { /* ... */ }
}
```

**Java 特点**:
- 类继承
- 多态通过继承实现
- 设计模式大量使用继承

### Go: 基于组合的设计
```go
// Go: 组合而非继承
type Animal struct {
    Name string
}

func (a *Animal) Eat() { /* ... */ }

type Dog struct {
    Animal  // 匿名嵌入，类似继承但只是组合
    Breed string
}

func (d *Dog) Bark() { /* ... */ }
```

**Go 特点**:
- 没有继承，只有组合
- 通过接口实现多态
- 匿名嵌入只是语法糖，不是父子关系

**心智迁移**:
- 不要尝试用 struct 模拟类继承
- 多用小接口，少用大接口
- 使用组合来复用代码

---

## 2. 异常处理 vs 值错误

### Java: 异常机制
```java
// Java: 异常
public User getUser(Long id) throws UserNotFoundException {
    User user = repository.findById(id);
    if (user == null) {
        throw new UserNotFoundException("User not found");
    }
    return user;
}

try {
    User user = getUser(123);
} catch (UserNotFoundException e) {
    // 处理异常
}
```

**Java 特点**:
- 异常分为 Checked 和 Unchecked
- 异常会中断控制流
- try-catch-finally 处理

### Go: 值错误
```go
// Go: 多返回值 + error
func GetUser(id int64) (*User, error) {
    user := repository.FindByID(id)
    if user == nil {
        return nil, fmt.Errorf("user not found: %d", id)
    }
    return user, nil
}

user, err := GetUser(123)
if err != nil {
    // 处理错误
    return err
}
// 继续执行
```

**Go 特点**:
- 错误是普通值
- 多返回值返回错误
- 显式处理每个错误

**心智迁移**:
- 不要用 panic 代替 error
- 错误是控制流的一部分，不是异常
- 错误处理要显式，不要忽略 error

---

## 3. IoC 容器 vs 手动依赖注入

### Java: Spring IoC 容器
```java
// Java: Spring 自动注入
@Service
public class UserService {
    
    @Autowired
    private UserRepository repository;
    
    @Autowired
    private EmailService emailService;
}
```

**Java 特点**:
- 容器管理对象生命周期
- 自动注入依赖
- 配置驱动

### Go: 手动依赖注入
```go
// Go: 手动注入
type UserService struct {
    repo  UserRepository
    email EmailService
}

func NewUserService(repo UserRepository, email EmailService) *UserService {
    return &amp;UserService{
        repo:  repo,
        email: email,
    }
}

// main 中组装
repo := NewUserRepository()
email := NewEmailService()
userService := NewUserService(repo, email)
```

**Go 特点**:
- 手动组装依赖
- 通过构造函数传递
- 显式、清晰

**心智迁移**:
- 不要寻找 Spring 那样的 IoC 容器
- 在 main 函数中组装所有依赖
- 构造函数应该是纯的，只做参数验证和赋值

---

## 4. 线程池 vs Goroutine

### Java: 线程池
```java
// Java: 线程池管理
ExecutorService pool = Executors.newFixedThreadPool(10);

for (int i = 0; i &lt; 100; i++) {
    pool.submit(() -&gt; {
        // 任务
    });
}

pool.shutdown();
```

**Java 特点**:
- 线程开销大，需要池化
- 线程池大小需要谨慎配置
- 上下文切换昂贵

### Go: Goroutine
```go
// Go: 直接用 Goroutine
for i := 0; i &lt; 100000; i++ {
    go func() {
        // 任务
    }()
}

// 可选：用 Channel 或工作池限制并发
sem := make(chan struct{}, 100)
for i := 0; i &lt; 100000; i++ {
    sem &lt;- struct{}{}
    go func() {
        defer func() { &lt;-sem }()
        // 任务
    }()
}
```

**Go 特点**:
- Goroutine 开销极小，通常不需要池化
- GOMAXPROCS 控制并行度
- Channel 用于同步和控制

**心智迁移**:
- 不要一上来就建 Goroutine 池
- 先考虑直接用 Goroutine
- 需要限制并发时用 Channel 做信号量

---

## 5. Null 安全

### Java: NullPointerException
```java
// Java: NPE 风险
User user = getUser();
String name = user.getName();  // 如果 user 是 null，抛 NPE

// Java 8+: Optional
Optional&lt;User&gt; userOpt = getUser();
userOpt.ifPresent(u -&gt; { /* ... */ });
```

### Go: nil 检查
```go
// Go: 显式 nil 检查
user, err := GetUser()
if err != nil {
    return err
}
if user == nil {
    return errors.New("user not found")
}
name := user.Name  // 安全

// 接口 nil 的坑！
var u *User = nil
var i interface{} = u
fmt.Println(i == nil)  // false! 注意这个坑！
```

**心智迁移**:
- 总是检查 error
- 显式检查 nil
- 注意接口 nil 判断的坑

---

## 6. 总结：Go 的设计哲学

| 原则 | 说明 |
|------|------|
| **简单优于复杂** | 避免过度设计，保持代码简单直接 |
| **组合优于继承** | 使用组合复用代码，而非继承 |
| **明确优于隐式** | 显式处理错误，显式传递依赖 |
| **并发是一等公民** | Goroutine + Channel，CSP 模型 |
| **标准库优先** | 优先使用标准库，避免不必要的依赖 |

**Java 开发者的反模式（不要这样写 Go）**:

❌ 错误：用 struct 模拟类继承
```go
// 不要这样！
type Base struct{}
type Derived struct { Base }
```

❌ 错误：忽略 error
```go
// 不要这样！
user, _ := GetUser()
```

❌ 错误：用 panic 代替 error
```go
// 不要这样！
if err != nil {
    panic(err)
}
```

❌ 错误：过度设计
```go
// 不要这样！
// 不要写复杂的工厂、抽象层
// 保持简单！
```
