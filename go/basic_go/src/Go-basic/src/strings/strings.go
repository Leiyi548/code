package main

import (
	"fmt"
	"strings"
)

// 字符串去除空格
func removeSpace() {
	s1 := "   hello,world"
	s2 := "hello,world   "
	trimStr1 := strings.TrimSpace(s1)
	trimStr2 := strings.TrimSpace(s2)
	fmt.Println(trimStr1)
	fmt.Println(trimStr2)
}

func main() {
	remveSpace()
}
