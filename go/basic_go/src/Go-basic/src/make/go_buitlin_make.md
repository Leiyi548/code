# go_buitlin_make

`make()` 是 `Go` 语言内存分配的内置函数，默认有三个参数。

`make(Type, len, cap)`

- `Type`：数据类型，必要参数，`Type` 的值只能是 `slice`、 `map`、 `channel` 这三种数据类型。
- `len`：数据类型实际占用的内存空间长度，`map`、 `channel` 是可选参数，`slice` 是必要参数。
- `cap`：为数据类型提前预留的内存空间长度，可选参数。所谓的提前预留是当前为数据类型申请内存空间的时候，提前申请好额外的内存空间，这样可以避免二次分配内存带来的开销，大大提高程序的性能。

为了能更好的理解这些参数的含义，我们先来看下 make() 的三种不同用法：

```go
make(map[string]string)

make([]int, 2)

make([]int, 2, 4)
```

- 第一种，只传类型，不指定实际占用的内存空间和提前预留的内存空间，适用于 `map` 和 `channel` 。
- 第二种，指定实际占用的内存空间为 2，不指定提前预留的内存空间。
- 第三种，指定实际占用的内存空间为 2，指定提前预留的内存空间是 4。

看到这里你有没有这样的疑惑，既然在初始化的时候已经指定数据的大小了，那为什么还要指定预留的大小呢？这是因为 `make()` 使用的是一种动态数组算法，一开始先向操作系统申请一小块内存，这个就是 `cap`，等 `cap` 被 `len` 占用满以后就需要扩容，扩容就是动态数组再去向操作系统申请当前长度的两倍的内存，然后将旧数据复制到新内存空间中。

为了更好的理解动态数组向扩容，这里先写一段演示代码：

```go
data := make([]int, 0)

for i, n := 0, 20; i < n; i++ {
		data = append(data, 1)
		fmt.Printf("len=%d cap=%d\n", len(data), cap(data))
}
```

运行结果：

```text
len=1 cap=1		# 第一次扩容
len=2 cap=2		# 第二次扩容
len=3 cap=4		# 第三次扩容
len=4 cap=4
len=5 cap=8		# 第四次扩容
len=6 cap=8
len=7 cap=8
len=8 cap=8
len=9 cap=16	# 第五次扩容
len=10 cap=16
len=11 cap=16
len=12 cap=16
len=13 cap=16
len=14 cap=16
len=15 cap=16
len=16 cap=16
len=17 cap=32	# 第六次扩容
len=18 cap=32
len=19 cap=32
len=20 cap=32
```

结果确实是动态数组每次扩容都是原来的两倍，那么思考一下动态数组会一直这样扩容下去吗？

```go
data := make([]int, 0)

for i, n := 0, 10000; i < n; i++ {
		data = append(data, 1)
		if len(data) == cap(data) {
				fmt.Printf("len=%d cap=%d\n", len(data), cap(data))
		}
}
```

运行结果：

```text
len=1 cap=1
len=2 cap=2
len=4 cap=4
len=8 cap=8
len=16 cap=16
len=32 cap=32
len=64 cap=64
len=128 cap=128
len=256 cap=256
len=512 cap=512
len=1024 cap=1024
len=1280 cap=1280
len=1696 cap=1696
len=2304 cap=2304
len=3072 cap=3072
len=4096 cap=4096
len=5120 cap=5120
len=7168 cap=7168
len=9216 cap=9216
```

理论上扩容应该是呈指数增长，但我们从运行结果中可以看到，前 11 次扩容确实是每次扩展一倍的长度，不过在第 12 次扩容并没有按照预期扩展到 2048，这是为什么呢？

要回答这个问题就需要回到动态数组算法上。动态数组刚刚也说了，就是每次内存容量占满就需要扩容。但不同的程序员对扩容多少的理解是不一样的，不同的语言也有不同的算法，不过最后肯定在扩容 n 次后就不能按照倍数来扩了，这是因为有物理内存的限制，避免一次申请过多从而导致内存申请失败和内存浪费的情况
