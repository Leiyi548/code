package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	// 先对数组进行排序
	sort.Ints(nums)
	var res [][]int
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	for first := 0; first < n; first++ {
		// 排除重复的数字
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 固定了第一个数那么，我们现在的问题就变成了三数之和
		// 另外3个数分别是second,third,four
		// 那么 nums[first] + nums[second] + nums[third] + nums[four] = target
		// nums[second] + nums[third] + nums[four] = target- nums[frist]
		res = append(res, threeSum(nums[first+1:], target-nums[first], nums[first])...)
	}
	return res
}

func threeSum(nums []int, target int, firstNum int) [][]int {
	var res [][]int
	n := len(nums)
	if n < 3 {
		return nil
	}
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second := first + 1
		third := n - 1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == target {
				res = append(res, []int{firstNum, nums[first], nums[second], nums[third]})
				// 到下一个不重复的数字
				secondCopy := second
				thirdCopy := third
				for secondCopy <= third && nums[secondCopy] == nums[second] {
					secondCopy++
				}
				for thirdCopy >= second && nums[thirdCopy] == nums[third] {
					thirdCopy--
				}
				second = secondCopy
				third = thirdCopy
				if second == third {
					break
				}
			}
			if sum > target {
				// 排除重复的数字
				thirdCopy := third - 1
				for thirdCopy > second && nums[thirdCopy] == nums[third] {
					thirdCopy--
				}
				third = thirdCopy
			}
			if sum < target {
				// 排除重复的数字
				secondCopy := second + 1
				for third > secondCopy && nums[secondCopy] == nums[second] {
					secondCopy++
				}
				second = secondCopy
			}
		}
	}
	return res
}

func main() {
	nums := [...]int{1, 0, -1, 0, -2, 2}
	target := 0
	nums2 := [...]int{2, 2, 2, 2, 2}
	target2 := 8
	nums3 := [...]int{-3, -2, -1, 0, 0, 1, 2, 3}
	target3 := 0
	nums4 := [...]int{-2, -1, -1, 1, 1, 2, 2}
	target4 := 0
	fmt.Printf("fourSum(nums[:], target): %v\n", fourSum(nums[:], target))
	fmt.Printf("fourSum(nums[:], target): %v\n", fourSum(nums2[:], target2))
	fmt.Printf("fourSum(nums[:], target): %v\n", fourSum(nums3[:], target3))
	fmt.Printf("fourSum(nums[:], target): %v\n", fourSum(nums4[:], target4))
}

// helloWorld
