package main

import "fmt"

func removeDuplicates(nums []int) int {
	// 使用双指针，left记录这个数字的开头，right记录下一个数字
	n := len(nums)
	prev, cur := 0, 0
	// index 记录当前不重复数字的下标，一开始在0
	index := 0
	// special case
	if n == 1 {
		return 1
	}
	for cur < n {
		if nums[cur] != nums[prev] {
			index++
			nums[index] = nums[cur]
			prev = cur
		}
		cur++
	}
	return index + 1
}

func main() {
	nums := [...]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums[:]))
}
