package main

import "fmt"

func searchInsert01(nums []int, target int) int {
	index := 0
	for _, v := range nums {
		if v < target {
			index++
		} else {
			return index
		}
	}
	return index
}

func searchInsert02(nums []int, target int) int {
	n := len(nums)
	left, right, ans := 0, n-1, n
	for left <= right {
		mid := ((right - left) >> 1) + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func main() {
	nums := [...]int{1, 3, 5, 6}
	target := 7
	fmt.Println(searchInsert01(nums[:], target))
	fmt.Println(searchInsert02(nums[:], target))
}
