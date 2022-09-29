package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)

	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		// 每次移动滑动窗口，都是最左边的字符，删除它向左走。
		if i != 0 {
			delete(m, s[i-1])
		}
		// m[s[rk+1]] == 0 ，代表s[rk+1]没有出现过
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	s := "bbsssbbbbbbbbbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
