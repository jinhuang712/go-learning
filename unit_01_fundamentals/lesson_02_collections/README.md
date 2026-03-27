# Lesson 2: 集合类型 (Collections)

本课程主要介绍 Go 语言中的核心集合类型：数组、切片 (Slice) 和 映射 (Map)。

## 核心知识点

### 1. 数组 (Array) vs 切片 (Slice)
| 特性 | 数组 (Array) | 切片 (Slice) | Java 对应 |
| :--- | :--- | :--- | :--- |
| **长度** | **固定** (`[3]int`) | **动态** (`[]int`) | Array vs ArrayList |
| **传递** | 值传递 (Copy) | 引用传递 (Reference) | - |

**重点**: 切片底层是对数组的引用。修改切片元素可能会影响底层数组和其他共享该数组的切片。

### 2. Map (映射)
- **定义**: `map[KeyType]ValueType`
- **操作**:
  - 创建: `make(map[string]int)`
  - 查找: `val, ok := m[key]` ("comma ok" idiom)
  - 遍历: `for k, v := range m` (注意：顺序随机)

## 练习
1. 创建一个 Slice，观察 `append` 导致的 `cap` (容量) 变化。
2. 尝试从 Map 中删除一个 Key: `delete(m, key)`。
