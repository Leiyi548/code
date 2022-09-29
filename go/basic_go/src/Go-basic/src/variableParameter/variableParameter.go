package main

import "fmt"

func myfunc(args ...int) {
	fmt.Println(args)
	for _, arg := range args {
		fmt.Print(arg)
	}
	fmt.Println()
}

func main() {
	myfunc(1)
	myfunc(1, 2)
	myfunc(1, 2, 3)
}
