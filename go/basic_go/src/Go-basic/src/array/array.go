package main

import (
	"fmt"
	"sort"
)

func test01() {
	// 初始化一个数组
	var a [3]int  // 定义三个整数的数组 默认为0
	var b [3]bool // 定义三个布尔值的数组 默认为false
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素
	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

func test02() {
	// initialized a empty array
	var binDimensionArray = [][]int{}
	fmt.Printf("binDimensionArray: %v\n", binDimensionArray)
	binDimensionArray = append(binDimensionArray, []int{1, 2, 3})
	fmt.Printf("binDimensionArray: %v\n", binDimensionArray)
}

type intList []int

func (array intList) Len() int {
	return len(array)
}

func (array intList) Less(i, j int) bool {
	return array[i] < array[j]
}

func (array intList) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

func test03() {
	// arr := [...]int{9, 3, 7, 2, 6}
	// sort.Sort(arr)
	arr := intList{
		8,
		1,
		2,
		9,
		3,
		8,
	}
	sort.Sort(arr)
	fmt.Printf("arr: %v\n", arr)
}

func main() {
	// test02()
	test03()
}
