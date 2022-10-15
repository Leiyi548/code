package main

import "fmt"

// go 定义全局变量
var ans []string

func generateParenthesis(n int) []string {
	dfs("", n, 0, 0)
	return ans
}

func dfs(s string, n int, left int, right int) {
	// 退出递归的条件
	if left == n && right == n {
		ans = append(ans, s)
	}
	// 左括号，右括号
	if left > n || right > n || right > left {
		return
	}
	// 先加入左括号
	dfs(s+"(", n, left+1, right)
	// 再加入右括号
	dfs(s+")", n, left, right+1)
}

func generateParenthesis_simply(n int) []string {
	var ans []string
	// 声明函数
	var dfs func(string, int, int, int)
	dfs = func(s string, n int, left int, right int) {
		// 退出递归的条件
		if left == n && right == n {
			ans = append(ans, s)
		}
		// 左括号，右括号
		if left > n || right > n || right > left {
			return
		}
		// 先加入左括号
		dfs(s+"(", n, left+1, right)
		// 再加入右括号
		dfs(s+")", n, left, right+1)
	}
	dfs("", n, 0, 0)
	return ans
}

func main() {
	fmt.Printf("generateParenthesis(1): %v\n", generateParenthesis(1))
}
