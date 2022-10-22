package main

import (
	"fmt"
	"math"
)

// dividend 被除数
// divisor 除数
func divide(dividend int, divisor int) int {
	ans := 0
	// 溢出情况
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	// sign 记录正负号
	// sign 为1，代表结果为正
	// sign 为2，代表结果为负
	sign := -1
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = 1
	}
	// 全部转换成负数防止溢出
	if dividend > 0 {
		dividend = -dividend
	}
	if divisor > 0 {
		divisor = -divisor
	}

	// 陪乘法，注意都是负数了，比较的时候与正数相反
	// 简单点理解，就是每次减去除法的 2^x 陪数，剩下的部分继续按这样的规则继续
	for dividend <= divisor {
		tmp, count := divisor, 1
		// 这里注意不要写成 tmp + tmp >= dividend，这样写加法有可能会溢出导致死循环
		// 比如，我这里的 tmp超过最大负数的一半，两个相加不就溢出了。
		for tmp >= dividend-tmp {
			// tmp和count 每次增加一陪，所以叫陪增
			tmp += tmp
			count += count
		}
		// 被除法减去除法的 2^x陪数做为新的被除数
		dividend -= tmp
		// count 即 2^x
		ans += count
	}

	if sign == -1 {
		return -ans
	} else {
		return ans
	}
}

func main() {
	dividend := 10
	divisor := 3
	fmt.Println(divide(dividend, divisor))
}
