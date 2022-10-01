package main

import (
	"fmt"
)

func myAtoi01(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	// flag为1则为正数
	// flag为-1则为负数
	flag := 1
	i := 0
	// 读入字符，删除无用的前导空格
	for _, ch := range s {
		if ch != ' ' {
			break
		} else {
			i++
		}
	}
	s = s[i:]
	if len(s) == 0 {
		return 0
	}
	// 检查下一个字符是正数还是负数
	if s[0:1] == "-" {
		s = s[1:]
		flag = -1
	} else if s[0:1] == "+" {
		s = s[1:]
		flag = 1
	} else {
		flag = 1
	}
	// 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
	for index, ch := range s {
		if ch == ' ' {
			s = s[:index]
			break
		}
		tmpNum := int(ch - '0')
		if tmpNum < 0 || tmpNum > 9 {
			s = s[:index]
			break
		}
	}
	// 将截取的字符串转化成数字
	for _, val := range s {
		tmp := int(val - '0')
		fmt.Println(tmp)
		if res > 214748364 && flag == 1 || (res == 214748364 && tmp >= 7 && flag == 1) {
			fmt.Println("plus")
			return 1<<31 - 1
		}
		if res > 214748364 && flag == -1 || (res == 214748364 && tmp >= 8 && flag == -1) {
			fmt.Println("del")
			return -1 << 31
		}
		res = res*10 + tmp
	}
	if flag == -1 {
		res = -res
	}
	return res
}

func myAtoi02(s string) int {
	// 结果
	res := 0
	// 下标
	i := 0
	n := len(s)
	// flag为1则为正数
	// flag为-1则为负数
	flag := 1
	// 读入字符，删除无用的前导空格
	for ; i < n && s[i] == ' '; i++ {
	}
	if i >= n {
		return 0
	}
	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		flag = -1
	}
	for ; i < n; i++ {
		tmp := int(s[i] - '0')
		// 不是数字直接退出
		if s[i] < 48 || s[i] > 57 {
			break
		}
		if res > 214748364 && flag == 1 || (res == 214748364 && tmp >= 7 && flag == 1) {
			fmt.Println("plus")
			return 1<<31 - 1
		}
		if res > 214748364 && flag == -1 || (res == 214748364 && tmp >= 8 && flag == -1) {
			fmt.Println("del")
			return -1 << 31
		}
		res = res*10 + int(tmp)
	}
	return res * flag
}

func main() {
	fmt.Println(myAtoi01(" "))
	fmt.Println(myAtoi02(" "))
}
