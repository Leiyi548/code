# go_builtin_map

Go 语言中 map 是一种特殊的数据结构，一种元素对（`pair`）的无序集合，`pair` 对应一个 `key`（索引）和一个 `value`（值），所以这个结构也称为关联数组或字典，这是一种能够快速寻找值的理想结构，给定 `key`，就可以迅速找到对应的 `value`。

`map` 这种数据结构在其他编程语言中也称为字典（`Python`）、`hash` 和 `HashTable` 等。

`map` 是引用类型，可以使用如下方式声明：

```go
 var mapname map[keytype]valuetype
```

- mapname 为 map 的变量名。
- keytype 为键类型。
- valuetype 是键对应的值类型

> 提示：[keytype] 和 valuetype 之间允许有空格。

**在声明的时候不需要知道 map 的长度，因为 map 是可以动态增长的，未初始化的 map 的值是 nil，使用函数 len() 可以获取 map 中 pair 的数目。**
