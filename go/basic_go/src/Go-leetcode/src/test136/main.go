package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}

func main() {
	nums_example1 := []int{2, 2, 1}
	nums_example2 := []int{4, 1, 2, 1, 2}
	fmt.Printf("singleNumber(nums_example1): %v\n", singleNumber(nums_example1))
	fmt.Printf("singleNumber(nums_example2): %v\n", singleNumber(nums_example2))
}
