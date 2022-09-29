package main

import (
	"fmt"
	"reflect"
)

func test01() {
	// 区间
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:3]) // [1 2 3] [2 3]
	// 中间到尾部的所有元素
	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b, b[1:]) // [1 2 3 4 5] [2 3 4 5]
	// 开头到中间指定位置的所有元素
	var c = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} //  [1 2 3 4 5 6 7 8 9] [1 2]
	fmt.Println(c, c[:2])
}

func test02() {
	s := "cdbbd"
	fmt.Println(s[1:5])
}

func main() {
	test01()
	s := "abc"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(s[0]))
	test02()
}
