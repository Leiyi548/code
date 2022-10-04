package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func maxArea01(height []int) int {
	n := len(height)
	maxArea := 0
	for left := 0; left < n; left++ {
		for right := left + 1; right < n; right++ {
			if height[left] > height[right] {
				tmpArea := (right - left) * height[right]
				maxArea = max(maxArea, tmpArea)
			} else {
				tmpArea := (right - left) * height[left]
				maxArea = max(maxArea, tmpArea)
			}
		}
	}
	return maxArea
}

func maxArea02(height []int) int {
	lp := 0
	rp := len(height) - 1
	res := 0

	for lp < rp {
		if height[lp] < height[rp] {
			// 因为长度是由短板决定的，所以我们要移动到长板去！
			res = max(res, (rp-lp)*height[lp])
			lp++
		} else {
			res = max(res, (rp-lp)*height[rp])
			rp--
		}
	}
	return res
}

func main() {
	arr := [...]int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea01(arr[:]))
	fmt.Println(maxArea02(arr[:]))
}
