package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	sort.Ints(nums)
	if n < 0 {
		return nil
	}

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 的后续的增加
			// 就不会满足 a+b+c=0 并且 b < c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func main() {
	nums01 := [...]int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("threeSum(nums[:]): %v\n", threeSum(nums01[:]))
	nums02 := [...]int{0, 1, 1}
	fmt.Printf("threeSum(nums[:]): %v\n", threeSum(nums02[:]))
	nums03 := [...]int{0, 0, 0}
	fmt.Printf("threeSum(nums[:]): %v\n", threeSum(nums03[:]))
}
