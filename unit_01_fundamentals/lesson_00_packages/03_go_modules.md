# Section 3: Go Module 系统

## 1. 知识点核心说明

### 什么是 Go Module
Go Module 是 Go 1.11+ 引入的官方依赖管理系统，用于管理 Go 项目的依赖。

### `go.mod` 文件
每个 Go Module 根目录下都有一个 `go.mod` 文件：

```go
module example.com/myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.0
)
```

### Module 路径
Module 路径是模块的唯一标识，通常是代码仓库的地址：
- `github.com/username/project`
- `example.com/myproject`

### 常用命令

| 命令 | 说明 |
|------|------|
| `go mod init &lt;module-path&gt;` | 初始化新 module |
| `go mod tidy` | 添加缺失的依赖，删除未使用的依赖 |
| `go get &lt;package&gt;` | 添加或更新依赖 |
| `go mod download` | 下载依赖到本地缓存 |
| `go list -m all` | 列出所有依赖 |

---

## 2. Java 与 Go 的对比说明

| 特性 | Java Maven/Gradle | Go Modules |
|------|------------------|------------|
| **配置文件** | `pom.xml` / `build.gradle` | `go.mod` |
| **依赖锁文件** | `pom.xml` (Maven) / `gradle.lockfile` | `go.sum` |
| **依赖标识** | `groupId:artifactId:version` | `module-path@version` |
| **本地仓库** | `~/.m2/repository` | `GOPATH/pkg/mod` |
| **传递依赖** | 自动处理 | 自动处理 |
| **依赖范围** | compile, test, provided | 无，通过 build tag 或 Go 版本区分 |

### 代码对照

**Maven (pom.xml)**:
```xml
&lt;project&gt;
    &lt;groupId&gt;com.example&lt;/groupId&gt;
    &lt;artifactId&gt;myproject&lt;/artifactId&gt;
    &lt;version&gt;1.0.0&lt;/version&gt;
    
    &lt;dependencies&gt;
        &lt;dependency&gt;
            &lt;groupId&gt;org.springframework.boot&lt;/groupId&gt;
            &lt;artifactId&gt;spring-boot-starter-web&lt;/artifactId&gt;
            &lt;version&gt;3.0.0&lt;/version&gt;
        &lt;/dependency&gt;
    &lt;/dependencies&gt;
&lt;/project&gt;
```

**Go (go.mod)**:
```go
module example.com/myproject

go 1.21

require github.com/gin-gonic/gin v1.9.0
```

---

## 3. 可运行代码示例

### 示例 1: 初始化和使用 Go Module

```bash
# 初始化新 module
go mod init example.com/myproject

# 添加依赖
go get github.com/gin-gonic/gin

# 整理依赖
go mod tidy

# 查看所有依赖
go list -m all
```

```go
// main.go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "hello world",
        })
    })
    r.Run()
}
```

---

## 4. 生产环境坑点提示

⚠️ **坑点 1: 提交 vendor 目录到 Git**
```bash
# 可选 - 将依赖提交到仓库
go mod vendor

# 是否提交 vendor/ 到 Git？
# - 优点：构建时不需要网络
# - 缺点：仓库体积变大
# - 建议：CI/CD 环境不提交，离线构建场景考虑提交
```

⚠️ **坑点 2: 直接修改 go.mod**
```bash
# 不推荐 - 手动编辑 go.mod 容易出错

# 推荐 - 使用 go 命令管理
go get github.com/example/pkg@v1.2.3
go mod tidy
```

⚠️ **坑点 3: 忽略 go.sum**
```
# .gitignore 中不要忽略 go.sum
go.sum  # 不要忽略！它是依赖的校验和，保证依赖不被篡改
```

---

## 5. 练习题

### 练习 3.1: 创建一个使用 Go Module 的项目
**目标**: 创建一个新的 Go 项目，初始化 Go Module，添加一个第三方依赖（如 `github.com/google/uuid`），并在代码中使用它。

**验收标准**:
- 正确初始化 go.mod
- 成功添加依赖
- 代码能正常编译和运行
- go.mod 和 go.sum 文件正确

### 练习 3.2: 理解 Go Module（概念题）
**问题**:
1. `go.mod` 和 `go.sum` 分别有什么作用？
2. Go Module 和 Java Maven 在依赖管理上有什么主要区别？

**验收标准**: 能够清晰解释这两个问题。
