package main

import "fmt"

func intToRoman(num int) string {
	res := ""
	for num > 0 {
		for num >= 1000 {
			num -= 1000
			res += "M"
		}
		for num >= 900 {
			num -= 900
			res += "CM"
		}
		for num >= 500 {
			num -= 500
			res += "D"
		}
		for num >= 400 {
			num -= 400
			res += "CD"
		}
		for num >= 100 {
			num -= 100
			res += "C"
		}
		for num >= 90 {
			num -= 90
			res += "XC"
		}
		for num >= 50 {
			num -= 50
			res += "L"
		}
		for num >= 40 {
			num -= 40
			res += "XL"
		}
		for num >= 10 {
			num -= 10
			res += "X"
		}
		for num >= 9 {
			num -= 9
			res += "IX"
		}
		for num >= 5 {
			num -= 5
			res += "V"
		}
		for num >= 4 {
			num -= 4
			res += "IV"
		}
		for num >= 1 {
			num -= 1
			res += "I"
		}
	}
	return res
}

func main() {
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(3900))
}
