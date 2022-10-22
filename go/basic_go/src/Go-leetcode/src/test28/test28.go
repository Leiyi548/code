package main

import "fmt"

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	findIndex := 0
	for findIndex < n-m+1 {
		// 如果第一個字符不同的話，那麽直接跳過
		for haystack[findIndex] != needle[0] && findIndex < n-m {
			findIndex++
		}
		if haystack[findIndex:findIndex+m] == needle {
			return findIndex
		}
		findIndex++
	}
	return -1
}

func main() {
	haystack01 := "sadbutsad"
	needle01 := "sad"
	haystack02 := "leetcode"
	needle02 := "leeto"
	fmt.Println(strStr(haystack01, needle01))
	fmt.Println(strStr(haystack02, needle02))
}
