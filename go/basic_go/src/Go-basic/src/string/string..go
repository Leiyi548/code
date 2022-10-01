package main

import "fmt"

func readMultiString() {
	// more information please see <http://c.biancheng.net/view/17.html>
	// 多行字符串
	const multiString01 = ` first_line second_line third_line `
	const multiString02 = `
first_line
second_line
third_line
`
	fmt.Println(multiString01)
	fmt.Println(multiString02)
}

func modifyString() {
	// Go 语言的字符串无法直接修改每一个字符元素，只能通过重新构造新的字符串并赋值给原来的字符串变量实现。
	// more information please see <https://www.cnblogs.com/niuben/p/12523396.html#:~:text=Go,语言的字符串无法直接修改每一个字符元素，只能通过重新构造新的字符串并赋值给原来的字符串变量实现。>
	var s string = "abc"
	b := []byte(s)
	b[1] = 'z'
	fmt.Printf("%s\n", s)
	fmt.Printf("%s", b)
}

func neoModifyString() {
	s2 := "小白兔"
	s3 := []rune(s2)        //把字符串强制转成rune切片
	s3[0] = '大'             //注意 这里需要使用单引号的字符，而不是双引号的字符串
	fmt.Println(string(s3)) //把rune类型的s3强转成字符串
}

func readString() {
	res := 0
	s := "12345"
	// fmt.Printf("%c\n", s[1])
	for _, tmp := range s {
		res = res*10 + int(tmp-'0')
		fmt.Println(tmp - '0')
	}
	fmt.Println(res)
}

func main() {
	readString()
}
