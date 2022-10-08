package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest01(nums []int, target int) int {
	res := 0
	n := len(nums)
	minDistance := 1<<31 - 1
	// special case
	if n == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	for first := 0; first < n; first++ {
		for second := first + 1; second < n; second++ {
			for third := second + 1; third < n; third++ {
				sum := nums[first] + nums[second] + nums[third]
				if math.Abs(float64(target)-float64(sum)) < float64(minDistance) {
					res = sum
					minDistance = int(math.Abs(float64(target) - float64(sum)))
				}
			}
		}
	}
	return res
}

// 双指针
func threeSumClosest02(nums []int, target int) int {
	// 首先先对数组进行排序
	sort.Ints(nums)
	n := len(nums)
	best := math.MaxInt32
	// 枚举 first
	for first := 0; first < n; first++ {
		// 保证和上一次枚举的元素不相等
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 使用双指针枚举b和c
		second, third := first+1, n-1
		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			// 更新best
			if abs(sum-target) < abs(best-target) {
				best = sum
			}
			// 如果直接等于target的话就直接返回
			if sum == target {
				return target
			}
			// 这个时候sum大于target，我们需要移动third指针来减小sum,让sum接近target
			if sum > target {
				// 为什么用thirdCopy因为后面我们需要用这个third来判断是否有重复数字
				thirdCopy := third
				// 排除掉重复元素
				for second < thirdCopy && nums[thirdCopy] == nums[third] {
					thirdCopy--
				}
				// 更新third
				third = thirdCopy
			}
			// 同理如上
			if sum < target {
				secondCopy := second
				for secondCopy < third && nums[secondCopy] == nums[second] {
					secondCopy++
				}
				second = secondCopy
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	nums := [...]int{-1, 2, 1, -4}
	target := 1
	fmt.Printf("threeSumClosest(nums[:], target): %v\n", threeSumClosest01(nums[:], target))
	fmt.Printf("threeSumClosest(nums[:], target): %v\n", threeSumClosest02(nums[:], target))
}
