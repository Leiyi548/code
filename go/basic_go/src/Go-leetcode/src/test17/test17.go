package main

import (
	"fmt"
)

func letterCombinationsBack(digits string) []string {
	n := len(digits)
	var letterMap map[string][]string = map[string][]string{"2": {"a", "b", "c"},
		"3": {"d", "e", "f"}, "4": {"g", "h", "i"}, "5": {"j", "k", "l"},
		"6": {"m", "n", "o"}, "7": {"p", "q", "r", "s"}, "8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"}}
	fmt.Printf("letterMap: %v\n", letterMap)
	if n == 0 {
		return nil
	}
	for _, v := range digits {
		fmt.Printf("v: %c\n", v)
		fmt.Printf("letterMap[v]: %v\n", letterMap[string(v)])
	}
	return nil
}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

func main() {
	digits := "23"
	fmt.Printf("letterCombinations(digits): %v\n", letterCombinations(digits))
}
