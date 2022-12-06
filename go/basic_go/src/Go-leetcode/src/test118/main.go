package main

import "fmt"

func generate(numRows int) [][]int {
	// 创建一个二维的空数组！
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}

func main() {
	numRows := 5
	fmt.Printf("generate(numRows): %v\n", generate(numRows))
}
