package main

import "fmt"

func forSum(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func testRange(s string) {
	for _, value := range s {
		fmt.Printf("%c", value)
	}
}

func main() {
	fmt.Println(forSum(3))
	testRange("abc")
}
