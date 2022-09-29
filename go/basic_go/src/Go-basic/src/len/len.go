package main

import "fmt"

func ReadString() {
	s := "Hello王"
	// 我们发现英文可以正常输出，但是中文通过这种方式输出会乱码
	// H e l l o ç
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 {
			fmt.Printf("%c ", s[i])
		} else {
			fmt.Printf("%c", s[i])
		}
	}
	fmt.Println()
}

func ReadZh() {
	s := "Hello王"
	// H e l l o 王
	for _, c := range s {
		fmt.Printf("%c ", c) //%c 字符
	}
	fmt.Println()
}

func basicLen() {
	// len()获得的是 byte 字节的数量，一个中文占用3个字节
	s := "Hello王"
	sHello := "Hello"
	sWang := "王"
	//len()获得的是 byte 字节的数量
	fmt.Println(len(s))      // 8
	fmt.Println(len(sHello)) // 5
	fmt.Println(len(sWang))  // 3
}

func main() {
	ReadString()
	ReadZh()
}
