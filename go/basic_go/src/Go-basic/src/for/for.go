package main

import "fmt"

func forSum(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(forSum(3))
}
