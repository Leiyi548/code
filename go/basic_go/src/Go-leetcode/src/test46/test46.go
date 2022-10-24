package main

import "fmt"

func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		// 递归出口
		if len(path) == len(nums) {
			// copy slice
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			// 退出这层
			return
		}
		for _, num := range nums {
			if visited[num] {
				continue
			}
			path = append(path, num)
			visited[num] = true
			dfs(path)
			// 回溯
			path = path[:len(path)-1]
			visited[num] = false
		}
	}
	dfs([]int{})
	return res
}

func main() {
	nums := [...]int{1, 2, 3}
	visited := map[int]bool{}
	fmt.Println(permute(nums[:]))
	fmt.Println(visited)
}
