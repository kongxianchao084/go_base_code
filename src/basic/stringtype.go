package main

import "fmt"

func main() {
	/*
		字符串：
		1.概念：多个bute的集合，理解为一个字符序列
		2.语法：使用双引号
	*/

	// 1.定义字符串
	var s1 string
	s1 = "王二狗"
	fmt.Printf("%T,%s\n", s1, s1)

	s2 := "hello world"
	fmt.Println(s2)

	// 2.区别：""和''
	v1 := 'A' // rune/int32
	v2 := "A"
	v3 := '中'
	fmt.Printf("%T,%d,%c,%q\n", v1, v1, v1, v1)
	fmt.Printf("%T,%s\n", v2, v2)
	fmt.Printf("%T,%s\n", v3, v3)
	/*
		int32,65,A,'A'
		string,A
		int32,%!s(int32=20013)
	*/

}
