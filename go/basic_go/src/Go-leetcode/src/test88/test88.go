package main

import (
	"fmt"
	"sort"
)

func merge01(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	index := 0
	first := 0
	second := 0
	for index < m+n {
		if nums1[first] < nums2[second] && first < m {
			nums1[index] = nums1[first]
			index++
			first++
		} else {
			nums1[index] = nums2[second]
			index++
			second++
		}
	}
}

func merge02(nums1 []int, m int, nums2 []int, n int) {
	// 将两个数组合并，然后再进行排序
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func main() {
	nums1 := [...]int{1, 2, 3, 0, 0, 0}
	nums2 := [...]int{2, 5, 6}
	m := 3
	n := 3
	merge02(nums1[:], m, nums2[:], n)
	fmt.Println(nums1)
}
