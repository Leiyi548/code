package main

import (
	"fmt"
)

func isValid(s string) bool {
	n := len(s)
	// 如果长度不为偶数，肯定是不是有效的括号
	if n%2 != 0 {
		return false
	}
	pairs := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		// 如果这个字符是左括号，将它放入栈
		if pairs[s[i]] > 0 {
			stack = append(stack, s[i])
		} else {
			// 去除括号
			// 如果长度为0，那么这个时候，有右括号就没有匹配的直接返回false
			// 匹配失败，直接返回false
			if len(stack) == 0 || s[i] != pairs[stack[len(stack)-1]] {
				return false
			}
			// 如果匹配成功就把左括号从这个栈中去除
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
