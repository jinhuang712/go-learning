# Java-Go 设计模式对照

## 前言
本文档对比常用设计模式在 Java 和 Go 中的不同实现方式，帮助 Java 开发者理解 Go 语言的设计哲学。

Go 语言的设计哲学：
- 组合优于继承 (Composition over Inheritance)
- 简单优于复杂 (Simple is better than complex)
- 明确优于隐式 (Explicit is better than implicit)

---

## 1. Singleton (单例模式)

### Java 实现
```java
public class Singleton {
    private static volatile Singleton instance;
    
    private Singleton() {}
    
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

### Go 实现
```go
package singleton

import "sync"

type Singleton struct{}

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

**对比说明**:
- Java: 需要双重检查锁 + volatile
- Go: 使用 `sync.Once`，更简洁安全

---

## 2. Factory (工厂模式)

### Java 实现
```java
public interface Shape {
    void draw();
}

public class Circle implements Shape {
    public void draw() { System.out.println("Circle"); }
}

public class Square implements Shape {
    public void draw() { System.out.println("Square"); }
}

public class ShapeFactory {
    public static Shape getShape(String type) {
        if ("circle".equals(type)) return new Circle();
        if ("square".equals(type)) return new Square();
        return null;
    }
}
```

### Go 实现
```go
package factory

type Shape interface {
    Draw()
}

type Circle struct{}
func (c *Circle) Draw() { println("Circle") }

type Square struct{}
func (s *Square) Draw() { println("Square") }

func NewShape(shapeType string) Shape {
    switch shapeType {
    case "circle":
        return &amp;Circle{}
    case "square":
        return &amp;Square{}
    default:
        return nil
    }
}
```

**对比说明**:
- 结构类似，都是通过接口实现多态
- Go 的构造函数通常命名为 `NewXxx`

---

## 3. Strategy (策略模式)

### Java 实现
```java
public interface PaymentStrategy {
    void pay(int amount);
}

public class CreditCardPayment implements PaymentStrategy {
    public void pay(int amount) { /* ... */ }
}

public class ShoppingCart {
    private PaymentStrategy strategy;
    
    public void setStrategy(PaymentStrategy strategy) {
        this.strategy = strategy;
    }
}
```

### Go 实现
```go
package strategy

type PaymentStrategy interface {
    Pay(amount int)
}

type CreditCardPayment struct{}
func (c *CreditCardPayment) Pay(amount int) { /* ... */ }

type ShoppingCart struct {
    strategy PaymentStrategy
}

func (s *ShoppingCart) SetStrategy(strategy PaymentStrategy) {
    s.strategy = strategy
}
```

**Go 的另一种方式：函数作为策略**
```go
type PaymentFunc func(amount int)

func CreditCardPay(amount int) { /* ... */ }

type ShoppingCart struct {
    payFunc PaymentFunc
}

func (s *ShoppingCart) SetPaymentFunc(f PaymentFunc) {
    s.payFunc = f
}
```

**对比说明**:
- 经典实现类似
- Go 可以用函数作为策略，更灵活

---

## 4. Observer (观察者模式)

### Java 实现
```java
import java.util.ArrayList;
import java.util.List;

public interface Observer {
    void update(String event);
}

public class Subject {
    private List&lt;Observer&gt; observers = new ArrayList&lt;&gt;();
    
    public void register(Observer o) { observers.add(o); }
    public void notifyObservers(String event) {
        observers.forEach(o -&gt; o.update(event));
    }
}
```

### Go 实现
```go
package observer

type Observer interface {
    Update(event string)
}

type Subject struct {
    observers []Observer
}

func (s *Subject) Register(o Observer) {
    s.observers = append(s.observers, o)
}

func (s *Subject) NotifyObservers(event string) {
    for _, o := range s.observers {
        o.Update(event)
    }
}
```

**Go 的另一种方式：使用 Channel**
```go
type Event struct {
    Type string
    Data interface{}
}

type Subject struct {
    observers []chan&lt;- Event
}

func (s *Subject) Notify(event Event) {
    for _, ch := range s.observers {
        select {
        case ch &lt;- event:
        default: // 非阻塞，避免慢消费者
        }
    }
}
```

**对比说明**:
- 经典实现类似
- Go 可以用 Channel 实现异步非阻塞的观察者模式

---

## 5. Decorator (装饰器模式)

### Java 实现
```java
public interface Coffee {
    double cost();
}

public class SimpleCoffee implements Coffee {
    public double cost() { return 5.0; }
}

public class MilkDecorator implements Coffee {
    private Coffee coffee;
    public MilkDecorator(Coffee coffee) { this.coffee = coffee; }
    public double cost() { return coffee.cost() + 1.5; }
}
```

### Go 实现
```go
package decorator

type Coffee interface {
    Cost() float64
}

type SimpleCoffee struct{}
func (s *SimpleCoffee) Cost() float64 { return 5.0 }

type MilkDecorator struct {
    coffee Coffee
}
func (m *MilkDecorator) Cost() float64 { return m.coffee.Cost() + 1.5 }

func WithMilk(c Coffee) Coffee {
    return &amp;MilkDecorator{coffee: c}
}
```

**对比说明**:
- 结构类似
- Go 常用函数式的装饰器（如 HTTP Middleware）

---

## 总结

| 模式 | Java 特点 | Go 特点 |
|------|----------|---------|
| Singleton | 双重检查锁 + volatile | sync.Once 更简洁 |
| Factory | 继承 + 多态 | 组合 + 接口 |
| Strategy | 接口 + 实现 | 接口 + 实现 + 函数 |
| Observer | List + 遍历 | Slice + 遍历 + Channel |
| Decorator | 包装类 | 包装类 + 函数式 |

**核心心智迁移**:
- 抛弃继承思维，多用组合
- 善用函数作为一等公民
- 利用 Channel 实现并发模式
