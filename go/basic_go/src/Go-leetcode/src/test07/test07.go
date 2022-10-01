package main

import "fmt"

func reverse(x int) int {
	res := 0
	for x != 0 {
		// 每次取末尾数字
		tmp := x % 10
		// 因为 -2^31 <= x <= 2^31 - 1
		// 判断是否大于最大32位整
		if res > 214748364 || (res == 214748364 && tmp > 7) {
			return 0
		}
		// 判断是否小于最小32位整
		if res < -214748364 || (res == -214748364 && tmp < -8) {
			return 0
		}
		res = res*10 + tmp
		x /= 10
	}
	return res
}

func main() {
	// x := 123
	// y := -123
	// fmt.Println(reverse(x))
	// fmt.Println(reverse(y))
	x := -1 / 10
	fmt.Println(x)
}
