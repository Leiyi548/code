package main

import "fmt"

// 暴力解法
func longestPalindrome01(s string) string {
	n := len(s)
	var res string
	maxLen := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if isPalindrome(s[i:j]) && (j-i) > maxLen {
				res = s[i:j]
				maxLen = j - i
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	end := len(s) - 1
	start := 0
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// 中心扩散
func longestPalindrome02(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 奇数
		left1, right1 := expandAroundCenter(s, i, i)
		// 偶数
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

// 动态规划
func longestPalindrome03(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	maxSubStr := ""
	maxLen := 0
	var dp [1001][1001]bool
	// 初始化动态规划的条件
	// 所有长度为1的子串都是回文串
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	for col := 1; col < n; col++ {
		for row := 0; row <= col; row++ {
			if s[row] == s[col] {
				if (col-row == 1) || dp[row+1][col-1] {
					dp[row][col] = true
				}
			}
			if dp[row][col] && col-row+1 > maxLen {
				maxLen = col - row + 1
				maxSubStr = s[row : col+1]
			}
		}
	}
	return maxSubStr
}

func main() {
	s := "bb"
	//fmt.Println(longestPalindrome01(s))
	fmt.Println(longestPalindrome02(s))
	fmt.Println(longestPalindrome03(s))
}
