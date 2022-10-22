package main

import "fmt"

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 左指针
	index := 0
	// 右指针
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	nums = nums[:index]
	fmt.Println(nums)
	return index
}

func main() {
	nums1 := [...]int{0, 1, 2, 2, 3, 0, 4, 2}
	val1 := 2
	nums2 := [...]int{3, 2, 2, 3}
	val2 := 3
	fmt.Println(removeElement(nums1[:], val1))
	fmt.Println(removeElement(nums2[:], val2))
}
