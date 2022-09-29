package main

// golang 分配内存主要有内置函数new和make。

// make只能为slice, map, channel分配内存，并返回一个初始化的值。首先来看下make有以下三种不同的用

// 第一种用法，即缺少长度的参数，只传类型，这种用法只能用在类型为map或chan的场景，例如make([]int)是会报错的。这样返回的空间长度都是默认为0的。


func main() {

}
