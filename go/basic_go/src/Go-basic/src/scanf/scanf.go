package main

import "fmt"

func main() {
	var length int
	var num int
	fmt.Println("Please input length")
	scanfNumber, err := fmt.Scanf("%d%d", &length, &num)
	if err != nil {
		return
	}
	fmt.Println(scanfNumber)
}
