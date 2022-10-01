package main

import (
	"fmt"
	"strings"
)

// more information please see https://leetcode.cn/problems/zigzag-conversion/solution/by-chao-yue-2-l2wn/
func convert02(s string, numRows int) string {
	if len(s) <= 2 {
		return s
	}
	if numRows <= 1 {
		return s
	}
	sArray := make([][]string, numRows, numRows)
	index := 0
	// flag为1代表向下，-1代表向上
	flag := 1
	for i := 0; i < len(s); i++ {
		char := s[i : i+1]
		sArray[index] = append(sArray[index], char)
		// 根据 flag 来判断方向
		if flag == 1 {
			index++
		} else {
			index--
		}
		//改变flag方向
		if index == numRows-1 {
			flag = -1
		} else if index == 0 {
			flag = 1
		}
	}
	//将数组转换成字符串
	var arr []string
	for j := 0; j < len(sArray); j++ {
		arr = append(arr, sArray[j]...)
	}
	result := strings.Join(arr, "")
	return string(result)
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert02(s, numRows))
}
