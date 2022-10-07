package main

import "fmt"

func longestCommonPrefix01(strs []string) (maxPrefix string) {
	if len(strs) == 0 {
		return ""
	}
	maxPrefix = strs[0]
	for i := 1; i < len(strs); i++ {
		maxPrefix = lcp(maxPrefix, strs[i])
		if len(maxPrefix) == 0 {
			break
		}
	}
	return
}

func longestCommonPrefix02(strs []string) (maxPrefix string) {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return
}

func lcp(s1, s2 string) string {
	length := min(len(s1), len(s2))
	index := 0
	for index < length && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func main() {
	s1 := [...]string{"flower", "flow", "flight"}
	s2 := [...]string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix01(s1[:]))
	fmt.Println(longestCommonPrefix02(s2[:]))
}
